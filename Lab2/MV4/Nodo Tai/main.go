package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	"google.golang.org/grpc"
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
	fmt.Println("Vida restante: ", s.vida)
	if s.vida <= 0 {
		fmt.Println(`
			██████╗ ███████╗██████╗ ██████╗  ██████╗ ████████╗ █████╗ 
			██╔══██╗██╔════╝██╔══██╗██╔══██╗██╔═══██╗╚══██╔══╝██╔══██╗
			██║  ██║█████╗  ██████╔╝██████╔╝██║   ██║   ██║   ███████║
			██║  ██║██╔══╝  ██╔══██╗██╔══██╗██║   ██║   ██║   ██╔══██║
			██████╔╝███████╗██║  ██║██║  ██║╚██████╔╝   ██║   ██║  ██║
			╚═════╝ ╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝    ╚═╝   ╚═╝  ╚═╝
			⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠏⠙⠋⠉⠛⠛⠛⢿⣿⣿⣿⣿⣿⣿⣿⣿
			⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠏⣀⣀⠀⠀⠀⣀⣀⣀⡨⢿⣿⣿⣿⣿⣿⣿
			⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⢠⣾⣿⣿⡷⠒⣾⣿⣿⣿⣿⡄⢿⣿⣿⣿⣿⣿
			⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⣾⣿⣿⡿⣃⣼⣿⠿⣿⣿⣿⡇⣼⣿⣿⣿⣿⣿
			⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠉⠉⠁⠀⠛⠉⢙⠁⠀⠀⠀⠠⣿⣿⣿⣿⣿⣿
			⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣄⡀⠤⣂⣀⢀⡘⠺⠦⣄⣠⡤⣿⣿⣿⣿⣿⣿
			⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡅⠈⡲⠄⠀⡖⠋⠙⠒⣾⠎⠇⢿⣿⣿⣿⣿⣿
			⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⠀⣷⡖⣮⠄⠀⠒⣲⣿⢸⣇⣿⣿⣿⣿⣿⣿
			⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡆⢻⣿⠛⠉⠉⢲⣯⠇⣸⣿⣿⣿⣿⣿⣿⣿
			⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⠒⠛⠓⡦⢶⠟⠛⢒⣿⣿⣿⣿⣿⣿⣿⣿
			⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣆⠀⠀⠀⠀⠑⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿
			⣿⣿⣿⣿⠟⠋⠉⢹⠿⣿⣿⣿⠛⢿⣿⡷⠀⠒⠀⢾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
			⣿⣿⣿⣧⡀⠀⠀⢾⡀⠉⠉⠁⠀⠀⢀⣀⠀⠀⠀⠈⠁⠀⢻⡿⠟⠛⠟⠛⠉⠛
			⣿⡿⠛⠛⠉⠂⠀⠀⠀⢀⡂⠓⠄⠀⠈⠉⠗⣉⠈⠁⠂⠀⠀⠈⠁⠀⠀⠀⠀⠀
			⣿⡀⠀⠈⠆⠀⠀⠀⠀⠀⠉⠘⠐⠿⠿⠁⠀⠘⠃⠤⠤⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		`)
		response := &pb.AttackResponse{Success: true}
		go func() {
			s.enviarSenalTermino("derrota")
		}()
		return response, nil
	}
	return &pb.AttackResponse{Success: false}, nil
}

func (s *servidorTai) enviarSenalTermino(resultado string) {
	_, err := s.primaryNodeClient.SendTerminationSignal(s.ctx, &pb.TerminateProcess{Result: resultado})
	if err != nil {
		fmt.Printf("Error al enviar señal de término al Nodo Primario: %v\n", err)
	} else {
		fmt.Println("Señal de término enviada al Nodo Primario")
	}
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
	VI, err := leerVI()
	if err != nil {
		log.Fatalf("Error al leer input.txt: %v", err)
	}

	s := &servidorTai{
		vida:         VI,
		terminarChan: make(chan struct{}),
	}
	s.ctx, s.cancelar = context.WithCancel(context.Background())

	go func() {
		lis, err := net.Listen("tcp", ":50058")
		if err != nil {
			log.Fatalf("No se pudo escuchar en el puerto 50058: %v", err)
		}
		servidorGRPC := grpc.NewServer()
		pb.RegisterTaiServiceServer(servidorGRPC, s)
		if err := servidorGRPC.Serve(lis); err != nil {
			log.Fatalf("No se pudo servir: %v", err)
		}
	}()

	connDiaboromon, err := grpc.Dial("diaboromon_container:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con Diaboromon: %v", err)
	}
	defer connDiaboromon.Close()
	s.diaboromonClient = pb.NewDiaboromonServiceClient(connDiaboromon)

	connPrimaryNode, err := grpc.Dial("primary_node_container:50057", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el Nodo Primario: %v", err)
	}
	defer connPrimaryNode.Close()
	s.primaryNodeClient = pb.NewPrimaryNodeServiceClient(connPrimaryNode)

	_, err = s.diaboromonClient.StartDiaboromon(s.ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error al iniciar Diaboromon: %v", err)
	}
	fmt.Println("Tai ha enviado la señal de inicio a Diaboromon.")

	go func() {
		<-s.terminarChan
		fmt.Println("Nodo Tai: Terminando ejecución.")
		os.Exit(0)
	}()

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
			fmt.Println("Recolectando información de los Digimons...")
			respuesta, err := s.primaryNodeClient.GetAttackData(s.ctx, &pb.TaiRequest{Message: "Solicitud de datos para ataque"})
			if err != nil {
				fmt.Printf("Error al obtener datos del Nodo Primario: %v\n", err)
				continue
			}
			datosRecolectados := respuesta.DataCollected

			fmt.Println("Nodo Tai atacó a Diaboromon")
			ataqueRespuesta, err := s.diaboromonClient.AttackDiaboromon(s.ctx, &pb.AttackRequest{AttackValue: datosRecolectados})
			if err != nil {
				fmt.Printf("Error al atacar a Diaboromon: %v\n", err)
				continue
			}
			if ataqueRespuesta.Success {
				fmt.Println(`
					██╗   ██╗██╗ ██████╗████████╗ ██████╗ ██████╗ ██╗ █████╗ 
					██║   ██║██║██╔════╝╚══██╔══╝██╔═══██╗██╔══██╗██║██╔══██╗
					██║   ██║██║██║        ██║   ██║   ██║██████╔╝██║███████║
					╚██╗ ██╔╝██║██║        ██║   ██║   ██║██╔══██╗██║██╔══██║
					 ╚████╔╝ ██║╚██████╗   ██║   ╚██████╔╝██║  ██║██║██║  ██║
					  ╚═══╝  ╚═╝ ╚═════╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝╚═╝╚═╝  ╚═╝
					⣘⠿⣿⣿⣿⣿⣿⣿⠋⠉⠀⠹⡌⢧⠀⠀⠀⠀⠀⠘⡦⣄⣀⡀⠀⠀⠀⠀⣠⣾⠖⠋⠀⠀⣠⣾⣿⣿⣿⣿⣿⣿
					⢹⡦⣄⡉⠛⠿⣿⠁⠀⠀⠀⠀⢹⡜⣆⠀⠀⣀⡤⠖⠋⠀⠈⠉⠓⠦⣄⡀⠛⢳⣶⠤⠴⠞⣿⣿⣿⣿⣿⣿⣿⣿
					⢸⣷⣮⣙⠶⣤⡈⠳⡄⠀⣠⠞⠁⠙⣼⣴⠮⠍⠀⠀⠀⠀⠠⠤⣤⣀⡀⠉⠳⣄⣨⡿⠓⡄⠸⣿⣿⣿⣿⣿⣿⣿
					⢸⠿⠿⠛⢳⡶⠭⠷⢿⡜⠁⢀⡴⠞⢛⣠⠴⠚⠉⠀⠀⠀⠀⠀⠀⣀⣉⣟⣲⣦⣝⠛⠷⢦⣄⢻⣿⣿⣿⣿⣿⣿
					⢸⢀⣀⣀⣸⠁⢻⣯⠟⢀⡴⣫⠴⠚⠉⠀⠀⠀⠀⠀⠀⠀⢀⡴⠊⢹⡟⢿⣿⣿⠻⡇⠈⠛⠻⢮⣿⣿⣿⡿⠻⢻
					⢸⣿⣿⣿⣿⢀⣾⢏⣴⣷⡯⣅⡀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⢧⣤⣾⣿⡆⢿⣿⡆⡇⠀⠀⠀⠀⠈⢻⣴⠀⣀⣼
					⢸⣿⣿⣿⣿⠚⢨⣿⣿⣏⢀⣠⣿⡷⣤⣀⠀⢀⣀⣴⣿⣿⣿⣇⠙⠛⠋⢀⣾⣿⡇⡇⠀⠀⠀⠀⠀⠀⠙⢿⣿⣿
					⢸⣿⣿⣿⣿⠀⢸⣿⣿⡌⠿⣿⡿⢳⡿⠛⢻⣩⠏⠉⠀⠈⠉⠙⠿⣶⡾⢽⣿⣿⡴⠃⠀⠀⣠⡤⠖⠒⠛⠋⠉⠉
					⢸⣿⣿⣿⡟⠀⠀⢿⣿⣿⣤⣤⠞⠉⠀⠀⠸⠋⢀⣀⡤⠤⠤⣤⣀⣀⠉⠡⠈⠉⠀⠀⣠⣴⡏⠀⠀⠀⠀⠀⠀⠀
					⢸⢛⣿⡟⠀⠀⠀⠈⠛⠿⠟⢁⡤⠖⠒⠒⢲⡞⠋⠀⠀⠀⠀⢀⣠⢦⠩⡝⠓⠲⣖⡋⢡⡿⠀⠀⠀⠀⠀⠀⠀⠀
					⢸⣼⣿⠀⠀⣠⠶⠺⡗⠒⠚⠙⣦⣤⣀⠀⠀⣷⠀⠀⠀⣀⡴⣿⠁⢠⠆⣹⡞⠋⣹⣷⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀
					⡼⣿⡟⢀⣞⡵⠖⠶⣿⠗⢸⣄⢻⢸⡟⠙⠶⢼⣦⠴⠋⠁⠀⠈⢦⣸⡴⢳⡇⠀⢿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀
					⡇⠈⣿⣾⣯⣄⡉⠀⢹⣷⡈⢿⡷⣼⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣻⠀⢸⠃⠀⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀
					⡇⠀⣿⣧⣭⣻⣿⣷⣶⣿⣷⡀⢳⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⠇⠀⣿⠀⠀⣿⣿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀
					⠃⠀⢸⣿⣧⣿⣿⣿⣿⣿⣿⣷⠀⢧⠀⠀⠀⢀⣤⢤⡀⠀⠀⢀⡟⠀⠀⠟⠀⠀⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
					⠀⠀⠈⡟⣏⣉⠙⠛⠛⢿⣿⣿⡆⠘⡆⠀⢠⡎⣼⠀⠙⢦⣠⣞⣀⣀⠀⠀⠀⠀⣿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
					⠀⠀⠀⢷⢹⡀⠀⠀⣴⣿⣿⣿⣧⠀⠹⠴⣏⠀⡇⠀⠀⢀⠟⠉⠉⠉⠁⠀⠀⢀⡏⠀⠀⠀⠀⠀⠀⠀⢀⡀⠀⠀
					⠀⠀⠀⢸⠈⣧⠀⠀⠈⢹⣿⣿⣿⡀⠀⠀⠘⢦⡂⣀⣴⠋⠀⠀⣀⣤⣴⣶⣶⣿⠃⠀⠀⠀⠀⣀⡤⠖⠉⠀⠀⠀
					⠀⠀⠀⠸⡇⢹⡄⠀⢶⣿⣿⣿⣿⣿⣷⣦⡀⠀⠉⢱⠃⣠⣴⣿⣿⣿⣿⣿⣿⡏⠀⠀⣠⠖⠋⠁⠀⠀⠀⠀⠀⠀
					⠀⠀⠀⠀⢻⣌⠳⢦⣌⠙⣿⣿⣿⣿⣿⣿⣿⣶⣤⣉⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⠖⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
					⠀⠀⠀⠀⠀⠙⠻⣶⣽⣷⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⣤⣤⠴⠒⠒⠒⠒⠒⠒⠀⠀
				`)
				go func() {
					s.enviarSenalTermino("victoria")
				}()
				<-s.terminarChan
				os.Exit(0)
			} else {
				fmt.Println("Diaboromon repelió el ataque")
				s.mu.Lock()
				s.vida -= 10
				fmt.Println("Vida restante: ", s.vida)
				if s.vida <= 0 {
					s.mu.Unlock()
					fmt.Println("Diaboromon ha ganado")
					go func() {
						s.enviarSenalTermino("derrota")
					}()
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

	<-s.terminarChan
	os.Exit(0)
}
