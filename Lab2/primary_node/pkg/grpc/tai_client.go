package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "primary_node/pkg/grpc/protobuf"
)

// TaiClient es el cliente gRPC para comunicarse con el Nodo Tai
type TaiClient struct {
	client pb.TaiServiceClient
}

// NewTaiClient crea una nueva instancia del cliente gRPC para el Nodo Tai
func NewTaiClient(taiAddress string) (*TaiClient, error) {
	conn, err := grpc.Dial(taiAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &TaiClient{
		client: pb.NewTaiServiceClient(conn),
	}, nil
}

// SendDataAmount envía la cantidad de datos acumulados al Nodo Tai
func (c *TaiClient) SendDataAmount(dataAmount int32) error {
	_, err := c.client.DiaboromonAttack(context.Background(), &pb.Empty{})
	if err != nil {
		log.Printf("Error al iniciar ataque de Diaboromon: %v", err)
		return err
	}
	log.Printf("Se envió la cantidad de datos: %d", dataAmount)
	return nil
}
