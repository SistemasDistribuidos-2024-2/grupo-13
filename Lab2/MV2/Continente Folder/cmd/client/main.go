package main

import (
	"log"
	"time"

	"regional_server/pkg/data"
	"regional_server/pkg/grpc"
)

const (
	primaryNodeAddress    = "dist052:50057"
	regionalServerAddress = "dist050:50054"
)

func main() {
	// Cargar la configuración del input
	config, err := data.LoadInputConfig()
	if err != nil {
		log.Fatalf("Error al cargar configuración: %v", err)
	}

	// Cargar la información de Digimons
	digimons, err := data.LoadDigimons()
	if err != nil {
		log.Fatalf("Error al cargar Digimons: %v", err)
	}

	// Crear e iniciar el servidor regional
	server := grpc.NewRegionalServer()

	go func() {
		if err := grpc.StartServer(regionalServerAddress, server); err != nil {
			log.Fatalf("Error al iniciar el servidor regional: %v", err)
		}
	}()

	// Esperar 10 segundos antes de iniciar el cliente
	log.Println("[CONTINENTE FOLDER] Esperando 10 segundos antes de iniciar el cliente...")
	time.Sleep(10 * time.Second)

	// Crear el cliente gRPC para el Primary Node
	client, err := grpc.NewClient(primaryNodeAddress, config, digimons)
	if err != nil {
		log.Fatalf("Error al conectar con Primary Node: %v", err)
	}
	defer client.Close()

	// Enviar 6 mensajes aleatorios al iniciar el cliente
	go func() {
		client.SendRandomData(6)

		// Configurar un ticker para enviar mensajes en intervalos de TE segundos
		ticker := time.NewTicker(time.Duration(config.TE) * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			client.SendRandomData(1)
		}
	}()

	// Esperar la señal de terminación del servidor regional
	<-server.QuitChan
	log.Println("[CONTINENTE FOLDER] Programa finalizado.")
}
