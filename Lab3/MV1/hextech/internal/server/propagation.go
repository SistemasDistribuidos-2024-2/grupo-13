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
		fmt.Println("[Servidor Hextech] Propagando cambios a los peers...")

		for _, peer := range s.peers {
			for region, regionData := range s.storage {
				if len(regionData.ChangeLog) > 0 {
					// Construir la solicitud de propagación
					req := &proto.PropagationRequest{
						Region:      region,
						ChangeLog:   regionData.ChangeLog,
						VectorClock: regionData.VectorClock,
					}

					// Intentar enviar la solicitud de propagación
					_, err := peer.PropagateChanges(context.Background(), req)
					if err != nil {
						fmt.Printf("[Servidor Hextech] Error al propagar al peer: %v\n", err)

						// Si la región no existe en el peer, intentamos crearla
						if strings.Contains(err.Error(), "Región no encontrada") {
							err := createRegionOnPeer(peer, region, regionData)
							if err != nil {
								fmt.Printf("[Servidor Hextech] Error al crear región en peer: %v\n", err)
							}
						}
					}
				}
			}
		}

		// Limpiar los logs de la región después de propagar
		for _, regionData := range s.storage {
			regionData.ClearLogs()
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
