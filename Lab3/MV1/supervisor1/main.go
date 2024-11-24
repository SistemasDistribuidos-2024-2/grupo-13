package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-server/proto/grpc-server/proto"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	brokerAddress = "localhost:50050" // Dirección del broker
)

var vectorClocks = make(map[string][]int32) // Almacén de relojes vectoriales por región
var mutex sync.Mutex                        // Para proteger el acceso al almacén de relojes

func main() {
	// Verificar que se pasó el argumento
	if len(os.Args) < 2 {
		log.Fatal("Por favor, especifica el argumento: 1 para Supervisor 1 o 2 para Supervisor 2")
	}

	// Convertir el argumento a un entero
	supervisorID, err := strconv.Atoi(os.Args[1])
	if err != nil || (supervisorID != 1 && supervisorID != 2) {
		log.Fatal("El argumento debe ser 1 o 2")
	}

	// Ejecutar el supervisor correspondiente
	runSupervisor(supervisorID)
}

// runSupervisor configura y ejecuta el supervisor según el ID especificado
func runSupervisor(supervisorID int) {
	// Configurar el puerto según el supervisor ID
	port := 50050 + supervisorID
	log.Printf("Supervisor %d inicializado en el puerto %d\n", supervisorID, port)

	// Mostrar menú interactivo
	for {
		fmt.Println("\nMenú de opciones:")
		fmt.Println("1. AgregarProducto")
		fmt.Println("2. RenombrarProducto")
		fmt.Println("3. ActualizarValor")
		fmt.Println("4. BorrarProducto")
		fmt.Println("5. Salir")
		fmt.Print("Selecciona una opción: ")

		reader := bufio.NewReader(os.Stdin)
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			handleRequest("AgregarProducto", reader)
		case "2":
			handleRequest("RenombrarProducto", reader)
		case "3":
			handleRequest("ActualizarValor", reader)
		case "4":
			handleRequest("BorrarProducto", reader)
		case "5":
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

// handleRequest procesa las solicitudes del menú
func handleRequest(action string, reader *bufio.Reader) {
	var region, product, newProduct string
	var value int

	fmt.Print("Ingresa el nombre de la región: ")
	region, _ = reader.ReadString('\n')
	region = strings.TrimSpace(region)

	fmt.Print("Ingresa el nombre del producto: ")
	product, _ = reader.ReadString('\n')
	product = strings.TrimSpace(product)

	if action == "ActualizarValor" || action == "AgregarProducto" {
		fmt.Print("Ingresa el valor: ")
		fmt.Scanf("%d\n", &value)
	} else if action == "RenombrarProducto" {
		fmt.Print("Ingresa el nuevo nombre del producto: ")
		newProduct, _ = reader.ReadString('\n')
		newProduct = strings.TrimSpace(newProduct)
	}

	fmt.Printf("Petición enviada al broker: %s %s %s [%d]\n", action, region, product, value)
	requestBroker(action, region, product, newProduct, value)
}

// requestBroker se conecta al broker y maneja la respuesta
func requestBroker(action, region, product, newProduct string, value int) {
	// Conexión al broker
	conn, err := grpc.Dial(brokerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar con el broker: %v", err)
	}
	defer conn.Close()

	client := pb.NewHextechServiceClient(conn)
	var address string

	// Solicitar dirección al broker
	switch action {
	case "AgregarProducto":
		resp, err := client.AddProductBroker(context.Background(), &pb.AddProductRequest{
			Region:   region,
			Product:  product,
			Quantity: int32(value),
		})
		if err != nil {
			log.Fatalf("Error al enviar la solicitud al broker: %v", err)
		}
		address = resp.Address
	case "RenombrarProducto":
		resp, err := client.RenameProductBroker(context.Background(), &pb.RenameProductRequest{
			Region:     region,
			OldProduct: product,
			NewProduct: newProduct,
		})
		if err != nil {
			log.Fatalf("Error al enviar la solicitud al broker: %v", err)
		}
		address = resp.Address
	case "ActualizarValor":
		resp, err := client.UpdateProductBroker(context.Background(), &pb.UpdateProductRequest{
			Region:   region,
			Product:  product,
			Quantity: int32(value),
		})
		if err != nil {
			log.Fatalf("Error al enviar la solicitud al broker: %v", err)
		}
		address = resp.Address
	case "BorrarProducto":
		resp, err := client.DeleteProductBroker(context.Background(), &pb.DeleteProductRequest{
			Region:  region,
			Product: product,
		})
		if err != nil {
			log.Fatalf("Error al enviar la solicitud al broker: %v", err)
		}
		address = resp.Address
	}

	fmt.Printf("Conectando con el servidor %s\n", address)
	sendToServer(action, address, region, product, newProduct, value)
}

// sendToServer se conecta al servidor para reenviar la solicitud
func sendToServer(action, address, region, product, newProduct string, value int) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar con el servidor: %v", err)
	}
	defer conn.Close()

	client := pb.NewHextechServiceClient(conn)
	var clockResponse *pb.ClockResponse

	switch action {
	case "AgregarProducto":
		clockResponse, err = client.AddProductServer(context.Background(), &pb.AddProductRequest{
			Region:   region,
			Product:  product,
			Quantity: int32(value),
		})
	case "RenombrarProducto":
		clockResponse, err = client.RenameProductServer(context.Background(), &pb.RenameProductRequest{
			Region:     region,
			OldProduct: product,
			NewProduct: newProduct,
		})
	case "ActualizarValor":
		clockResponse, err = client.UpdateProductServer(context.Background(), &pb.UpdateProductRequest{
			Region:   region,
			Product:  product,
			Quantity: int32(value),
		})
	case "BorrarProducto":
		clockResponse, err = client.DeleteProductServer(context.Background(), &pb.DeleteProductRequest{
			Region:  region,
			Product: product,
		})
	}

	if err != nil {
		log.Fatalf("Error al enviar la solicitud al servidor: %v", err)
	}

	fmt.Printf("Petición enviada al servidor: %s %s %s [%d]\n", action, region, product, value)
	handleClockConsistency(region, clockResponse.VectorClock)
}

// handleClockConsistency maneja la consistencia del reloj vectorial
func handleClockConsistency(region string, receivedClock []int32) {
	mutex.Lock()
	defer mutex.Unlock()

	// Verificar si ya existe un reloj vectorial para la región
	if storedClock, exists := vectorClocks[region]; exists {
		for i, v := range receivedClock {
			if v < storedClock[i] {
				fmt.Printf("Violación detectada, enviada al broker: Forzar Merge %s %v\n", region, receivedClock)
				forceMergeWithBroker(region, receivedClock)
				return
			}
		}
	}

	// Actualizar el reloj vectorial
	vectorClocks[region] = receivedClock
	fmt.Printf("[%s] actualizada: %v\n", region, receivedClock)
}

// forceMergeWithBroker envía una solicitud al broker para forzar un merge
func forceMergeWithBroker(region string, vectorClock []int32) {
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
		log.Fatalf("Error al enviar la solicitud de merge al broker: %v", err)
	}

	fmt.Println("Merge forzado confirmado por el broker.")
}