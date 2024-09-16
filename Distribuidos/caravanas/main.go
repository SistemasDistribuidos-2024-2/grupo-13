package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	pb "grpc-server/proto/grpc-server/proto"
)

const (
	serverAddress       = "caravanas_container:50052"
	serverRemoteAddress = "konzu_container:50053"
)

var waitingTimeForSecondPackage time.Duration 

type Caravan struct {
	Id       string
	Priority bool
	Busy     bool
	Packages []*pb.PackageRequest
	mu       sync.Mutex
	LogFile  *os.File
}

type CaravanManager struct {
	CaravanP1  *Caravan
	CaravanP2  *Caravan
	CaravanN   *Caravan
	NextPIndex int
	mu         sync.Mutex
}

func NewCaravanManager() *CaravanManager {
	logFileP1, _ := os.Create("CaravanP1_log.txt")
	logFileP2, _ := os.Create("CaravanP2_log.txt")
	logFileN, _ := os.Create("CaravanN_log.txt")

	return &CaravanManager{
		CaravanP1: &Caravan{Id: "CaravanP1", Priority: true, Busy: false, LogFile: logFileP1},
		CaravanP2: &Caravan{Id: "CaravanP2", Priority: true, Busy: false, LogFile: logFileP2},
		CaravanN:  &Caravan{Id: "CaravanN", Priority: false, Busy: false, LogFile: logFileN},
	}
}

func (cm *CaravanManager) AssignPackage(packageRequest *pb.PackageRequest) (*Caravan, error) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if packageRequest.PackageType == "Ostronitas" || packageRequest.PackageType == "Prioritario" {
		if cm.NextPIndex == 0 && !cm.CaravanP1.Busy {
			cm.CaravanP1.Busy = true
			return cm.CaravanP1, nil
		} else if cm.NextPIndex == 1 && !cm.CaravanP2.Busy {
			cm.CaravanP2.Busy = true
			return cm.CaravanP2, nil
		}
		cm.NextPIndex = (cm.NextPIndex + 1) % 2
	} else if packageRequest.PackageType == "Normal" && !cm.CaravanN.Busy {
		cm.CaravanN.Busy = true
		return cm.CaravanN, nil
	}

	return nil, fmt.Errorf("no hay caravanas disponibles para el paquete")
}

func (cm *CaravanManager) FreeCaravan(caravan *Caravan) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	caravan.Busy = false
}

func (c *Caravan) waitForSecondPackage(waitTime time.Duration) {
	fmt.Printf("%s esperando un segundo paquete por %v...\n", c.Id, waitTime)
	time.Sleep(waitTime)
}

func (c *Caravan) sortPackagesByValue() {
	if len(c.Packages) > 1 {
		if c.Packages[0].Value < c.Packages[1].Value {
			c.Packages[0], c.Packages[1] = c.Packages[1], c.Packages[0]
		}
	}
}

func isGrineerDeliveryProfitable(pkg *pb.PackageRequest, failedAttempts int) bool {
	remainingValue := pkg.Value - float32(failedAttempts*100)
	return remainingValue > 0
}

func (c *Caravan) deliverPackages(client pb.CaravanServiceClient) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.sortPackagesByValue()

	for _, pkg := range c.Packages {
		maxAttempts := 3
		failedAttempts := 0
		var delivered bool

		for attempts := 0; attempts < maxAttempts; attempts++ {
			if pkg.PackageType == "Grineer" && attempts > 0 {
				if !isGrineerDeliveryProfitable(pkg, failedAttempts) {
					fmt.Printf("No es rentable seguir intentando entregar el paquete %s después de %d intentos fallidos.\n", pkg.PackageId, failedAttempts)
					c.reportStatus(client, pkg.PackageId, "No Entregado", failedAttempts)
					break
				}
			}

			fmt.Printf("Caravana %s intentando entregar el paquete %s. Intento #%d\n", c.Id, pkg.PackageId, attempts+1)
			if rand.Float32() > 0.15 {
				fmt.Printf("Entrega exitosa del paquete %s en el intento #%d\n", pkg.PackageId, attempts+1)
				c.reportStatus(client, pkg.PackageId, "Entregado", failedAttempts)
				delivered = true
				break
			} else {
				fmt.Printf("Fallo en la entrega del paquete %s en el intento #%d\n", pkg.PackageId, attempts+1)
				failedAttempts++
			}
			time.Sleep(3 * time.Second)
		}

		if !delivered && failedAttempts >= maxAttempts {
			fmt.Printf("Fallo definitivo en la entrega del paquete %s después de %d intentos\n", pkg.PackageId, maxAttempts)
			c.reportStatus(client, pkg.PackageId, "No Entregado", failedAttempts)
		}

		c.logDelivery(pkg, failedAttempts, delivered)
	}

	c.Packages = nil
}

func (c *Caravan) logDelivery(pkg *pb.PackageRequest, attempts int, delivered bool) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	status := "Not Delivered"
	if delivered {
		status = "Delivered"
	}

	logEntry := fmt.Sprintf("Fecha: %s, ID Paquete: %s, Tipo: %s, Valor: %.2f, Escolta: %s, Destino: %s, Intentos fallidos: %d, Estado: %s\n",
		currentTime, pkg.PackageId, pkg.PackageType, pkg.Value, pkg.Escort, pkg.Destination, attempts, status)

	_, err := c.LogFile.WriteString(logEntry)
	if err != nil {
		log.Fatalf("Error al escribir en el log de %s: %v", c.Id, err)
	}
}

func (c *Caravan) reportStatus(client pb.CaravanServiceClient, packageId string, status string, attempts int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.StatusRequestCaravana{
		PackageId: packageId,
		Status:    status,
		Attempts:  int32(attempts),
	}
	res, err := client.ReportStatus(ctx, req)
	if err != nil {
		log.Fatalf("Error reporting status: %v", err)
	}
	fmt.Printf("Estado reportado para el paquete %s: %s\n", packageId, res.Message)
}

type server struct {
	pb.UnimplementedCaravanServiceServer
	CaravanManager *CaravanManager
	Client         pb.CaravanServiceClient
}

func (s *server) AssignPackage(ctx context.Context, req *pb.PackageRequest) (*pb.PackageResponse, error) {
	caravan, err := s.CaravanManager.AssignPackage(req)
	if err != nil {
		return nil, err
	}

	caravan.mu.Lock()
	caravan.Packages = append(caravan.Packages, req)
	caravan.mu.Unlock()

	fmt.Printf("Asignando paquete %s de tipo %s a la caravana %s\n", req.PackageId, req.PackageType, caravan.Id)

	go func() {
		caravan.waitForSecondPackage(waitingTimeForSecondPackage)
		caravan.deliverPackages(s.Client)
		s.CaravanManager.FreeCaravan(caravan)
	}()

	return &pb.PackageResponse{Message: fmt.Sprintf("Paquete asignado a %s", caravan.Id)}, nil
}

func (s *server) CheckCaravanStatus(ctx context.Context, req *pb.EmptyRequest) (*pb.CaravanStatusResponse, error) {
	fmt.Println("Estado de las caravanas solicitado")
	return &pb.CaravanStatusResponse{
		Caravan_P_1: !s.CaravanManager.CaravanP1.Busy,
		Caravan_P_2: !s.CaravanManager.CaravanP2.Busy,
		Caravan_N:   !s.CaravanManager.CaravanN.Busy,
	}, nil
}

func main() {
	waitingTimeStr := os.Getenv("TIEMPO_OPERACION")

	if waitingTimeStr == "" {
		waitingTimeForSecondPackage = 5 * time.Second
	} else {
		waitingTime, err := strconv.Atoi(waitingTimeStr)
		if err != nil {
			fmt.Printf("Error al convertir TIEMPO_OPERACION: %v\n", err)
			waitingTimeForSecondPackage = 5 * time.Second
		} else {
			waitingTimeForSecondPackage = time.Duration(waitingTime) * time.Second
		}
	}

	fmt.Printf("Tiempo de espera para el segundo paquete: %v\n", waitingTimeForSecondPackage)

	rand.Seed(time.Now().UnixNano())
	conn, err := grpc.Dial(serverRemoteAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar al servidor remoto: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error cerrando la conexión: %v", err)
		}
	}(conn)

	client := pb.NewCaravanServiceClient(conn)

	caravanManager := NewCaravanManager()

	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Fallo al escuchar en el puerto %s: %v", serverAddress, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCaravanServiceServer(grpcServer, &server{CaravanManager: caravanManager, Client: client})

	fmt.Printf("Servidor gRPC escuchando en el puerto %s...\n", serverAddress)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fallo al servir: %v", err)
	}
}
