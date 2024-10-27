package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"google.golang.org/grpc"
	pb "grpc-server/proto/grpc-server/proto"
)

type server struct {
	pb.UnimplementedDataNodeServiceServer
	mu             sync.Mutex
	dataNodeNumber int
	filename       string
	grpcServer     *grpc.Server
}

func (s *server) StoreDigimon(ctx context.Context, in *pb.DigimonInfo) (*pb.StoreDigimonResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Printf("[DATA NODE %d] Recibida informacion de primary node: ID=%d.\n", s.dataNodeNumber, in.Id)
	file, err := os.OpenFile(s.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error al abrir el archivo: %v", err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	_, err = file.WriteString(fmt.Sprintf("%d,%s\n", in.Id, in.Attribute))
	if err != nil {
		log.Printf("Error al escribir en el archivo: %v", err)
		return nil, err
	}
	fmt.Printf("[DATA NODE %d] Digimon ID=%d escritura completa. \n", s.dataNodeNumber, in.Id)
	return &pb.StoreDigimonResponse{Message: "Success"}, nil
}

func (s *server) GetDigimonAttribute(ctx context.Context, in *pb.DigimonRequest) (*pb.DigimonResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Printf("[DATA NODE %d] Solicitud de Primary Node recibida: ID=%d.\n", s.dataNodeNumber, in.Id)

	file, err := os.Open(s.filename)
	if err != nil {
		log.Printf("Error al abrir el archivo: %v", err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) == 2 {
			idStr := parts[0]
			attribute := parts[1]

			id, err := strconv.Atoi(idStr)
			if err != nil {
				continue
			}
			if int32(id) == in.Id {
				fmt.Printf("[DATA NODE %d] Respuesta enviada al Primary Node: Atributo=%s.\n", s.dataNodeNumber, attribute)
				return &pb.DigimonResponse{Attribute: attribute}, nil
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error al leer el archivo: %v", err)
		return nil, err
	}

	fmt.Printf("[DATA NODE %d] Respuesta enviada al Primary Node: Atributo no encontrado.\n", s.dataNodeNumber)
	return nil, fmt.Errorf("Atributo no encontrado para ID %d", in.Id)
}

func (s *server) Terminate(ctx context.Context, in *pb.TerminateRequest) (*pb.TerminateResponse, error) {
	fmt.Printf("[DATA NODE %d] Señal de terminación recibida del Primary Node.\n", s.dataNodeNumber)

	go func() {
		s.grpcServer.GracefulStop()
	}()

	return &pb.TerminateResponse{Message: "Data Node terminating"}, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Uso: %s <DataNodeNumber>", os.Args[0])
	}

	dataNodeNumber, err := strconv.Atoi(os.Args[1])
	if err != nil || (dataNodeNumber != 1 && dataNodeNumber != 2) {
		log.Fatalf("Número de Data Node inválido: %s", os.Args[1])
	}

	filename := fmt.Sprintf("DATA_%d.txt", dataNodeNumber)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50050+dataNodeNumber))
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	grpcServer := grpc.NewServer()
	dataNodeServer := &server{
		dataNodeNumber: dataNodeNumber,
		filename:       filename,
		grpcServer:     grpcServer,
	}

	pb.RegisterDataNodeServiceServer(grpcServer, dataNodeServer)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		fmt.Printf("\n[DATA NODE %d] Terminando ejecución por señal del sistema.\n", dataNodeServer.dataNodeNumber)
		grpcServer.GracefulStop()
	}()

	fmt.Printf("Data Node %d escuchando en el puerto %d\n", dataNodeNumber, 50050+dataNodeNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}

	fmt.Printf("[DATA NODE %d] Ejecución terminada.\n", dataNodeNumber)
}
