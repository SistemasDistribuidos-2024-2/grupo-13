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
	port          = ":50050"                   // Puerto donde el broker escucha
	arbitraryAddr = "hextech1_container:50054" // Dirección del servidor arbitrario
)

var (
	serverAddresses = []string{"hextech1_container:50054", "hextech2_container:50055", "hextech3_container:50056"}
	mutex           sync.Mutex
)

type brokerServer struct {
	pb.UnimplementedHextechServiceServer
}

func main() {
	// Inicializar el servidor gRPC
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

// getRandomServer selecciona una dirección aleatoria de los servidores
func getRandomServer() string {
	rand.Seed(time.Now().UnixNano())
	return serverAddresses[rand.Intn(len(serverAddresses))]
}

// AddProductBroker maneja la solicitud AgregarProducto y retorna una dirección aleatoria
func (b *brokerServer) AddProductBroker(ctx context.Context, req *pb.AddProductRequest) (*pb.AddressResponse, error) {
	log.Printf("Solicitud recibida: AgregarProducto %s %s [%d]\n", req.Region, req.Product, req.Quantity)
	address := getRandomServer()
	log.Printf("Dirección retornada: %s\n", address)
	return &pb.AddressResponse{Address: address}, nil
}

// RenameProductBroker maneja la solicitud RenombrarProducto y retorna una dirección aleatoria
func (b *brokerServer) RenameProductBroker(ctx context.Context, req *pb.RenameProductRequest) (*pb.AddressResponse, error) {
	log.Printf("Solicitud recibida: RenombrarProducto %s %s -> %s\n", req.Region, req.OldProduct, req.NewProduct)
	address := getRandomServer()
	log.Printf("Dirección retornada: %s\n", address)
	return &pb.AddressResponse{Address: address}, nil
}

// UpdateProductBroker maneja la solicitud ActualizarValor y retorna una dirección aleatoria
func (b *brokerServer) UpdateProductBroker(ctx context.Context, req *pb.UpdateProductRequest) (*pb.AddressResponse, error) {
	log.Printf("Solicitud recibida: ActualizarValor %s %s [%d]\n", req.Region, req.Product, req.Quantity)
	address := getRandomServer()
	log.Printf("Dirección retornada: %s\n", address)
	return &pb.AddressResponse{Address: address}, nil
}

// DeleteProductBroker maneja la solicitud BorrarProducto y retorna una dirección aleatoria
func (b *brokerServer) DeleteProductBroker(ctx context.Context, req *pb.DeleteProductRequest) (*pb.AddressResponse, error) {
	log.Printf("Solicitud recibida: BorrarProducto %s %s\n", req.Region, req.Product)
	address := getRandomServer()
	log.Printf("Dirección retornada: %s\n", address)
	return &pb.AddressResponse{Address: address}, nil
}

// GetProductBroker maneja la solicitud ObtenerProducto y retorna una dirección aleatoria
func (b *brokerServer) GetProductBroker(ctx context.Context, req *pb.GetProductRequest) (*pb.AddressResponse, error) {
	log.Printf("Solicitud recibida: ObtenerProducto %s %s\n", req.Region, req.Product)
	address := getRandomServer()
	log.Printf("Dirección retornada: %s\n", address)
	return &pb.AddressResponse{Address: address}, nil
}

// ForceMerge maneja solicitudes de error y las reenvía al servidor arbitrario
func (b *brokerServer) ForceMerge(ctx context.Context, req *pb.ErrorMergeRequest) (*pb.ConfirmationError, error) {
	log.Printf("Error recibido: Forzar Merge %s %v\n", req.Region, req.VectorClock)

	// Confirmación al cliente
	confirmation := &pb.ConfirmationError{
		Confirmation: "Merge forzado recibido por el Broker",
	}
	log.Println("Confirmación enviada al cliente.")

	// Reenviar al servidor arbitrario
	log.Printf("Reenviando a servidor arbitrario: %s\n", arbitraryAddr)
	sendToArbitraryServer(req)

	return confirmation, nil
}

// sendToArbitraryServer reenvía el error al servidor arbitrario
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
