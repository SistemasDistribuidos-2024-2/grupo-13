package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "grpc-server/proto/grpc-server/proto"
)

const brokerAddress = "broker_container:50050" // Dirección del broker

var mutex sync.Mutex
var serverRegions = map[string]string{
	"hextech1_container:50054": "S1",
	"hextech2_container:50055": "S2",
	"hextech3_container:50056": "S3",
}

func main() {
	time.Sleep(1 * time.Second)
	for {
		fmt.Println("\nMenú de Jayce:")
		fmt.Println("1. Consultar producto")
		fmt.Println("2. Salir")
		fmt.Print("Selecciona una opción: ")

		reader := bufio.NewReader(os.Stdin)
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			handleGetProduct(reader)
		case "2":
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

// handleGetProduct procesa la solicitud de consulta de un producto
func handleGetProduct(reader *bufio.Reader) {
	fmt.Print("Ingresa el nombre de la región: ")
	region, _ := reader.ReadString('\n')
	region = strings.TrimSpace(region)

	fmt.Print("Ingresa el nombre del producto: ")
	product, _ := reader.ReadString('\n')
	product = strings.TrimSpace(product)

	address := requestBroker(region, product)
	if address == "" {
		log.Println("No se pudo obtener una dirección del broker.")
		return
	}

	fmt.Printf("Conectando con el servidor en %s\n", address)
	queryServer(address, region, product)
}

// requestBroker consulta al broker para obtener la dirección de un servidor
func requestBroker(region, product string) string {
	conn, err := grpc.Dial(brokerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar con el broker: %v", err)
	}
	defer conn.Close()

	client := pb.NewHextechServiceClient(conn)
	resp, err := client.GetProductBroker(context.Background(), &pb.GetProductRequest{
		Region:  region,
		Product: product,
	})
	if err != nil {
		log.Fatalf("Error al enviar la solicitud al broker: %v", err)
	}

	fmt.Printf("Dirección recibida del broker: %s\n", resp.Address)
	return resp.Address
}

// queryServer consulta al servidor y maneja la respuesta
func queryServer(address, region, product string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar con el servidor: %v", err)
	}
	defer conn.Close()

	client := pb.NewHextechServiceClient(conn)
	resp, err := client.GetProductServer(context.Background(), &pb.GetProductRequest{
		Region:  region,
		Product: product,
	})
	if err != nil {
		log.Fatalf("Error al consultar al servidor: %v", err)
	}

	fmt.Printf("Respuesta del servidor: Reloj vectorial %v, Cantidad %d\n", resp.VectorClock, resp.Quantity)
	handleConsistency(address, region, product, resp.VectorClock, resp.Quantity)
}

// handleConsistency verifica la consistencia Monotonic Read y actualiza los registros
func handleConsistency(address, region, product string, vectorClock []int32, quantity int32) {
	serverID := serverRegions[address]
	filename := fmt.Sprintf("%s_%s.txt", serverID, region)

	mutex.Lock()
	defer mutex.Unlock()

	// Leer el archivo si existe
	existingClock := readClockFromFile(filename)

	// Verificar consistencia Monotonic Read
	if !isMonotonicRead(existingClock, vectorClock) {
		fmt.Printf("Inconsistencia detectada, notificando al broker: Forzar Merge %s %v\n", region, vectorClock)
		notifyBroker(region, vectorClock)
		return
	}

	// Actualizar el archivo
	updateFile(filename, vectorClock, product, quantity)
	fmt.Printf("[%s] [%s] actualizado correctamente.\n", serverID, region)
}

// isMonotonicRead verifica si un reloj satisface la consistencia Monotonic Read
func isMonotonicRead(existingClock, newClock []int32) bool {
	if len(existingClock) == 0 {
		return true // Si no hay un reloj previo, no hay inconsistencia
	}

	for i := range newClock {
		if newClock[i] < existingClock[i] {
			return false
		}
	}
	return true
}

// readClockFromFile lee el reloj vectorial del archivo, si existe
func readClockFromFile(filename string) []int32 {
	file, err := os.Open(filename)
	if err != nil {
		return nil // Archivo no existe
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		return parseClock(line)
	}
	return nil
}

// updateFile actualiza el archivo con el nuevo reloj vectorial y producto
func updateFile(filename string, vectorClock []int32, product string, quantity int32) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error al actualizar el archivo %s: %v", filename, err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, _ = writer.WriteString(formatClock(vectorClock) + "\n")
	_, _ = writer.WriteString(fmt.Sprintf("%s: %d\n", product, quantity))
	writer.Flush()
}

// parseClock convierte una cadena de texto en un reloj vectorial
func parseClock(line string) []int32 {
	parts := strings.Split(line, ",")
	clock := make([]int32, len(parts))
	for i, p := range parts {
		val, _ := strconv.Atoi(strings.TrimSpace(p))
		clock[i] = int32(val)
	}
	return clock
}

// formatClock convierte un reloj vectorial en una cadena de texto
func formatClock(clock []int32) string {
	parts := make([]string, len(clock))
	for i, v := range clock {
		parts[i] = strconv.Itoa(int(v))
	}
	return strings.Join(parts, ",")
}

// notifyBroker notifica al broker sobre una inconsistencia
func notifyBroker(region string, vectorClock []int32) {
	conn, err := grpc.Dial(brokerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar con el broker: %v", err)
	}
	defer conn.Close()

	client := pb.NewHextechServiceClient(conn)
	_, err = client.ForceMerge(context.Background(), &pb.ErrorMergeRequest{
		Region:      region,
		VectorClock: vectorClock,
	})
	if err != nil {
		log.Fatalf("Error al notificar al broker: %v", err)
	}

	fmt.Println("Notificación de merge forzado enviada al broker.")
}
