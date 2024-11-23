package storage
import (
	"errors"
	"fmt"
	"os"
	"strings"
)


type RegionData struct {
	FilePath    string
	VectorClock []int32
	ChangeLog   []string
}

func NewRegionData(filePath string, clockSize int) *RegionData {
	return &RegionData{
		FilePath:    filePath,
		VectorClock: make([]int32, clockSize),
		ChangeLog:   []string{},
	}
}

func (r *RegionData) AddLog(entry string, serverID int) {
	r.ChangeLog = append(r.ChangeLog, entry)
	r.VectorClock[serverID]++
}

func (r *RegionData) ClearLogs() {
	r.ChangeLog = []string{}
}

var ErrProductNotFound = errors.New("producto no encontrado")

func GetProductQuantity(filePath, product string) (int32, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 3 && parts[1] == product {
			var quantity int32
			fmt.Sscanf(parts[2], "%d", &quantity)
			return quantity, nil
		}
	}

	return 0, ErrProductNotFound
}


// UpdateValueInFile actualiza o agrega un producto con su cantidad en el archivo
func UpdateValueInFile(filePath,region string, product string, newQuantity int32) error {
	content, err := os.ReadFile(filePath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	lines := strings.Split(string(content), "\n")
	updated := false

	for i, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 3 && parts[1] == product {
			lines[i] = fmt.Sprintf("%s %s %d", parts[0], product, newQuantity)
			updated = true
			break
		}
	}

	if !updated { // Si no se encontró, agregarlo
		lines = append(lines, fmt.Sprintf("%s %s %d", region, product, newQuantity))
	}

	updatedContent := strings.Join(lines, "\n")
	return os.WriteFile(filePath, []byte(updatedContent), 0644)
}

func RemoveProductFromFile(filePath, product string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	found := false

	// Filtra las líneas que no contienen el producto
	newLines := []string{}
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 3 && parts[1] == product {
			found = true
			continue // No incluir esta línea en el archivo actualizado
		}
		newLines = append(newLines, line)
	}

	if !found {
		return ErrProductNotFound
	}

	// Sobreescribe el archivo con las líneas restantes
	updatedContent := strings.Join(newLines, "\n")
	return os.WriteFile(filePath, []byte(updatedContent), 0644)
}

func ApplyLogToFile(filePath, log string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	parts := strings.Fields(log)
	if len(parts) < 3 {
		return fmt.Errorf("log mal formado: %s", log)
	}

	action := parts[0]
	region := parts[1]
	product := parts[2]
	var quantity int32
	if len(parts) > 3 {
		fmt.Sscanf(parts[3], "%d", &quantity)
	}

	switch action {
	case "AgregarProducto":
		return UpdateValueInFile(filePath, region, product, quantity)
	case "BorrarProducto":
		return RemoveProductFromFile(filePath, product)
	case "RenombrarProducto":
		if len(parts) < 4 {
			return fmt.Errorf("log mal formado para renombrar: %s", log)
		}
		newProduct := parts[3]
		return UpdateFile(filePath, product, newProduct)
	case "ActualizarValor":
		return UpdateValueInFile(filePath, region, product, quantity)
	default:
		return fmt.Errorf("acción desconocida: %s", action)
	}
}
