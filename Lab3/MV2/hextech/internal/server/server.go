package server

import (
	"context"
	"fmt"
	"os"
	"sync"
	"crypto/sha256"
	"encoding/hex"
	"hextech/internal/storage"
	"hextech/proto"
)

type HextechServer struct {
	mu          sync.Mutex
	storage     map[string]*storage.RegionData
	serverID    int
	peers       []proto.HextechServiceClient
	appliedLogs map[string]bool

	proto.UnimplementedHextechServiceServer
}

func generateLogHash(log string) string {
	hash := sha256.Sum256([]byte(log))
	return hex.EncodeToString(hash[:])
}

func NewHextechServer(serverID int) *HextechServer {
	return &HextechServer{
		storage:     make(map[string]*storage.RegionData),
		serverID:    serverID,
		peers:       []proto.HextechServiceClient{},
		appliedLogs: make(map[string]bool),
	}
}

func (s *HextechServer) getLogFilePath() string {
	return fmt.Sprintf("HextechLogs_%d.txt", s.serverID)
}

func (s *HextechServer) AddPeer(peer proto.HextechServiceClient) {
	s.peers = append(s.peers, peer)
	fmt.Printf("[Servidor Hextech] Peer agregado. Total de peers: %d\n", len(s.peers))
}

func (s *HextechServer) AddProductServer(ctx context.Context, req *proto.AddProductRequest) (*proto.ClockResponse, error) {
	fmt.Printf("[Servidor Hextech][Recibida solicitud] [AddProduct]: Región=%s, Producto=%s, Cantidad=%d\n", req.Region, req.Product, req.Quantity)
	return s.AddProduct(ctx, req)
}

func (s *HextechServer) DeleteProductServer(ctx context.Context, req *proto.DeleteProductRequest) (*proto.ClockResponse, error) {
	fmt.Printf("[Servidor Hextech][Recibida solicitud] [DeleteProduct]: Región=%s, Producto=%s\n", req.Region, req.Product)
	return s.DeleteProduct(ctx, req)
}

func (s *HextechServer) UpdateProductServer(ctx context.Context, req *proto.UpdateProductRequest) (*proto.ClockResponse, error) {
	fmt.Printf("[Servidor Hextech][Recibida solicitud] [UpdateProduct]: Región=%s, Producto=%s, Cantidad=%d\n", req.Region, req.Product, req.Quantity)
	return s.UpdateProduct(ctx, req)
}

func (s *HextechServer) RenameProductServer(ctx context.Context, req *proto.RenameProductRequest) (*proto.ClockResponse, error) {
	fmt.Printf("[Servidor Hextech][Recibida solicitud] [RenameProduct]: Región=%s, ProductoAntiguo=%s, ProductoNuevo=%s\n", req.Region, req.OldProduct, req.NewProduct)
	return s.RenameProduct(ctx, req)
}

func (s *HextechServer) PropagateChanges(ctx context.Context, req *proto.PropagationRequest) (*proto.PropagationResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	region := req.Region
	fmt.Printf("[Servidor Hextech] Recibiendo cambios para la región [%s]\n", region)

	if _, exists := s.storage[region]; !exists {
		s.storage[region] = storage.NewRegionData(fmt.Sprintf("%s.txt", region), len(s.peers)+1)
		err := os.WriteFile(s.storage[region].FilePath, []byte{}, 0644)
		if err != nil {
			fmt.Printf("[Servidor Hextech] Error al crear archivo de región [%s]: %v\n", region, err)
			return nil, err
		}
	}

	regionData := s.storage[region]

	for _, log := range req.ChangeLog {
		logHash := generateLogHash(log)

		if s.appliedLogs[logHash] {
			fmt.Printf("[Servidor Hextech] Log duplicado ignorado: %s\n", log)
			continue
		}

		err := storage.ApplyLogToFile(regionData.FilePath, log, s.serverID)
		if err != nil {
			fmt.Printf("[Servidor Hextech] Error al aplicar log [%s]: %v\n", log, err)
			return nil, err
		}

		s.appliedLogs[logHash] = true
	}

	for i, value := range req.VectorClock {
		if value > regionData.VectorClock[i] {
			regionData.VectorClock[i] = value
		}
	}

	fmt.Printf("[Servidor Hextech] Cambios aplicados a la región [%s].\n", region)
	return &proto.PropagationResponse{Status: "success"}, nil
}

func (s *HextechServer) ForceMerge(ctx context.Context, req *proto.ErrorMergeRequest) (*proto.ConfirmationError, error) {
	fmt.Printf("[Servidor Hextech][Merge forzado solicitado]: Región=%s\n", req.Region)
	return &proto.ConfirmationError{
		Confirmation: "Merge forzado completado",
	}, nil
}
func (s *HextechServer) AddProduct(ctx context.Context, req *proto.AddProductRequest) (*proto.ClockResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	region := req.Region
	product := req.Product
	quantity := req.Quantity

	if _, exists := s.storage[region]; !exists {
		regionData := storage.NewRegionData(fmt.Sprintf("%s.txt", region), len(s.peers)+1)
		s.storage[region] = regionData

		err := os.WriteFile(regionData.FilePath, []byte{}, 0644)
		if err != nil {
			fmt.Printf("[Servidor Hextech] Error al crear archivo para región: %v\n", err)
			return nil, err
		}
		fmt.Printf("[Servidor Hextech] Nueva región creada: %s\n", region)
	}

	regionData := s.storage[region]

	existingQuantity, err := storage.GetProductQuantity(regionData.FilePath, product)
	if err != nil && err != storage.ErrProductNotFound {
		fmt.Printf("[Servidor Hextech] Error al verificar producto en archivo: %v\n", err)
		return nil, err
	}

	newQuantity := quantity
	if err == nil { 
		newQuantity += existingQuantity
	}

	err = storage.UpdateValueInFile(regionData.FilePath, region, product, newQuantity)
	if err != nil {
		fmt.Printf("[Servidor Hextech] Error al actualizar producto en archivo: %v\n", err)
		return nil, err
	}

	logEntry := fmt.Sprintf("AgregarProducto %s %s %d", region, product, quantity)
	regionData.AddLog(logEntry, s.serverID)
	s.writeToLogFile(logEntry)

	fmt.Printf("[Servidor Hextech] Producto agregado o actualizado: Región=%s, Producto=%s, NuevaCantidad=%d\n", region, product, newQuantity)
	return &proto.ClockResponse{VectorClock: regionData.VectorClock}, nil
}

func (s *HextechServer) DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest) (*proto.ClockResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	region := req.Region
	product := req.Product

	if regionData, exists := s.storage[region]; exists {
		err := storage.RemoveProductFromFile(regionData.FilePath, product)
		if err != nil {
			if err == storage.ErrProductNotFound {
				fmt.Printf("[Servidor Hextech] Producto no encontrado: Región=%s, Producto=%s\n", region, product)
				return &proto.ClockResponse{VectorClock: regionData.VectorClock}, nil
			}
			fmt.Printf("[Servidor Hextech] Error al intentar eliminar producto del archivo: %v\n", err)
			return &proto.ClockResponse{VectorClock: regionData.VectorClock}, err
		}

		logEntry := fmt.Sprintf("BorrarProducto %s %s", region, product)
		s.writeToLogFile(logEntry)
		regionData.AddLog(logEntry, s.serverID)

		fmt.Printf("[Servidor Hextech] Producto eliminado exitosamente: Región=%s, Producto=%s\n", region, product)
		return &proto.ClockResponse{VectorClock: regionData.VectorClock}, nil
	}

	fmt.Printf("[Servidor Hextech] Región no encontrada: %s\n", region)
	return &proto.ClockResponse{VectorClock: make([]int32, len(s.peers)+1)}, nil
}


func (s *HextechServer) RenameProduct(ctx context.Context, req *proto.RenameProductRequest) (*proto.ClockResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	region := req.Region
	oldProduct := req.OldProduct
	newProduct := req.NewProduct

	if regionData, exists := s.storage[region]; exists {
		exists, err := storage.CheckProductExists(regionData.FilePath, oldProduct)
		if err != nil {
			fmt.Printf("[Servidor Hextech] Error al verificar producto en archivo: %v\n", err)
			return &proto.ClockResponse{VectorClock: regionData.VectorClock}, err
		}

		if !exists {
			fmt.Printf("[Servidor Hextech] Producto no encontrado: Región=%s, Producto=%s\n", region, oldProduct)
			return &proto.ClockResponse{VectorClock: regionData.VectorClock}, nil
		}

		err = storage.UpdateFile(regionData.FilePath, oldProduct, newProduct)
		if err != nil {
			fmt.Printf("[Servidor Hextech] Error al renombrar producto en archivo: %v\n", err)
			return &proto.ClockResponse{VectorClock: regionData.VectorClock}, err
		}

		logEntry := fmt.Sprintf("RenombrarProducto %s %s %s", region, oldProduct, newProduct)
		s.writeToLogFile(logEntry)
		regionData.AddLog(logEntry, s.serverID)

		fmt.Printf("[Servidor Hextech] Producto renombrado: Región=%s, ProductoAntiguo=%s, ProductoNuevo=%s\n", region, oldProduct, newProduct)
		return &proto.ClockResponse{VectorClock: regionData.VectorClock}, nil
	}

	fmt.Printf("[Servidor Hextech] Error: Región no encontrada: %s\n", region)
	return &proto.ClockResponse{VectorClock: make([]int32, len(s.peers)+1)}, nil
}


func (s *HextechServer) UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest) (*proto.ClockResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	region := req.Region
	product := req.Product
	quantity := req.Quantity

	if regionData, exists := s.storage[region]; exists {
		regionData.AddLog(fmt.Sprintf("ActualizarValor %s %s %d", region, product, quantity), s.serverID)

		if err := storage.UpdateValueInFile(regionData.FilePath, region, product, quantity); err != nil {
			fmt.Printf("[Servidor Hextech] Error al actualizar producto en archivo: %v\n", err)
			return &proto.ClockResponse{VectorClock: regionData.VectorClock}, err
		}

		logEntry := fmt.Sprintf("ActualizarValor %s %s %d", region, product, quantity)
		s.writeToLogFile(logEntry)

		fmt.Printf("[Servidor Hextech] Producto actualizado: Región=%s, Producto=%s, NuevaCantidad=%d\n", region, product, quantity)
		return &proto.ClockResponse{VectorClock: regionData.VectorClock}, nil
	}

	fmt.Printf("[Servidor Hextech] Error: Región no encontrada: %s\n", region)
	return &proto.ClockResponse{VectorClock: make([]int32, len(s.peers)+1)}, nil
}


func (s *HextechServer) writeToLogFile(message string) {
	logFilePath := fmt.Sprintf("HextechLogs_%d.txt", s.serverID)

	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("[Servidor Hextech] Error al abrir archivo de logs [%s]: %v\n", logFilePath, err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(message + "\n"); err != nil {
		fmt.Printf("[Servidor Hextech] Error al escribir en archivo de logs [%s]: %v\n", logFilePath, err)
	}
}

func (s *HextechServer) GetProductServer(ctx context.Context, req *proto.GetProductRequest) (*proto.ProductResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	region := req.Region
	product := req.Product

	if regionData, exists := s.storage[region]; exists {
		quantity, err := storage.GetProductQuantity(regionData.FilePath, product)
		if err != nil {
			if err == storage.ErrProductNotFound {
				fmt.Printf("[Servidor Hextech] Producto no encontrado: Región=%s, Producto=%s\n", region, product)
				return &proto.ProductResponse{
					Quantity:    0,
					VectorClock: regionData.VectorClock,
				}, nil
			}
			fmt.Printf("[Servidor Hextech] Error al obtener producto: %v\n", err)
			return nil, err
		}

		fmt.Printf("[Servidor Hextech] Producto encontrado: Región=%s, Producto=%s, Cantidad=%d\n", region, product, quantity)
		return &proto.ProductResponse{
			Quantity:    quantity,
			VectorClock: regionData.VectorClock,
		}, nil
	}

	fmt.Printf("[Servidor Hextech] Región no encontrada: %s\n", region)
	return &proto.ProductResponse{
		Quantity:    0,
		VectorClock: make([]int32, len(s.peers)+1),
	}, nil
}

