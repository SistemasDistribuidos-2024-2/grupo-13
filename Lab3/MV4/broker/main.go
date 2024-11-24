package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	pb "grpc-server/proto/grpc-server/proto"

	"google.golang.org/grpc"
)

const (
	port          = ":50050"
	arbitraryAddr = "dist049:50054"
)

var (
	serverAddresses = []string{"dist049:50054", "dist050:50055", "dist051:50056"}
	mutex           sync.Mutex
)

type brokerServer struct {
	pb.UnimplementedHextechServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor en el puerto %s: %v", port, err)
	}

	s := grpc.NewServer()
	pb.RegisterHextechServiceServer(s, &brokerServer{})

	log.Printf("Broker escuchando en %s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
	}
}

func getRandomServer() string {
	rand.Seed(time.Now().UnixNano())
	return serverAddresses[rand.Intn(len(serverAddresses))]
}

func (b *brokerServer) AddProductBroker(ctx context.Context, req *pb.AddProductRequest) (*pb.AddressResponse, error) {
	log.Printf("Solicitud recibida: AgregarProducto %s %s [%d]\n", req.Region, req.Product, req.Quantity)
	address := getRandomServer()
	log.Printf("Dirección retornada: %s\n", address)
	return &pb.AddressResponse{Address: address}, nil
}

func (b *brokerServer) RenameProductBroker(ctx context.Context, req *pb.RenameProductRequest) (*pb.AddressResponse, error) {
	log.Printf("Solicitud recibida: RenombrarProducto %s %s -> %s\n", req.Region, req.OldProduct, req.NewProduct)
	address := getRandomServer()
	log.Printf("Dirección retornada: %s\n", address)
	return &pb.AddressResponse{Address: address}, nil
}

func (b *brokerServer) UpdateProductBroker(ctx context.Context, req *pb.UpdateProductRequest) (*pb.AddressResponse, error) {
	log.Printf("Solicitud recibida: ActualizarValor %s %s [%d]\n", req.Region, req.Product, req.Quantity)
	address := getRandomServer()
	log.Printf("Dirección retornada: %s\n", address)
	return &pb.AddressResponse{Address: address}, nil
}

func (b *brokerServer) DeleteProductBroker(ctx context.Context, req *pb.DeleteProductRequest) (*pb.AddressResponse, error) {
	log.Printf("Solicitud recibida: BorrarProducto %s %s\n", req.Region, req.Product)
	address := getRandomServer()
	log.Printf("Dirección retornada: %s\n", address)
	return &pb.AddressResponse{Address: address}, nil
}

func (b *brokerServer) GetProductBroker(ctx context.Context, req *pb.GetProductRequest) (*pb.AddressResponse, error) {
	log.Printf("Solicitud recibida: ObtenerProducto %s %s\n", req.Region, req.Product)
	address := getRandomServer()
	log.Printf("Dirección retornada: %s\n", address)
	return &pb.AddressResponse{Address: address}, nil
}

func (b *brokerServer) ForceMerge(ctx context.Context, req *pb.ErrorMergeRequest) (*pb.ConfirmationError, error) {
	log.Printf("Error recibido: Forzar Merge %s %v\n", req.Region, req.VectorClock)

	confirmation := &pb.ConfirmationError{
		Confirmation: "Merge forzado recibido por el Broker",
	}
	log.Println("Confirmación enviada al cliente.")

	log.Printf("Reenviando a servidor arbitrario: %s\n", arbitraryAddr)
	sendToArbitraryServer(req)

	return confirmation, nil
}

func sendToArbitraryServer(req *pb.ErrorMergeRequest) {
	conn, err := grpc.Dial(arbitraryAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar con el servidor arbitrario: %v", err)
	}
	defer conn.Close()

	client := pb.NewHextechServiceClient(conn)
	_, err = client.ForceMerge(context.Background(), req)
	if err != nil {
		log.Fatalf("Error al reenviar la solicitud al servidor arbitrario: %v", err)
	}

	log.Println("Servidor arbitrario confirmó el merge forzado.")
}
