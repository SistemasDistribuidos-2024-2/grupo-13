package grpc

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	"regional_server/pkg/crypto"
	"regional_server/pkg/data"
	"regional_server/pkg/models"

	pb "regional_server/pkg/grpc/protobuf"

	"google.golang.org/grpc"
)

// Cliente gRPC para enviar datos al Primary Node
type Client struct {
	conn     *grpc.ClientConn
	client   pb.PrimaryNodeServiceClient
	config   *data.InputConfig
	digimons []models.Digimon
}

func NewClient(address string, config *data.InputConfig, digimons []models.Digimon) (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewPrimaryNodeServiceClient(conn)

	return &Client{
		conn:     conn,
		client:   client,
		config:   config,
		digimons: digimons,
	}, nil
}

func (c *Client) SendEncryptedMessage(digimon models.Digimon) error {
	sacrificed := shouldSacrifice(c.config.PS)

	plaintext := digimon.Name + "," + digimon.Type + "," + strconv.FormatBool(sacrificed)

	encryptedMessage, err := crypto.EncryptAES(plaintext)
	if err != nil {
		return err
	}
	encryptedData := &pb.EncryptedMessage{
		EncryptedData: encryptedMessage,
	}
	_, err = c.client.ReceiveEncryptedMessage(context.Background(), encryptedData)
	if err != nil {
		return err
	}

	log.Printf("[CONTINENTE FOLDER] Enviado (encriptado): %s", encryptedMessage)
	return nil
}

func (c *Client) SendRandomData(count int) {
	for i := 0; i < count; i++ {
		if len(c.digimons) == 0 {
			log.Println("[CONTINENTE FOLDER] No quedan más Digimon para enviar.")
			return
		}

		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(c.digimons))

		if err := c.SendEncryptedMessage(c.digimons[index]); err != nil {
			log.Printf("Error al enviar mensaje encriptado: %v", err)
		}

		c.digimons = append(c.digimons[:index], c.digimons[index+1:]...)
	}
}

func (c *Client) Close() {
	c.conn.Close()
}

func shouldSacrifice(ps float64) bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64() <= ps
}
