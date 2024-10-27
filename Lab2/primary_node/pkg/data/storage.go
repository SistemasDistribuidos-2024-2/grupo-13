package data

import (
	"fmt"
	"os"
)

// Verifica si INFO.txt existe, si no lo crea
func InitInfoFile() error {
	_, err := os.Stat("pkg/data/INFO.txt")
	if os.IsNotExist(err) {
		file, err := os.Create("pkg/data/INFO.txt")
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

// Escribe una entrada en INFO.txt
func WriteInfo(id int, dataNode int, name, status string) error {
	entry := fmt.Sprintf("%d,%d,%s,%s\n", id, dataNode, name, status)
	file, err := os.OpenFile("pkg/data/INFO.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(entry)
	return err
}
