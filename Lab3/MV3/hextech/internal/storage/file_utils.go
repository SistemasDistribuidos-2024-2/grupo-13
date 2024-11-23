package storage

import (
	"os"
	"strings"
	"fmt"
)

func WriteToFile(filePath, record string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(record + "\n")
	return err
}

func WriteAllToFile(filePath string, data map[string]string) error {
    f, err := os.Create(filePath) // Crea o sobrescribe el archivo
    if err != nil {
        return err
    }
    defer f.Close()

    for key, value := range data {
        _, err := f.WriteString(fmt.Sprintf("%s %s\n", key, value))
        if err != nil {
            return err
        }
    }

    return nil
}

func UpdateFile(filePath, oldName, newName string) error {
    content, err := os.ReadFile(filePath)
    if err != nil {
        return err
    }

    updatedContent := strings.ReplaceAll(string(content), oldName, newName)
    return os.WriteFile(filePath, []byte(updatedContent), 0644)
}


func RemoveFromFile(filePath, record string) error {
	content, err := os.ReadFile(filePath) // Lee el archivo completo
	if err != nil {
		return err
	}

	lines := []string{}
	for _, line := range strings.Split(string(content), "\n") {
		if line != record && line != "" { // Filtra la línea a eliminar
			lines = append(lines, line)
		}
	}

	return os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644) // Sobrescribe el archivo
}

