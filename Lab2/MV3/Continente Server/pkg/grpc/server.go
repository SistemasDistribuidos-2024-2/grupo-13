package grpc

import (
	"context"
	"log"
	"net"

	pb "regional_server/pkg/grpc/protobuf"

	"google.golang.org/grpc"
)

// Servidor Regional Server
type RegionalServer struct {
	pb.UnimplementedRegionalServerServiceServer
	QuitChan chan bool // Canal para comunicar la terminación
}

// NewRegionalServer crea una nueva instancia del servidor regional
func NewRegionalServer() *RegionalServer {
	return &RegionalServer{
		QuitChan: make(chan bool), // Inicializa el canal
	}
}

// TerminateRegional maneja la señal de terminación del Primary Node
func (s *RegionalServer) TerminateRegional(ctx context.Context, req *pb.TerminateRequest) (*pb.TerminateResponse, error) {
	log.Printf("[CONTINENTE SERVER] Señal de terminación recibida: %s", req.Message)
	s.QuitChan <- true // Notifica la terminación al proceso principal
	return &pb.TerminateResponse{Message: "Servidor Regional terminado correctamente"}, nil
}

// StartServer inicia el servidor gRPC del Regional Server
func StartServer(address string, server *RegionalServer) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("[ERROR][CONTINENTE SERVER] Error al iniciar el servidor: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRegionalServerServiceServer(grpcServer, server)

	log.Printf("[CONTINENTE SERVER] Servidor escuchando en %s", address)
	return grpcServer.Serve(lis)
}
