package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	// Importa el paquete proto generado (reemplaza con tu ruta actual)
	pb "grpc-server/proto/grpc-server/proto"
)

type servidorTai struct {
	pb.UnimplementedTaiServiceServer
	mu                sync.Mutex
	vida              int32
	ctx               context.Context
	cancelar          context.CancelFunc
	diaboromonClient  pb.DiaboromonServiceClient
	primaryNodeClient pb.PrimaryNodeServiceClient
	terminarChan      chan struct{}
}

func (s *servidorTai) DiaboromonAttack(ctx context.Context, req *pb.Empty) (*pb.AttackResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.vida -= 10
	fmt.Println("Diaboromon ha atacado, restando 10 puntos de vida")
	if s.vida <= 0 {
		fmt.Println("Diaboromon ha ganado")
		// Responder true al ataque final de Diaboromon
		response := &pb.AttackResponse{Success: true}
		// Iniciar proceso de término después de responder
		go func() {
			s.enviarSenalTermino("derrota")
		}()
		return response, nil
	}
	// Responder false para que Diaboromon continúe atacando
	return &pb.AttackResponse{Success: false}, nil
}

func (s *servidorTai) enviarSenalTermino(resultado string) {
	_, err := s.primaryNodeClient.SendTerminationSignal(s.ctx, &pb.TerminateProcess{Result: resultado})
	if err != nil {
		fmt.Printf("Error al enviar señal de término al Nodo Primario: %v\n", err)
	} else {
		fmt.Println("Señal de término enviada al Nodo Primario")
	}
	// Señal para terminar la ejecución
	s.terminarChan <- struct{}{}
}

func leerVI() (int32, error) {
	archivo, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	if scanner.Scan() {
		linea := scanner.Text()
		valores := strings.Split(linea, ",")
		for i := range valores {
			valores[i] = strings.TrimSpace(valores[i])
		}
		if len(valores) < 5 {
			return 0, fmt.Errorf("input.txt no tiene suficientes valores")
		}
		// Posiciones: PS=0, TE=1, TD=2, CD=3, VI=4
		viFloat, err := strconv.ParseFloat(valores[4], 64)
		if err != nil {
			return 0, err
		}
		VI := int32(viFloat)
		return VI, nil
	} else {
		return 0, fmt.Errorf("input.txt está vacío")
	}
}

func main() {
	// Leer vida inicial de Tai desde input.txt
	VI, err := leerVI()
	if err != nil {
		log.Fatalf("Error al leer input.txt: %v", err)
	}

	s := &servidorTai{
		vida:         VI,
		terminarChan: make(chan struct{}),
	}
	s.ctx, s.cancelar = context.WithCancel(context.Background())

	// Inicia el servidor gRPC de Tai
	go func() {
		lis, err := net.Listen("tcp", ":50052")
		if err != nil {
			log.Fatalf("No se pudo escuchar en el puerto 50052: %v", err)
		}
		servidorGRPC := grpc.NewServer()
		pb.RegisterTaiServiceServer(servidorGRPC, s)
		if err := servidorGRPC.Serve(lis); err != nil {
			log.Fatalf("No se pudo servir: %v", err)
		}
	}()

	// Conecta con Diaboromon
	connDiaboromon, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con Diaboromon: %v", err)
	}
	defer connDiaboromon.Close()
	s.diaboromonClient = pb.NewDiaboromonServiceClient(connDiaboromon)

	// Conecta con el Nodo Primario
	connPrimaryNode, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el Nodo Primario: %v", err)
	}
	defer connPrimaryNode.Close()
	s.primaryNodeClient = pb.NewPrimaryNodeServiceClient(connPrimaryNode)

	// Envía señal de inicio a Diaboromon
	_, err = s.diaboromonClient.StartDiaboromon(s.ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error al iniciar Diaboromon: %v", err)
	}
	fmt.Println("Tai ha enviado la señal de inicio a Diaboromon.")

	// Gorutina para esperar la señal de término
	go func() {
		<-s.terminarChan
		fmt.Println("Nodo Tai: Terminando ejecución.")
		os.Exit(0)
	}()

	// Bucle para leer comandos desde la consola
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Seleccione una opción:")
		fmt.Println("1. Atacar a Diaboromon")
		fmt.Println("2. Salir")
		fmt.Print("Ingrese el número de su elección: ")
		if !scanner.Scan() {
			break
		}
		entrada := strings.TrimSpace(scanner.Text())
		switch entrada {
		case "1":
			// Solicitar datos al Nodo Primario
			fmt.Println("Recolectando información de los Digimons...")
			respuesta, err := s.primaryNodeClient.GetAttackData(s.ctx, &pb.TaiRequest{Message: "Solicitud de datos para ataque"})
			if err != nil {
				fmt.Printf("Error al obtener datos del Nodo Primario: %v\n", err)
				continue
			}
			datosRecolectados := respuesta.DataCollected

			// Atacar a Diaboromon con los datos recolectados
			fmt.Println("Nodo Tai atacó a Diaboromon")
			ataqueRespuesta, err := s.diaboromonClient.AttackDiaboromon(s.ctx, &pb.AttackRequest{AttackValue: datosRecolectados})
			if err != nil {
				fmt.Printf("Error al atacar a Diaboromon: %v\n", err)
				continue
			}
			if ataqueRespuesta.Success {
				fmt.Println("¡Nodo Tai ha ganado!")
				// Enviar señal de término al Nodo Primario y terminar ejecución
				go func() {
					s.enviarSenalTermino("victoria")
				}()
				// Esperar a que se envíe la señal de término
				<-s.terminarChan
				os.Exit(0)
			} else {
				fmt.Println("Diaboromon repelió el ataque")
				s.mu.Lock()
				s.vida -= 10
				if s.vida <= 0 {
					s.mu.Unlock()
					fmt.Println("Diaboromon ha ganado")
					// Enviar señal de término al Nodo Primario y terminar ejecución
					go func() {
						s.enviarSenalTermino("derrota")
					}()
					// Esperar a que se envíe la señal de término
					<-s.terminarChan
					os.Exit(0)
				}
				s.mu.Unlock()
			}
		case "2":
			fmt.Println("Terminando el programa.")
			s.cancelar()
			os.Exit(0)
		default:
			fmt.Println("Opción no válida. Por favor, ingrese '1' para atacar o '2' para salir.")
		}
	}

	// Esperar a que se envíe la señal de término antes de salir
	<-s.terminarChan
	os.Exit(0)
}
