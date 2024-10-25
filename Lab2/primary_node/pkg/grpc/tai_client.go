package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "primary_node/pkg/grpc/protobuf"
)

// Cliente gRPC para comunicarse con el Nodo Tai
type TaiClient struct {
	client pb.TaiServiceClient
}

// Crear una nueva instancia del cliente gRPC para el Nodo Tai
func NewTaiClient(taiAddress string) (*TaiClient, error) {
	conn, err := grpc.Dial(taiAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &TaiClient{
		client: pb.NewTaiServiceClient(conn),
	}, nil
}

// Enviar la cantidad de datos acumulados al Nodo Tai
func (c *TaiClient) SendDataAmount(dataAmount int32) error {
	_, err := c.client.ReceiveDataAmount(context.Background(), &pb.AttackDataResponse{
		DataCollected: dataAmount,
	})
	if err != nil {
		log.Printf("Error al enviar datos al Nodo Tai: %v", err)
		return err
	}
	log.Printf("Cantidad de datos %d enviada al Nodo Tai", dataAmount)
	return nil
}
