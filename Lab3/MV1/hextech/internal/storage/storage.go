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

func ApplyLogToFile(filePath string, log string, serverID int) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		// Si el archivo no existe, crear uno nuevo
		if errors.Is(err, os.ErrNotExist) {
			return os.WriteFile(filePath, []byte(log+"\n"), 0644)
		}
		return err
	}

	lines := strings.Split(string(content), "\n") // Divide el contenido del archivo en líneas

	// Procesar el log para aplicar los cambios
	parts := strings.Fields(log)
	if len(parts) < 4 {
		return fmt.Errorf("log mal formado: %s", log)
	}

	action := parts[0]
	region := parts[1]
	product := parts[2]
	var quantity int32
	if len(parts) > 3 {
		fmt.Sscanf(parts[3], "%d", &quantity)
	}

	updated := false
	switch action {
	case "AgregarProducto":
		for i, line := range lines {
			lineParts := strings.Fields(line)
			if len(lineParts) == 3 && lineParts[1] == product {
				oldQuantity, _ := strconv.Atoi(lineParts[2])
				lines[i] = fmt.Sprintf("%s %s %d", region, product, oldQuantity+int(quantity))
				updated = true
				break
			}
		}
		if !updated {
			lines = append(lines, fmt.Sprintf("%s %s %d", region, product, quantity))
		}
	case "BorrarProducto":
		newLines := []string{}
		for _, line := range lines {
			lineParts := strings.Fields(line)
			if len(lineParts) == 3 && lineParts[1] == product {
				continue // Omitir esta línea
			}
			newLines = append(newLines, line)
		}
		lines = newLines
	case "RenombrarProducto":
		if len(parts) < 4 {
			return fmt.Errorf("log mal formado para renombrar: %s", log)
		}
		newProduct := parts[3]
		for i, line := range lines {
			lineParts := strings.Fields(line)
			if len(lineParts) == 3 && lineParts[1] == product {
				lines[i] = fmt.Sprintf("%s %s %s", region, newProduct, lineParts[2])
				updated = true
				break
			}
		}
	case "ActualizarValor":
		for i, line := range lines {
			lineParts := strings.Fields(line)
			if len(lineParts) == 3 && lineParts[1] == product {
				lines[i] = fmt.Sprintf("%s %s %d", region, product, quantity)
				updated = true
				break
			}
		}
	default:
		return fmt.Errorf("acción desconocida: %s", action)
	}

	// Escribir los cambios de vuelta al archivo de la región
	updatedContent := strings.Join(lines, "\n")
	err = os.WriteFile(filePath, []byte(updatedContent+"\n"), 0644)
	if err != nil {
		return err
	}

	// Registrar en el archivo de logs individual
	err = appendToLogFile(serverID, fmt.Sprintf("%s %s %s %d", action, region, product, quantity))
	if err != nil {
		return fmt.Errorf("error al registrar en el log del servidor [%d]: %v", serverID, err)
	}

	return nil
}


func appendToLogFile(serverID int, entry string) error {
	logFilePath := fmt.Sprintf("HextechLogs_%d.txt", serverID) // Archivo de logs individual

	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error al abrir el archivo de logs [%s]: %v", logFilePath, err)
	}
	defer f.Close()

	_, err = f.WriteString(entry + "\n")
	if err != nil {
		return fmt.Errorf("error al escribir en el archivo de logs [%s]: %v", logFilePath, err)
	}
	return nil
}

