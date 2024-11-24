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
	serverID := flag.Int("id", 0, "ID del servidor Hextech (obligatorio)")
	port := flag.String("port", "", "Puerto donde el servidor Hextech se ejecutará (obligatorio)")
	peerAddresses := flag.String("peers", "", "Direcciones de los peers en el formato host:port separados por comas (opcional)")

	flag.Parse()

	if *serverID < 0 {
		log.Fatalf("SERVER_ID debe ser un valor positivo. Valor recibido: %d", *serverID)
	}

	if *port == "" {
		log.Fatalf("PORT debe especificarse y no puede estar vacío")
	}

	if *peerAddresses == "" {
		log.Printf("[Advertencia] No se especificaron peers. El servidor funcionará de manera aislada.")
	}

	hexServer := server.NewHextechServer(*serverID)

	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("No se pudo iniciar el listener en el puerto %s: %v", *port, err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterHextechServiceServer(grpcServer, hexServer)

	if *peerAddresses != "" {
		peers := strings.Split(*peerAddresses, ",") 
		for _, peerAddress := range peers {
			if !strings.HasSuffix(peerAddress, ":"+*port) {
				conn, err := grpc.Dial(peerAddress, grpc.WithInsecure())
				if err == nil {
					hexServer.AddPeer(proto.NewHextechServiceClient(conn))
					fmt.Printf("[Servidor Hextech] Conectado al peer en [%s]\n", peerAddress)
				} else {
					fmt.Printf("[Servidor Hextech] Error al conectar con el peer en [%s]: %v\n", peerAddress, err)
				}
			}
		}
	}

	go hexServer.StartPropagation()

	fmt.Printf("[Servidor Hextech] Ejecutándose en el puerto [%s] con ID [%d]...\n", *port, *serverID)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
