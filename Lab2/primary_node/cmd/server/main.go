package main

import (
	"log"
	"primary_node/pkg/grpc"
)

const (
	primaryNodeAddress = "localhost:50051"
	dataNode1Address   = "localhost:50052"
	dataNode2Address   = "localhost:50053"
	taiNodeAddress     = "localhost:50054"
)

func main() {
	// Crear cliente gRPC para los Data Nodes
	dataNodeClient, err := grpc.NewDataNodeClient(dataNode1Address, dataNode2Address)
	if err != nil {
		log.Fatalf("Error al conectar con los Data Nodes: %v", err)
	}

	// Crear cliente gRPC para el Nodo Tai
	taiClient, err := grpc.NewTaiClient(taiNodeAddress)
	if err != nil {
		log.Fatalf("Error al conectar con el Nodo Tai: %v", err)
	}

	// Crear e iniciar el servidor Primary Node
	server, err := grpc.NewPrimaryNodeServer(dataNodeClient)
	if err != nil {
		log.Fatalf("Error al iniciar el Primary Node: %v", err)
	}

	// Ejemplo de envío de datos al Nodo Tai
	if err := taiClient.SendDataAmount(42); err != nil {
		log.Fatalf("Error al enviar datos al Nodo Tai: %v", err)
	}

	// Iniciar el servidor gRPC
	grpc.StartServer(primaryNodeAddress, server)
}
