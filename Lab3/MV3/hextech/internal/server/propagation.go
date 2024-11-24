package server

import (
	"context"
	"fmt"
	"time"
	"strings"
	"hextech/internal/storage"
	"hextech/proto"
)

func (s *HextechServer) StartPropagation() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		s.mu.Lock()
		fmt.Println("[Servidor Hextech] Iniciando propagación de cambios...")

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
							fmt.Printf("[Servidor Hextech] Error al propagar cambios a un peer: %v\n", err)
						}
					}(peer, req)
				}
			}
		}

		for _, regionData := range s.storage {
			regionData.ClearLogs()
		}

		s.mu.Unlock()
		fmt.Println("[Servidor Hextech] Propagación de cambios completada.")
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
