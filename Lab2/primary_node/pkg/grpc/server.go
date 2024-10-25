package grpc

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"primary_node/pkg/crypto"
	"primary_node/pkg/data"
	pb "primary_node/pkg/grpc/protobuf"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
)

// Servidor Primary Node
type server struct {
	pb.UnimplementedPrimaryNodeServiceServer
	dataNodeClient *DataNodeClient
}

// Crear un nuevo servidor Primary Node
func NewPrimaryNodeServer(client *DataNodeClient) (*server, error) {
	if err := data.InitInfoFile(); err != nil {
		return nil, fmt.Errorf("error al inicializar INFO.txt: %v", err)
	}
	return &server{dataNodeClient: client}, nil
}

// Generar un ID único
func generateID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100000)
}

// Recibir y procesar mensajes encriptados
func (s *server) ReceiveEncryptedMessage(ctx context.Context, msg *pb.EncryptedMessage) (*pb.Empty, error) {
	decryptedMessage, err := crypto.DecryptAES(msg.EncryptedData)
	if err != nil {
		log.Printf("Error al desencriptar mensaje: %v", err)
		return nil, err
	}

	parts := strings.Split(decryptedMessage, ",")
	if len(parts) != 3 {
		return nil, fmt.Errorf("formato de mensaje incorrecto")
	}

	id := generateID()
	name, attribute, status := parts[0], parts[1], parts[2]
	dataNode := 1
	if name[0] >= 'N' && name[0] <= 'Z' {
		dataNode = 2
	}

	if err := data.WriteInfo(id, dataNode, name, status); err != nil {
		return nil, err
	}

	if err := s.dataNodeClient.SendToDataNode(id, name, attribute); err != nil {
		return nil, err
	}

	log.Printf("Procesado: ID=%d, DataNode=%d, Nombre=%s, Estado=%s", id, dataNode, name, status)
	return &pb.Empty{}, nil
}

// Manejar la solicitud de Tai para obtener los datos acumulados
func (s *server) GetAttackData(ctx context.Context, req *pb.TaiRequest) (*pb.AttackDataResponse, error) {
	totalData := 0.0

	file, err := os.Open("pkg/data/INFO.txt")
	if err != nil {
		return nil, fmt.Errorf("error al abrir INFO.txt: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 4 {
			continue
		}

		id, _ := strconv.Atoi(parts[0])
		status := parts[3]

		if status == "Sacrificado" {
			attribute, err := s.dataNodeClient.GetAttributeFromDataNode(id)
			if err != nil {
				return nil, err
			}
			totalData += calculateDataAmount(attribute)
		}
	}

	return &pb.AttackDataResponse{DataCollected: int32(totalData)}, nil
}

// Calcular la cantidad de datos según el atributo del Digimon
func calculateDataAmount(attribute string) float64 {
	switch attribute {
	case "Vaccine":
		return 3.0
	case "Data":
		return 1.5
	case "Virus":
		return 0.8
	default:
		return 0.0
	}
}

// Iniciar el servidor gRPC
func StartServer(address string, s *server) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("error al iniciar el servidor: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPrimaryNodeServiceServer(grpcServer, s)

	log.Printf("Servidor escuchando en %s", address)
	return grpcServer.Serve(lis)
}
