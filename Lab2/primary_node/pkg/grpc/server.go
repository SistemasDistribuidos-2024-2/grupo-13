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

type Server struct {
	pb.UnimplementedPrimaryNodeServiceServer
	DataNodeClient  *DataNodeClient
	TaiClient       *TaiClient
	RegionalClients *RegionalServerClients
	QuitChan        chan bool // Exportado para acceso desde main.go
}

// NewPrimaryNodeServer crea una nueva instancia del servidor Primary Node
func NewPrimaryNodeServer(dataNodeClient *DataNodeClient, taiClient *TaiClient, regionalClients *RegionalServerClients) (*Server, error) {
	if err := data.InitInfoFile(); err != nil {
		log.Printf("[ERROR][PRIMARY NODE] Error al inicializar INFO.txt: %v", err)
		return nil, err
	}
	return &Server{
		DataNodeClient:  dataNodeClient,
		TaiClient:       taiClient,
		RegionalClients: regionalClients,
		QuitChan:        make(chan bool), // Inicialización del canal
	}, nil
}

// generateID genera un ID único para los Digimon
func generateID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100000)
}

// ReceiveEncryptedMessage procesa mensajes encriptados desde los servidores regionales
func (s *Server) ReceiveEncryptedMessage(ctx context.Context, msg *pb.EncryptedMessage) (*pb.Empty, error) {
	decryptedMessage, err := crypto.DecryptAES(msg.EncryptedData)
	if err != nil {
		log.Printf("[ERROR][PRIMARY NODE] Error al desencriptar mensaje: %v", err)
		return nil, err
	}

	log.Printf("[PRIMARY NODE] Mensaje recibido: %s", decryptedMessage)

	parts := strings.Split(decryptedMessage, ",")
	if len(parts) != 3 {
		log.Printf("[ERROR][PRIMARY NODE] Formato de mensaje incorrecto: %s", decryptedMessage)
		return nil, fmt.Errorf("formato de mensaje incorrecto")
	}

	id := generateID()
	name, attribute, status := parts[0], parts[1], parts[2]

	dataNode := 1
	if name[0] >= 'N' && name[0] <= 'Z' {
		dataNode = 2
	}

	if err := data.WriteInfo(id, dataNode, name, status); err != nil {
		log.Printf("[ERROR][PRIMARY NODE] Error al escribir en INFO.txt: %v", err)
		return nil, err
	}

	if err := s.DataNodeClient.SendToDataNode(id, name, attribute); err != nil {
		log.Printf("[ERROR][PRIMARY NODE] Error al enviar datos al Data Node: %v", err)
		return nil, err
	}

	log.Printf("[PRIMARY NODE] Procesado: ID=%d, DataNode=%d, Nombre=%s, Estado=%s", id, dataNode, name, status)

	return &pb.Empty{}, nil
}

// GetAttackData responde con la cantidad total de datos acumulados
func (s *Server) GetAttackData(ctx context.Context, req *pb.TaiRequest) (*pb.AttackDataResponse, error) {
	log.Println("[PRIMARY NODE] Solicitud de cantidad de datos acumulados recibida.")

	totalData := s.calculateTotalData()
	response := &pb.AttackDataResponse{DataCollected: int32(totalData)}

	log.Printf("[PRIMARY NODE] Respuesta enviada: %d datos acumulados", response.DataCollected)

	return response, nil
}

// calculateTotalData calcula la cantidad total de datos acumulados
func (s *Server) calculateTotalData() float64 {
	totalData := 0.0

	file, err := os.Open("pkg/data/INFO.txt")
	if err != nil {
		log.Printf("[ERROR][PRIMARY NODE] Error al abrir INFO.txt: %v", err)
		return 0.0
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 4 || parts[3] != "Sacrificado" {
			continue
		}

		id, _ := strconv.Atoi(parts[0])
		attribute, err := s.DataNodeClient.GetAttributeFromDataNode(id)
		if err != nil {
			log.Printf("[ERROR][PRIMARY NODE] Error al obtener atributo del Data Node: %v", err)
			continue
		}

		dataAmount := calculateDataAmount(attribute)
		log.Printf("[PRIMARY NODE] Atributo %s -> %f datos", attribute, dataAmount)
		totalData += dataAmount
	}

	if err := scanner.Err(); err != nil {
		log.Printf("[ERROR][PRIMARY NODE] Error al leer INFO.txt: %v", err)
	}

	return totalData
}

// calculateDataAmount devuelve el valor asociado a cada tipo de atributo
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

// SendTerminationSignal envía la señal de terminación a los servidores regionales y Data Nodes
func (s *Server) SendTerminationSignal(ctx context.Context, req *pb.TerminateProcess) (*pb.TerminateResponse, error) {
	log.Printf("[PRIMARY NODE] Señal de terminación recibida: %s", req.Result)

	if err := s.DataNodeClient.TerminateDataNodes(); err != nil {
		log.Printf("[ERROR][PRIMARY NODE] Error al terminar Data Nodes: %v", err)
		return nil, err
	}

	if err := s.RegionalClients.TerminateAll(); err != nil {
		log.Printf("[ERROR][PRIMARY NODE] Error al terminar Regional Servers: %v", err)
		return nil, err
	}

	log.Println("[PRIMARY NODE] Todos los nodos y servidores han sido terminados.")
	s.QuitChan <- true

	return &pb.TerminateResponse{Message: "Terminación completada"}, nil
}

// StartServer inicia el servidor gRPC del Primary Node
func StartServer(address string, s *Server) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Printf("[ERROR][PRIMARY NODE] Error al iniciar el servidor: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPrimaryNodeServiceServer(grpcServer, s)

	log.Printf("[PRIMARY NODE] Servidor escuchando en %s", address)
	return grpcServer.Serve(lis)
}
