package data

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regional_server/pkg/models"
	"strconv"
	"strings"
)

// Parámetros del sistema (cargados desde INPUT.txt)
type InputConfig struct {
	PS float64 // Probabilidad de sacrificio (0 <= PS <= 1)
	TE int     // Tiempo de espera para enviar información (TE > 0)
	TD int     // Tiempo de ataque de Diaboromon (TD > 0)
	CD int     // Cantidad de datos necesarios para evolucionar en Omegamon
	VI int     // Vida inicial para Greymon y Garurumon
}

// Carga de los parámetros del archivo INPUT.txt
func LoadInputConfig() (*InputConfig, error) {
	file, err := os.Open("pkg/data/INPUT.txt")
	if err != nil {
		log.Fatalf("Error al abrir INPUT.txt: %v", err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		if len(values) != 5 {
			return nil, errors.New("INPUT.txt debe contener exactamente 5 valores separados por comas")
		}
		ps, _ := strconv.ParseFloat(values[0], 64)
		te, _ := strconv.Atoi(values[1])
		td, _ := strconv.Atoi(values[2])
		cd, _ := strconv.Atoi(values[3])
		vi, _ := strconv.Atoi(values[4])
		return &InputConfig{
			PS: ps,
			TE: te,
			TD: td,
			CD: cd,
			VI: vi,
		}, nil
	}

	return nil, errors.New("No se pudo leer INPUT.txt")
}

// Carga de los Digimon desde el archivo DIGIMONS.TXT
func LoadDigimons() ([]models.Digimon, error) {
	file, err := os.Open("pkg/data/DIGIMONS.TXT")
	if err != nil {
		log.Fatalf("Error al abrir DIGIMONS.TXT: %v", err)
		return nil, err
	}
	defer file.Close()

	var digimons []models.Digimon
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		if len(line) != 2 {
			log.Println("Línea inválida, se omite:", line)
			continue
		}
		digimon := models.Digimon{
			Name:      line[0],
			Type:      line[1],
			Sacrifice: false,
		}
		digimons = append(digimons, digimon)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return digimons, nil
}
