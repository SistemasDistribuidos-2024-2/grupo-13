package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "primary_node/pkg/grpc/protobuf"
)

// Cliente gRPC para comunicarse con los Data Nodes
type DataNodeClient struct {
	dataNode1 pb.DataNodeServiceClient
	dataNode2 pb.DataNodeServiceClient
}

// Crear una nueva instancia del cliente gRPC
func NewDataNodeClient(node1Addr, node2Addr string) (*DataNodeClient, error) {
	conn1, err := grpc.Dial(node1Addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	conn2, err := grpc.Dial(node2Addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &DataNodeClient{
		dataNode1: pb.NewDataNodeServiceClient(conn1),
		dataNode2: pb.NewDataNodeServiceClient(conn2),
	}, nil
}

// Enviar el atributo del Digimon al Data Node correspondiente
func (c *DataNodeClient) SendToDataNode(id int, name, attribute string) error {
	var client pb.DataNodeServiceClient

	if name[0] >= 'A' && name[0] <= 'M' {
		client = c.dataNode1
	} else {
		client = c.dataNode2
	}

	_, err := client.StoreDigimon(context.Background(), &pb.DigimonInfo{
		Id:        int32(id),
		Attribute: attribute,
	})
	if err != nil {
		log.Printf("Error al enviar datos al Data Node: %v", err)
		return err
	}

	log.Printf("Datos enviados al Data Node: ID=%d, Atributo=%s", id, attribute)
	return nil
}

// Obtener el atributo del Digimon desde el Data Node
func (c *DataNodeClient) GetAttributeFromDataNode(id int) (string, error) {
	req := &pb.DigimonRequest{Id: int32(id)}

	resp, err := c.dataNode1.GetDigimonAttribute(context.Background(), req)
	if err == nil {
		return resp.Attribute, nil
	}

	resp, err = c.dataNode2.GetDigimonAttribute(context.Background(), req)
	if err != nil {
		return "", err
	}
	return resp.Attribute, nil
}
