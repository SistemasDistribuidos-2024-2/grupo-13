package server

import (
	"context"
	"fmt"
	"time"
	"strings"
	"hextech/internal/storage"
	"hextech/proto"
	"os"
)

func (s *HextechServer) StartPropagation() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		s.mu.Lock()

		for _, peer := range s.peers {
			for region, regionData := range s.storage {
				if len(regionData.ChangeLog) > 0 {
					req := &proto.PropagationRequest{
						Region:      region,
						ChangeLog:   regionData.ChangeLog,
						VectorClock: regionData.VectorClock,
					}
					go func(peer proto.HextechServiceClient, req *proto.PropagationRequest) {
						_, err := peer.PropagateChanges(context.Background(), req)
						if err != nil {
							fmt.Printf("[Servidor Hextech] Error al propagar a un peer: %v\n", err)
						} else {
							// Limpiar los logs después de la propagación exitosa
							regionData.ClearLogs()
							logFilePath := s.getLogFilePath()
							err := os.WriteFile(logFilePath, []byte{}, 0644)
							if err != nil {
								fmt.Printf("[Servidor Hextech] Error al limpiar archivo de logs: %v\n", err)
							}
						}
					}(peer, req)
				}
			}
		}

		s.mu.Unlock()
	}
}

func createRegionOnPeer(peer proto.HextechServiceClient, region string, regionData *storage.RegionData) error {
	fmt.Printf("[Servidor Hextech] Creando región [%s] en peer...\n", region)
	for _, log := range regionData.ChangeLog {
		parts := strings.Fields(log)
		if len(parts) >= 3 {
			product := parts[1]
			var quantity int32
			fmt.Sscanf(parts[2], "%d", &quantity)

			req := &proto.AddProductRequest{
				Region:   region,
				Product:  product,
				Quantity: quantity,
			}

			_, err := peer.AddProductServer(context.Background(), req)
			if err != nil {
				return fmt.Errorf("[Servidor Hextech] Error al crear región [%s] en peer: %v", region, err)
			}
		}
	}
	fmt.Printf("[Servidor Hextech] Región [%s] creada exitosamente en peer.\n", region)
	return nil
}
