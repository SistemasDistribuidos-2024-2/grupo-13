package server

import (
	"fmt"
	"strings"
	"context"

	"hextech/internal/storage"
	"hextech/proto"
)

// mergeWithPeers coordina el merge entre servidores utilizando la estrategia basada en ID.
func (s *HextechServer) mergeWithPeers(peers []proto.HextechServiceClient) {
	for region, localData := range s.storage {
		mergedClock := make([]int32, len(localData.VectorClock))
		copy(mergedClock, localData.VectorClock)
		mergedLog := append([]string{}, localData.ChangeLog...)

		// Determinar si este servidor es el dominante
		isDominant := true
		for _, peer := range peers {
			mergeRequest := &proto.MergeRequest{Region: region}
			response, err := peer.RequestMerge(context.Background(), mergeRequest)
			if err != nil {
				fmt.Printf("Error al obtener datos de peer para región %s: %v\n", region, err)
				continue
			}

			// Comparar relojes vectoriales y determinar el servidor dominante
			for i := range response.VectorClock {
				if response.VectorClock[i] > mergedClock[i] {
					mergedClock[i] = response.VectorClock[i]
				}
			}

			// Verificar si hay un servidor con menor ID
			if response.VectorClock[s.serverID] < mergedClock[s.serverID] {
				isDominant = false
			}

			// Combinar los logs del peer
			mergedLog = append(mergedLog, response.ChangeLog...)
		}

		if isDominant {
			// Este servidor es dominante, realiza el merge
			fmt.Printf("Servidor %d es dominante para la región %s\n", s.serverID, region)
			resolvedRecords := resolveConflicts(mergedLog)
			storage.WriteAllToFile(localData.FilePath, resolvedRecords)
			localData.VectorClock = mergedClock
			localData.ChangeLog = []string{}
		} else {
			fmt.Printf("Servidor %d no es dominante para la región %s\n", s.serverID, region)
		}
	}
}
func (s *HextechServer) RequestMerge(ctx context.Context, req *proto.MergeRequest) (*proto.MergeResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	region := req.Region
	if regionData, exists := s.storage[region]; exists {
		return &proto.MergeResponse{
			ChangeLog:   regionData.ChangeLog,
			VectorClock: regionData.VectorClock,
		}, nil
	}

	return nil, fmt.Errorf("Región %s no encontrada", region)
}

func resolveConflicts(changeLog []string) map[string]string {
    resolved := make(map[string]string)

    for _, entry := range changeLog {
        parts := strings.Fields(entry) // Divide la entrada en palabras
        if len(parts) < 3 {
            continue
        }

        action := parts[0]
        region := parts[1]
        product := parts[2]

        switch action {
        case "AgregarProducto", "ActualizarValor":
            // Usa el último valor encontrado
            if len(parts) > 3 {
                resolved[fmt.Sprintf("%s %s", region, product)] = parts[3]
            }
        case "RenombrarProducto":
            // Cambia el nombre del producto en los registros
            if len(parts) > 3 {
                resolved[fmt.Sprintf("%s %s", region, parts[3])] = resolved[fmt.Sprintf("%s %s", region, product)]
                delete(resolved, fmt.Sprintf("%s %s", region, product))
            }
        case "BorrarProducto":
            // Elimina el producto del registro
            delete(resolved, fmt.Sprintf("%s %s", region, product))
        }
    }

    return resolved
}
