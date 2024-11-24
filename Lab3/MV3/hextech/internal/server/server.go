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
	appliedLogs map[string]bool // Logs ya aplicados

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
	fmt.Printf("[Servidor Hextech] Propagando cambios en la región [%s]\n", region)

	// Verificar si la región existe; si no, crearla
	if _, exists := s.storage[region]; !exists {
		s.storage[region] = storage.NewRegionData(fmt.Sprintf("%s.txt", region), len(s.peers)+1)
		err := os.WriteFile(s.storage[region].FilePath, []byte{}, 0644)
		if err != nil {
			fmt.Printf("[Servidor Hextech] Error al crear archivo de región [%s]: %v\n", region, err)
			return nil, err
		}
	}

	regionData := s.storage[region]
	// Procesar cada log recibido
	for _, log := range req.ChangeLog {
		logHash := generateLogHash(log)

		// Si el log ya fue procesado, omitirlo
		if s.appliedLogs[logHash] {
			continue
		}

		// Aplicar el log al archivo de la región
		err := storage.ApplyLogToFile(regionData.FilePath, log, s.serverID)
		if err != nil {
			fmt.Printf("[Servidor Hextech] Error al aplicar log [%s]: %v\n", log, err)
			return nil, err
		}

		// Marcar el log como aplicado
		s.appliedLogs[logHash] = true
	}

	// Actualizar el reloj vectorial
	for i, value := range req.VectorClock {
		if value > regionData.VectorClock[i] {
			regionData.VectorClock[i] = value
		}
	}

	// Propagar los cambios a otros peers
	for _, peer := range s.peers {
		if peer != nil { // Asegurarse de que el peer esté conectado
			go func(peer proto.HextechServiceClient) {
				_, err := peer.PropagateChanges(context.Background(), req)
				if err != nil {
					fmt.Printf("[Servidor Hextech] Error al propagar cambios a un peer: %v\n", err)
				}
			}(peer)
		}
	}

	fmt.Printf("[Servidor Hextech] Cambios propagados a la región [%s].\n", region)
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

	// Verificar si la región existe en memoria
	if _, exists := s.storage[region]; !exists {
		// Crear la región en memoria y su archivo asociado
		regionData := storage.NewRegionData(fmt.Sprintf("%s.txt", region), len(s.peers)+1)
		s.storage[region] = regionData

		// Crear el archivo si no existe
		err := os.WriteFile(regionData.FilePath, []byte{}, 0644)
		if err != nil {
			fmt.Printf("[Servidor Hextech] Error al crear archivo para región: %v\n", err)
			return nil, err
		}
		fmt.Printf("[Servidor Hextech] Nueva región creada: %s\n", region)
	}

	regionData := s.storage[region]

	// Verificar si el producto ya existe en el archivo
	existingQuantity, err := storage.GetProductQuantity(regionData.FilePath, product)
	if err != nil && err != storage.ErrProductNotFound {
		fmt.Printf("[Servidor Hextech] Error al verificar producto en archivo: %v\n", err)
		return nil, err
	}

	newQuantity := quantity
	if err == nil { // Si el producto existe, sumar las cantidades
		newQuantity += existingQuantity
	}

	// Actualizar o añadir el producto en el archivo
	err = storage.UpdateValueInFile(regionData.FilePath, region, product, newQuantity)
	if err != nil {
		fmt.Printf("[Servidor Hextech] Error al actualizar producto en archivo: %v\n", err)
		return nil, err
	}

	// Registrar en el log de memoria y en el archivo global de logs
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

	// Verifica si la región existe en el almacenamiento
	if regionData, exists := s.storage[region]; exists {
		// Intenta eliminar el producto del archivo
		err := storage.RemoveProductFromFile(regionData.FilePath, product)
		if err != nil {
			if err == storage.ErrProductNotFound {
				fmt.Printf("[Servidor Hextech] Producto no encontrado: Región=%s, Producto=%s\n", region, product)
				return nil, fmt.Errorf("Producto %s no encontrado en región %s", product, region)
			}
			fmt.Printf("[Servidor Hextech] Error al eliminar producto del archivo: %v\n", err)
			return nil, err
		}

		// Registra en el archivo de logs
		logEntry := fmt.Sprintf("BorrarProducto %s %s", region, product)
		s.writeToLogFile(logEntry)
		regionData.AddLog(logEntry, s.serverID)

		fmt.Printf("[Servidor Hextech] Producto eliminado: Región=%s, Producto=%s\n", region, product)
		return &proto.ClockResponse{VectorClock: regionData.VectorClock}, nil
	}

	fmt.Printf("[Servidor Hextech] Error: Región no encontrada: %s\n", region)
	return nil, fmt.Errorf("Región %s no encontrada", region)
}



func (s *HextechServer) RenameProduct(ctx context.Context, req *proto.RenameProductRequest) (*proto.ClockResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	region := req.Region
	oldProduct := req.OldProduct
	newProduct := req.NewProduct

	if regionData, exists := s.storage[region]; exists {
		regionData.AddLog(fmt.Sprintf("RenombrarProducto %s %s %s", region, oldProduct, newProduct), s.serverID)

		if err := storage.UpdateFile(regionData.FilePath, oldProduct, newProduct); err != nil {
			fmt.Printf("[Servidor Hextech] Error al renombrar producto en archivo: %v\n", err)
			return nil, err
		}

		// Registra en el archivo de logs
		logEntry := fmt.Sprintf("RenombrarProducto %s %s %s", region, oldProduct, newProduct)
		s.writeToLogFile(logEntry)

		fmt.Printf("[Servidor Hextech] Producto renombrado: Región=%s, ProductoAntiguo=%s, ProductoNuevo=%s\n", region, oldProduct, newProduct)
		return &proto.ClockResponse{VectorClock: regionData.VectorClock}, nil
	}

	fmt.Printf("[Servidor Hextech] Error: Región no encontrada: %s\n", region)
	return nil, fmt.Errorf("Región %s no encontrada", region)
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
			return nil, err
		}

		// Registra en el archivo de logs
		logEntry := fmt.Sprintf("ActualizarValor %s %s %d", region, product, quantity)
		s.writeToLogFile(logEntry)

		fmt.Printf("[Servidor Hextech] Producto actualizado: Región=%s, Producto=%s, NuevaCantidad=%d\n", region, product, quantity)
		return &proto.ClockResponse{VectorClock: regionData.VectorClock}, nil
	}

	fmt.Printf("[Servidor Hextech] Error: Región no encontrada: %s\n", region)
	return nil, fmt.Errorf("Región %s no encontrada", region)
}


func (s *HextechServer) writeToLogFile(message string) {
	logFilePath := fmt.Sprintf("HextechLogs_%d.txt", s.serverID) // Basado en el ID del servidor

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

	// Verifica si la región existe
	if regionData, exists := s.storage[region]; exists {
		// Busca el producto en el archivo de la región
		quantity, err := storage.GetProductQuantity(regionData.FilePath, product)
		if err != nil {
			if err == storage.ErrProductNotFound {
				fmt.Printf("[Servidor Hextech] Producto no encontrado: Región=%s, Producto=%s\n", region, product)
				return nil, fmt.Errorf("Producto %s no encontrado en región %s", product, region)
			}
			fmt.Printf("[Servidor Hextech] Error al obtener producto: %v\n", err)
			return nil, err
		}

		// Retorna la respuesta con la cantidad y el reloj vectorial
		fmt.Printf("[Servidor Hextech] Producto encontrado: Región=%s, Producto=%s, Cantidad=%d\n", region, product, quantity)
		return &proto.ProductResponse{
			Quantity:    quantity,
			VectorClock: regionData.VectorClock,
		}, nil
	}

	// Región no encontrada
	fmt.Printf("[Servidor Hextech] Región no encontrada: %s\n", region)
	return nil, fmt.Errorf("Región %s no encontrada", region)
}
