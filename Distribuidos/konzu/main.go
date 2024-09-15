package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	pb "grpc-server/proto/grpc-server/proto"
)

type Paquete struct {
	Timestamp     time.Time
	IdPaquete     string
	Tipo          string
	Contenido     string
	Valor         float64
	Escolta       string
	Destino       string
	IdSeguimiento int
	Estado        string
	Intentos      int
}

var (
	ostrionitaQueue   []Paquete
	prioritarioQueue  []Paquete
	normalQueue       []Paquete
	todosPaquetes     []Paquete
	trackingIdCounter int
	mu                sync.Mutex
	clientConn        *grpc.ClientConn
	caravanClient     pb.CaravanServiceClient
	noActivity        bool
	lastActivityTime  time.Time
	rabbitConn        *amqp.Connection
	rabbitChannel     *amqp.Channel
)

type clientServiceServer struct {
	pb.UnimplementedClientServiceServer
}

func (s *clientServiceServer) CreateOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	log.Printf("Recibida nueva orden: ID %s, Tipo %s, Destino %s", req.PackageId, req.TypeOrder, req.Destination)

	paquete := Paquete{
		IdPaquete: req.PackageId,
		Tipo:      req.TypeOrder,
		Contenido: req.Content,
		Valor:     float64(req.Price),
		Escolta:   req.Escort,
		Destino:   req.Destination,
	}

	trackingId := recibirPaquete(paquete)
	log.Printf("Paquete asignado con ID de seguimiento: %d", trackingId)

	return &pb.OrderResponse{TrackingId: strconv.Itoa(trackingId)}, nil
}

func (s *clientServiceServer) CheckStatus(ctx context.Context, req *pb.TrackingRequest) (*pb.StatusResponse, error) {
	log.Printf("Verificando estado para Tracking ID: %s", req.TrackingId)

	mu.Lock()
	defer mu.Unlock()

	for _, paquete := range todosPaquetes {
		if strconv.Itoa(paquete.IdSeguimiento) == req.TrackingId {
			log.Printf("Estado del paquete %s: %s, Intentos: %d", paquete.IdPaquete, paquete.Estado, paquete.Intentos)
			return &pb.StatusResponse{
				Status:   paquete.Estado,
				Attempts: int32(paquete.Intentos),
			}, nil
		}
	}

	log.Printf("No se encontró el paquete con Tracking ID: %s", req.TrackingId)
	return &pb.StatusResponse{Status: "No encontrado", Attempts: 0}, nil
}

func recibirPaquete(paquete Paquete) int {
	mu.Lock()
	defer mu.Unlock()

	paquete.Timestamp = time.Now()
	paquete.IdSeguimiento = trackingIdCounter
	paquete.Estado = "En cetus"
	paquete.Intentos = 0
	trackingIdCounter++

	switch paquete.Tipo {
	case "Ostronitas":
		log.Printf("Paquete %s agregado a la cola de ostronitas", paquete.IdPaquete)
		ostrionitaQueue = append(ostrionitaQueue, paquete)
	case "Prioritario":
		log.Printf("Paquete %s agregado a la cola prioritaria", paquete.IdPaquete)
		prioritarioQueue = append(prioritarioQueue, paquete)
	default:
		log.Printf("Paquete %s agregado a la cola normal", paquete.IdPaquete)
		normalQueue = append(normalQueue, paquete)
	}

	todosPaquetes = append(todosPaquetes, paquete)
	noActivity = false
	lastActivityTime = time.Now()

	return paquete.IdSeguimiento
}

func enviarPaquete(paquete Paquete, tipoCaravana string) error {
	mu.Lock()
	defer mu.Unlock()

	log.Printf("Enviando paquete %s a caravana %s", paquete.IdPaquete, tipoCaravana)
	for i, p := range todosPaquetes {
		if p.IdPaquete == paquete.IdPaquete {
			todosPaquetes[i].Estado = "En camino"
			break
		}
	}

	req := &pb.PackageRequest{
		PackageId:   paquete.IdPaquete,
		PackageType: paquete.Tipo,
		Destination: paquete.Destino,
		Value:       float32(paquete.Valor),
	}

	_, err := caravanClient.AssignPackage(context.Background(), req)
	if err != nil {
		log.Printf("Error al asignar paquete %s: %v", paquete.IdPaquete, err)
		return err
	}

	log.Printf("Paquete %s asignado correctamente", paquete.IdPaquete)
	return nil
}

func verificarCaravanas() {
	for {
		time.Sleep(1 * time.Second)
		log.Println("Verificando disponibilidad de caravanas...")

		resp, err := caravanClient.CheckCaravanStatus(context.Background(), &pb.EmptyRequest{})
		if err != nil {
			log.Printf("Error al verificar estado de las caravanas: %v", err)
			continue
		}

		if len(ostrionitaQueue) > 0 && (resp.Caravan_P_1 || resp.Caravan_P_2) {
			paquete := ostrionitaQueue[0]
			log.Printf("Caravana prioritaria disponible, intentando enviar paquete %s", paquete.IdPaquete)
			if err := enviarPaquete(paquete, "Prioritaria"); err == nil {
				ostrionitaQueue = ostrionitaQueue[1:]
			}
		} else if len(prioritarioQueue) > 0 && (resp.Caravan_P_1 || resp.Caravan_P_2 || resp.Caravan_N) {
			paquete := prioritarioQueue[0]
			log.Printf("Caravana cualquiera disponible, intentando enviar paquete %s", paquete.IdPaquete)
			if err := enviarPaquete(paquete, "Cualquiera"); err == nil {
				prioritarioQueue = prioritarioQueue[1:]
			}
		} else if len(normalQueue) > 0 && resp.Caravan_N {
			paquete := normalQueue[0]
			log.Printf("Caravana normal disponible, intentando enviar paquete %s", paquete.IdPaquete)
			if err := enviarPaquete(paquete, "Normal"); err == nil {
				normalQueue = normalQueue[1:]
			}
		}
	}
}

type caravanServiceServer struct {
	pb.UnimplementedCaravanServiceServer
}

func (s *caravanServiceServer) ReportStatus(ctx context.Context, req *pb.StatusRequestCaravana) (*pb.StatusResponseCaravana, error) {
	mu.Lock()
	defer mu.Unlock()

	for i, paquete := range todosPaquetes {
		if paquete.IdPaquete == req.PackageId {
			todosPaquetes[i].Estado = req.Status
			todosPaquetes[i].Intentos = int(req.Attempts)
			log.Printf("Paquete %s actualizado con estado: %s e intentos: %d", req.PackageId, req.Status, req.Attempts)

			log.Println(todosPaquetes[i])
			go enviarMensajeAFinanzas(todosPaquetes[i], false)

			break
		}
	}

	if verificarPaquetesFinalizados() {
		log.Println("Mensaje Final enviado")
		go enviarMensajeAFinanzas(Paquete{}, true)
	}

	return &pb.StatusResponseCaravana{Message: "Estado actualizado con éxito"}, nil
}

func enviarMensajeAFinanzas(paquete Paquete, finalizar bool) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Printf("Error al conectar a RabbitMQ: %v", err)
		return
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Error al abrir un canal en RabbitMQ: %v", err)
		return
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {

		}
	}(ch)

	err = ch.Confirm(false)
	if err != nil {
		return
	}

	msg := map[string]interface{}{
		"valor":    paquete.Valor,
		"intentos": paquete.Intentos,
		"estado":   paquete.Estado,
		"tipo":     paquete.Tipo,
		"id":       paquete.IdPaquete,
	}
	if finalizar {
		msg["id"] = "-1"
	}

	body, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error al serializar el mensaje: %v", err)
		return
	}

	err = ch.Publish(
		"",         // exchange
		"finanzas", // routing key (queue)
		true,       // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		},
	)
	if err != nil {
		log.Printf("Error al publicar el mensaje: %v", err)
		return
	}

	log.Println("Mensaje publicado correctamente en la cola 'finanzas'")

	conf := ch.NotifyPublish(make(chan amqp.Confirmation, 1))

	go func() {
		for confirm := range conf {
			if confirm.Ack {
				log.Println("Mensaje entregado exitosamente.")
			} else {
				log.Println("Fallo en la entrega del mensaje.")
			}
		}
	}()
}

func verificarPaquetesFinalizados() bool {
	for _, paquete := range todosPaquetes {
		if paquete.Estado != "Entregado" && paquete.Estado != "No entregado" {
			return false
		}
	}
	return true
}

func initRabbitMQ() {
	var err error
	rabbitConn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	rabbitChannel, err = rabbitConn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	err = rabbitChannel.Confirm(false)
	if err != nil {
		return
	}
}

func startCaravanServer() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCaravanServiceServer(grpcServer, &caravanServiceServer{})

	log.Printf("Servidor de caravanas escuchando en :50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	initRabbitMQ()
	defer func(rabbitConn *amqp.Connection) {
		err := rabbitConn.Close()
		if err != nil {

		}
	}(rabbitConn)
	defer func(rabbitChannel *amqp.Channel) {
		err := rabbitChannel.Close()
		if err != nil {

		}
	}(rabbitChannel)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterClientServiceServer(grpcServer, &clientServiceServer{})

	clientConn, err = grpc.Dial("caravanas_container:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar con el servidor de caravanas: %v", err)
	}
	defer func(clientConn *grpc.ClientConn) {
		err := clientConn.Close()
		if err != nil {

		}
	}(clientConn)

	caravanClient = pb.NewCaravanServiceClient(clientConn)

	go verificarCaravanas()

	go startCaravanServer()

	go func() {
		ticker := time.NewTicker(40 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			mu.Lock()
			if time.Since(lastActivityTime) >= 2*time.Second {
				if !noActivity {
					noActivity = true
					log.Println("No se ha recibido un paquete durante 2 segundos")
				}
			}
			mu.Unlock()
		}
	}()

	log.Printf("Servidor de clientes escuchando en :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
