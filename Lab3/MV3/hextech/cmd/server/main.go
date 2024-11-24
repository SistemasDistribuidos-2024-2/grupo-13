package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"hextech/internal/server"
	"hextech/proto"
	"google.golang.org/grpc"
)

func main() {
	// Definir flags para los argumentos
	serverID := flag.Int("id", 0, "ID del servidor Hextech (obligatorio)")
	port := flag.String("port", "", "Puerto donde el servidor Hextech se ejecutará (obligatorio)")
	peerPorts := flag.String("peers", "", "Puertos de los peers separados por comas (opcional)")

	// Parsear los flags
	flag.Parse()

	// Validar los argumentos
	if *serverID < 0 {
		log.Fatalf("SERVER_ID debe ser un valor positivo. Valor recibido: %d", *serverID)
	}

	if *port == "" {
		log.Fatalf("PORT debe especificarse y no puede estar vacío")
	}

	if *peerPorts == "" {
		log.Printf("[Advertencia] No se especificaron peers. El servidor funcionará de manera aislada.")
	}

	// Crear el servidor Hextech
	hexServer := server.NewHextechServer(*serverID)

	// Configurar el servidor gRPC
	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("No se pudo iniciar el listener en el puerto %s: %v", *port, err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterHextechServiceServer(grpcServer, hexServer)

	// Conectar a los peers si se especificaron
	if *peerPorts != "" {
		peers := strings.Split(*peerPorts, ",") // Dividir la cadena en una lista
		for _, peerPort := range peers {
			if peerPort != *port { // Evitar conectarse a sí mismo
				conn, err := grpc.Dial("localhost:"+peerPort, grpc.WithInsecure())
				if err == nil {
					hexServer.AddPeer(proto.NewHextechServiceClient(conn))
					fmt.Printf("[Servidor Hextech] Conectado al peer en el puerto [%s]\n", peerPort)
				} else {
					fmt.Printf("[Servidor Hextech] Error al conectar con el peer en el puerto [%s]: %v\n", peerPort, err)
				}
			}
		}
	}

	// Iniciar la propagación en segundo plano
	go hexServer.StartPropagation()

	// Iniciar el servidor gRPC
	fmt.Printf("[Servidor Hextech] Ejecutándose en el puerto [%s] con ID [%d]...\n", *port, *serverID)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
