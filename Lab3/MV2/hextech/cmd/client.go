package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"hextech/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Verifica que se proporcionen argumentos
	if len(os.Args) < 3 {
		log.Fatalf("Uso: go run client.go <server_address> <comando> [argumentos]")
	}

	serverAddress := os.Args[1] // Dirección del servidor, por ejemplo, "localhost:5001"
	command := os.Args[2]       // Comando, como "AddProduct"
	args := os.Args[3:]         // Argumentos del comando

	// Conectar al servidor gRPC
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("No se pudo conectar al servidor: %v", err)
	}
	defer conn.Close()

	client := proto.NewHextechServiceClient(conn)

	// Ejecutar el comando
	switch strings.ToLower(command) {
	case "addproduct":
		addProduct(client, args)
	case "renameproduct":
		renameProduct(client, args)
	case "updateproduct":
		updateProduct(client, args)
	case "deleteproduct":
		deleteProduct(client, args)
	case "getproduct":
		getProduct(client, args)
	default:
		log.Fatalf("Comando desconocido: %s", command)
	}
}

func addProduct(client proto.HextechServiceClient, args []string) {
	if len(args) < 3 {
		log.Fatalf("Uso: AddProduct <Region> <Producto> <Cantidad>")
	}

	region := args[0]
	product := args[1]
	var quantity int32
	fmt.Sscanf(args[2], "%d", &quantity)

	req := &proto.AddProductRequest{
		Region:   region,
		Product:  product,
		Quantity: quantity,
	}

	resp, err := client.AddProductServer(context.Background(), req)
	if err != nil {
		log.Fatalf("Error al agregar producto: %v", err)
	}

	fmt.Printf("Producto agregado. Reloj vectorial: %v\n", resp.VectorClock)
}

func renameProduct(client proto.HextechServiceClient, args []string) {
	if len(args) < 3 {
		log.Fatalf("Uso: RenameProduct <Region> <ProductoAntiguo> <ProductoNuevo>")
	}

	region := args[0]
	oldProduct := args[1]
	newProduct := args[2]

	req := &proto.RenameProductRequest{
		Region:     region,
		OldProduct: oldProduct,
		NewProduct: newProduct,
	}

	resp, err := client.RenameProductServer(context.Background(), req)
	if err != nil {
		log.Fatalf("Error al renombrar producto: %v", err)
	}

	fmt.Printf("Producto renombrado. Reloj vectorial: %v\n", resp.VectorClock)
}

func updateProduct(client proto.HextechServiceClient, args []string) {
	if len(args) < 3 {
		log.Fatalf("Uso: UpdateProduct <Region> <Producto> <Cantidad>")
	}

	region := args[0]
	product := args[1]
	var quantity int32
	fmt.Sscanf(args[2], "%d", &quantity)

	req := &proto.UpdateProductRequest{
		Region:   region,
		Product:  product,
		Quantity: quantity,
	}

	resp, err := client.UpdateProductServer(context.Background(), req)
	if err != nil {
		log.Fatalf("Error al actualizar producto: %v", err)
	}

	fmt.Printf("Producto actualizado. Reloj vectorial: %v\n", resp.VectorClock)
}

func deleteProduct(client proto.HextechServiceClient, args []string) {
	if len(args) < 2 {
		log.Fatalf("Uso: DeleteProduct <Region> <Producto>")
	}

	region := args[0]
	product := args[1]

	req := &proto.DeleteProductRequest{
		Region:  region,
		Product: product,
	}

	resp, err := client.DeleteProductServer(context.Background(), req)
	if err != nil {
		log.Fatalf("Error al borrar producto: %v", err)
	}

	fmt.Printf("Producto borrado. Reloj vectorial: %v\n", resp.VectorClock)
}

func getProduct(client proto.HextechServiceClient, args []string) {
	if len(args) < 2 {
		log.Fatalf("Uso: GetProduct <Region> <Producto>")
	}

	region := args[0]
	product := args[1]

	req := &proto.GetProductRequest{
		Region:  region,
		Product: product,
	}

	resp, err := client.GetProductServer(context.Background(), req)
	if err != nil {
		log.Fatalf("Error al obtener producto: %v", err)
	}

	fmt.Printf("Producto encontrado. Cantidad: %d, Reloj vectorial: %v\n", resp.Quantity, resp.VectorClock)
}
