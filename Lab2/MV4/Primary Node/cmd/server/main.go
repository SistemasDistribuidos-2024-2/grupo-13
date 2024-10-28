package main

import (
	"log"
	"primary_node/pkg/grpc"
	"time"
)

const (
	primaryNodeAddress  = "primary_node_container:50057"
	dataNode1Address    = "data_node_1_container:50051"
	dataNode2Address    = "data_node_2_container:50052"
	taiNodeAddress      = "tai_container:50058"
	regionalServer1Addr = "isla_file_container:50053"
	regionalServer2Addr = "continente_folder_container:50054"
	regionalServer3Addr = "continente_server_container:50056"
)

func main() {
	// Crear los clientes gRPC requeridos
	dataNodeClient, err := grpc.NewDataNodeClient(dataNode1Address, dataNode2Address)
	if err != nil {
		log.Fatalf("[ERROR] Error al conectar con los Data Nodes: %v", err)
	}

	taiClient, err := grpc.NewTaiClient(taiNodeAddress)
	if err != nil {
		log.Fatalf("[ERROR] Error al conectar con el Nodo Tai: %v", err)
	}

	regionalClients, err := grpc.NewRegionalServerClients([]string{
		regionalServer1Addr, regionalServer2Addr, regionalServer3Addr,
	})
	if err != nil {
		log.Fatalf("[ERROR] Error al conectar con los servidores regionales: %v", err)
	}

	// Crear e iniciar el servidor Primary Node
	server, err := grpc.NewPrimaryNodeServer(dataNodeClient, taiClient, regionalClients)
	if err != nil {
		log.Fatalf("[ERROR] Error al iniciar el Primary Node: %v", err)
	}

	// Iniciar el servidor gRPC en una goroutine
	go func() {
		log.Println("[PRIMARY NODE] Servidor en funcionamiento, esperando señales...")
		if err := grpc.StartServer(primaryNodeAddress, server); err != nil {
			log.Fatalf("[ERROR] Error al ejecutar el servidor: %v", err)
		}
	}()

	// Esperar 10 segundos antes de que el Primary Node actúe como cliente
	log.Println("[PRIMARY NODE] Esperando 10 segundos antes de iniciar la conexión como cliente...")
	time.Sleep(10 * time.Second)

	// Lógica para actuar como cliente (enviar señales o realizar alguna acción como cliente)
	// Por ejemplo, podemos intentar iniciar alguna interacción con los nodos conectados

	// Esperar la señal de terminación del servidor
	<-server.QuitChan
	log.Println("[PRIMARY NODE] El proceso ha terminado. Finalizando el programa.")
}
