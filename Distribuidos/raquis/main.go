package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"sync"
	"time"
)

type Paquete struct {
	IdPaquete string  `json:"id"`
	Estado    string  `json:"estado"`
	Tipo      string  `json:"tipo"`
	Valor     float64 `json:"valor"`
	Intentos  int     `json:"intentos"`
}

var balanceFinal float64
var mu sync.Mutex

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@dist049:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {

		}
	}(ch)

	q, err := ch.QueueDeclare(
		"finanzas", // nombre de la cola
		true,       // durable
		false,      // autoDelete
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	file, err := os.OpenFile("finanzas.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	for {
		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			false,  // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		if err != nil {
			log.Fatalf("Failed to register a consumer: %v", err)
		}

		messages := make(chan bool)

		go func() {
			for d := range msgs {
				var paquete Paquete
				err := json.Unmarshal(d.Body, &paquete)
				if err != nil {
					log.Printf("Error parsing package: %v", err)
					err := d.Nack(false, true)
					if err != nil {
						return
					}
					continue
				}

				if paquete.IdPaquete == "-1" {
					log.Println("Paquete con ID -1, no se realiza ninguna acción.")
					messages <- true
					err := d.Ack(false)
					if err != nil {
						return
					}
					continue
				}

				log.Printf("Recibiendo paquete: ID: %s, Estado: %s, Tipo: %s, Valor: %.2f, Intentos: %d",
					paquete.IdPaquete, paquete.Estado, paquete.Tipo, paquete.Valor, paquete.Intentos)

				ganancia := calcularGanancias(paquete)
				balanceFinal += ganancia

				line := fmt.Sprintf("ID: %s, Estado: %s, Ganancia/Pérdida: %.2f\n", paquete.IdPaquete, paquete.Estado, ganancia)
				mu.Lock()
				_, err = file.WriteString(line)
				mu.Unlock()
				if err != nil {
					log.Printf("Error writing to file: %v", err)
				}

				err = d.Ack(false)
				if err != nil {
					return
				}

				messages <- true
			}
		}()

		select {
		case <-messages:
			log.Printf(" [*] Mensaje procesado.")
		case <-time.After(40 * time.Second):
			log.Printf("No se recibieron mensajes en 40 segundos. Generando balance final...")
			mu.Lock()
			_, err := fmt.Fprintf(file, "Balance final: %.2f créditos\n", balanceFinal)
			if err != nil {
				return
			}
			mu.Unlock()
			log.Printf("Balance final: %.2f créditos\n", balanceFinal)
			return
		}
	}
}

func calcularGanancias(paquete Paquete) float64 {
	var ganancia float64

	switch paquete.Tipo {
	case "Ostronitas":
		if paquete.Estado == "Entregado" {
			ganancia = paquete.Valor - float64(paquete.Intentos)*100
		} else {
			ganancia = paquete.Valor - float64(paquete.Intentos)*100
		}
	case "Prioritario":
		if paquete.Estado == "Entregado" {
			ganancia = paquete.Valor*1.3 - float64(paquete.Intentos)*100
		} else {
			ganancia = paquete.Valor*0.3 - float64(paquete.Intentos)*100
		}
	case "Normal":
		if paquete.Estado == "Entregado" {
			ganancia = paquete.Valor - float64(paquete.Intentos)*100
		} else {
			ganancia = -float64(paquete.Intentos) * 100
		}
	default:
		ganancia = 0
	}

	return ganancia
}
