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
	"time"

	"google.golang.org/grpc"
	pb "grpc-server/proto/grpc-server/proto"
)

type servidor struct {
	pb.UnimplementedDiaboromonServiceServer
	CD       int32
	TD       int32
	iniciado bool
	mu       sync.Mutex
	derrota  chan struct{}
	victoria chan struct{}
	inicio   chan struct{}
	ctx      context.Context
	cancelar context.CancelFunc
}

func (s *servidor) StartDiaboromon(ctx context.Context, req *pb.Empty) (*pb.StartResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.iniciado {
		s.iniciado = true
		fmt.Println("Diaboromon se está preparando para atacar...")
		s.inicio <- struct{}{}
	}
	return &pb.StartResponse{Message: "Diaboromon iniciado"}, nil
}

func (s *servidor) AttackDiaboromon(ctx context.Context, req *pb.AttackRequest) (*pb.AttackResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if req.AttackValue >= s.CD {
		fmt.Println(`
			██████╗ ███████╗██████╗ ██████╗  ██████╗ ████████╗ █████╗ 
			██╔══██╗██╔════╝██╔══██╗██╔══██╗██╔═══██╗╚══██╔══╝██╔══██╗
			██║  ██║█████╗  ██████╔╝██████╔╝██║   ██║   ██║   ███████║
			██║  ██║██╔══╝  ██╔══██╗██╔══██╗██║   ██║   ██║   ██╔══██║
			██████╔╝███████╗██║  ██║██║  ██║╚██████╔╝   ██║   ██║  ██║
			╚═════╝ ╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝    ╚═╝   ╚═╝  ╚═╝
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
		s.derrota <- struct{}{}
		s.cancelar()
		return &pb.AttackResponse{Success: true}, nil
	} else {
		fmt.Println("Diaboromon ha repelido el ataque.")
		return &pb.AttackResponse{Success: false}, nil
	}
}

func leerCDyTD() (int32, int32, error) {
	archivo, err := os.Open("input.txt")
	if err != nil {
		return 0, 0, err
	}
	defer func(archivo *os.File) {
		err := archivo.Close()
		if err != nil {

		}
	}(archivo)

	scanner := bufio.NewScanner(archivo)
	var CD int32
	var TD int32
	if scanner.Scan() {
		linea := scanner.Text()
		valores := strings.Split(linea, ",")
		for i := range valores {
			valores[i] = strings.TrimSpace(valores[i])
		}
		if len(valores) < 5 {
			return 0, 0, fmt.Errorf("input.txt no tiene suficientes valores")
		}
		tdFloat, err := strconv.ParseFloat(valores[2], 64)
		if err != nil {
			return 0, 0, err
		}
		TD = int32(tdFloat)
		cdFloat, err := strconv.ParseFloat(valores[3], 64)
		if err != nil {
			return 0, 0, err
		}
		CD = int32(cdFloat)
		return CD, TD, nil
	} else {
		return 0, 0, fmt.Errorf("input.txt está vacío")
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, err
	}
	return 0, 0, fmt.Errorf("No se pudo leer input.txt")
}

func main() {
	CD, TD, err := leerCDyTD()
	if err != nil {
		log.Fatalf("Error al leer input.txt: %v", err)
	}

	s := &servidor{
		CD:       CD,
		TD:       TD,
		derrota:  make(chan struct{}),
		victoria: make(chan struct{}),
		inicio:   make(chan struct{}),
	}
	s.ctx, s.cancelar = context.WithCancel(context.Background())

	go func() {
		lis, err := net.Listen("tcp", ":50055")
		if err != nil {
			log.Fatalf("No se pudo escuchar: %v", err)
		}
		servidorGRPC := grpc.NewServer()
		pb.RegisterDiaboromonServiceServer(servidorGRPC, s)
		if err := servidorGRPC.Serve(lis); err != nil {
			log.Fatalf("No se pudo servir: %v", err)
		}
	}()

	<-s.inicio

	go func() {
		conn, err := grpc.Dial("tai_container:50058", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("No se pudo conectar con Tai: %v", err)
		}
		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)

		cliente := pb.NewTaiServiceClient(conn)
		ticker := time.NewTicker(time.Duration(s.TD) * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-s.ctx.Done():
				return
			case <-ticker.C:
				fmt.Println("Diaboromon ha atacado.")
				res, err := cliente.DiaboromonAttack(s.ctx, &pb.Empty{})
				if err != nil {
					log.Fatalf("Error al atacar a Tai: %v", err)
				}
				if res.Success {
					fmt.Println(`
						██╗   ██╗██╗ ██████╗████████╗ ██████╗ ██████╗ ██╗ █████╗ 
						██║   ██║██║██╔════╝╚══██╔══╝██╔═══██╗██╔══██╗██║██╔══██╗
						██║   ██║██║██║        ██║   ██║   ██║██████╔╝██║███████║
						╚██╗ ██╔╝██║██║        ██║   ██║   ██║██╔══██╗██║██╔══██║
						 ╚████╔╝ ██║╚██████╗   ██║   ╚██████╔╝██║  ██║██║██║  ██║
						  ╚═══╝  ╚═╝ ╚═════╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝╚═╝╚═╝  ╚═╝
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
					s.victoria <- struct{}{}
					s.cancelar()
					return
				}
			}
		}
	}()

	select {
	case <-s.derrota:
	case <-s.victoria:
	}
	time.Sleep(1 * time.Second)
}
