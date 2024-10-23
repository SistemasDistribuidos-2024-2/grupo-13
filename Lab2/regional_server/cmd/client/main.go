package main

import (
	"log"
	"time"

	"regional_server/pkg/data"
	"regional_server/pkg/grpc"
)

const primaryNodeAddress = "localhost:50051"

func main() {
	// Carga de configuraciones
	config, err := data.LoadInputConfig()
	if err != nil {
		log.Fatalf("Error al cargar configuración: %v", err)
	}
	// Carga de informacion de digimons
	digimons, err := data.LoadDigimons()
	if err != nil {
		log.Fatalf("Error al cargar Digimons: %v", err)
	}

	// Crear cliente gRPC
	client, err := grpc.NewClient(primaryNodeAddress, config, digimons)
	if err != nil {
		log.Fatalf("Error al conectar con Primary Node: %v", err)
	}
	defer client.Close()

	// Envia 6 mensajes aleatorios al iniciar el programa
	client.SendRandomData(6)

	// Envia mensajes aleatorios en intervalos de TE segundos
	ticker := time.NewTicker(time.Duration(config.TE) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		client.SendRandomData(1)
	}
}
