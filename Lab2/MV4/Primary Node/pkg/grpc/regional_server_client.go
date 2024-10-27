package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "primary_node/pkg/grpc/protobuf"
)

// Cliente para comunicarse con múltiples servidores regionales
type RegionalServerClients struct {
	clients []pb.RegionalServerServiceClient
}

// Crear clientes gRPC para servidores regionales
func NewRegionalServerClients(addresses []string) (*RegionalServerClients, error) {
	var clients []pb.RegionalServerServiceClient

	for _, addr := range addresses {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Printf("[ERROR] No se pudo conectar al servidor regional en %s: %v", addr, err)
			return nil, err
		}
		client := pb.NewRegionalServerServiceClient(conn)
		clients = append(clients, client)
		log.Printf("[PRIMARY NODE] Conectado al servidor regional en %s", addr)
	}

	return &RegionalServerClients{clients: clients}, nil
}

// Enviar señal de terminación a todos los servidores regionales
func (c *RegionalServerClients) TerminateAll() error {
	for _, client := range c.clients {
		_, err := client.TerminateRegional(context.Background(), &pb.TerminateRequest{Message: "Fin de ejecución"})
		if err != nil {
			log.Printf("[ERROR] Error al enviar señal de terminación: %v", err)
			return err
		}
		log.Println("[PRIMARY NODE] Señal de terminación enviada a un servidor regional.")
	}
	return nil
}
