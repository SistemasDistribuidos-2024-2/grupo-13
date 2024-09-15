package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "grpc-server/proto/grpc-server/proto"
)

const (
	serverAddress = "konzu_container:50051"
)

type Package struct {
	PackageId   string
	TypeOrder   string
	Content     string
	Price       float32
	Escort      string
	Destination string
}

var sentPackages = struct {
	sync.Mutex
	packages []PackageTrackingInfo
}{}

type PackageTrackingInfo struct {
	Package    Package
	TrackingID string
}

func readPackagesFromFile(filename string) ([]Package, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open the file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	var packages []Package
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Split(line, ",")
		if len(fields) != 7 {
			return nil, fmt.Errorf("malformed line: %s", line)
		}

		pkg := Package{
			PackageId:   fields[0],
			TypeOrder:   fields[1],
			Content:     fields[2],
			Price:       parseFloat(fields[3]),
			Escort:      fields[4],
			Destination: fields[5],
		}

		packages = append(packages, pkg)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading the file: %v", err)
	}

	return packages, nil
}

func parseFloat(s string) float32 {
	var f float32
	if _, err := fmt.Sscanf(s, "%f", &f); err != nil {
		fmt.Println("Error parsing float:", err)
	}
	return f
}

func createOrder(client pb.ClientServiceClient, pkg Package) string {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &pb.OrderRequest{
		PackageId:   pkg.PackageId,
		TypeOrder:   pkg.TypeOrder,
		Price:       pkg.Price,
		Escort:      pkg.Escort,
		Destination: pkg.Destination,
		Content:     pkg.Content,
	}

	res, err := client.CreateOrder(ctx, req)
	if err != nil {
		log.Fatalf("Error creating order: %v", err)
	}

	return res.TrackingId
}

func sendPackagesInBackground(client pb.ClientServiceClient, packages []Package) {
	for _, pkg := range packages {
		trackingId := createOrder(client, pkg)

		sentPackages.Lock()
		sentPackages.packages = append(sentPackages.packages, PackageTrackingInfo{
			Package:    pkg,
			TrackingID: trackingId,
		})
		sentPackages.Unlock()

		time.Sleep(2 * time.Second)
	}
}

func showSentPackages() {
	sentPackages.Lock()
	defer sentPackages.Unlock()

	if len(sentPackages.packages) == 0 {
		fmt.Println("No se han enviado paquetes aún.")
		return
	}

	fmt.Println("Paquetes enviados:")
	for _, pkgInfo := range sentPackages.packages {
		fmt.Printf("Paquete: %s, Tracking ID: %s\n", pkgInfo.Package.Content, pkgInfo.TrackingID)
	}
}

func checkStatus(client pb.ClientServiceClient, trackingId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.TrackingRequest{
		TrackingId: trackingId,
	}

	res, err := client.CheckStatus(ctx, req)
	if err != nil {
		log.Fatalf("Error checking status: %v", err)
	}

	fmt.Printf("Order status with tracking %s: %s\n", trackingId, res.Status)
}

func main() {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing gRPC connection:", err)
		}
	}(conn)

	client := pb.NewClientServiceClient(conn)

	packages, err := readPackagesFromFile("packages.txt")
	if err != nil {
		log.Fatalf("Error reading packages file: %v", err)
	}

	go sendPackagesInBackground(client, packages)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nMenú:")
		fmt.Println("1) Ver paquetes enviados")
		fmt.Println("2) Ver estado de un paquete")
		fmt.Println("3) Salir")
		fmt.Print("Elige una opción: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			showSentPackages()

		case "2":
			fmt.Print("Introduce el tracking ID: ")
			scanner.Scan()
			trackingId := scanner.Text()
			checkStatus(client, trackingId)

		case "3":
			fmt.Println("Saliendo del programa...")
			return

		default:
			fmt.Println("Opción no válida, por favor elige de nuevo.")
		}
	}
}
