package commands

import (
	"context"
	"fmt"
	"hextech/proto"
	"strings"
)

type CommandHandler struct {
	client proto.HextechServiceClient
}

func NewCommandHandler(client proto.HextechServiceClient) *CommandHandler {
	return &CommandHandler{client: client}
}

func (h *CommandHandler) Execute(command string) {
	parts := strings.Fields(command)
	if len(parts) < 2 {
		fmt.Println("Comando inválido. Usa uno de los siguientes:")
		fmt.Println("AgregarProducto <Region> <Producto> [Valor]")
		fmt.Println("RenombrarProducto <Region> <Producto> <NuevoNombre>")
		fmt.Println("ActualizarValor <Region> <Producto> <NuevoValor>")
		fmt.Println("BorrarProducto <Region> <Producto>")
		return
	}

	switch parts[0] {
	case "AgregarProducto":
		h.addProduct(parts)
	case "RenombrarProducto":
		h.renameProduct(parts)
	case "ActualizarValor":
		h.updateProduct(parts)
	case "BorrarProducto":
		h.deleteProduct(parts)
	default:
		fmt.Println("Comando desconocido:", parts[0])
	}
}

func (h *CommandHandler) addProduct(parts []string) {
	if len(parts) < 3 {
		fmt.Println("Uso: AgregarProducto <Region> <Producto> [Valor]")
		return
	}

	region := parts[1]
	product := parts[2]
	quantity := int32(0)
	if len(parts) > 3 {
		fmt.Sscanf(parts[3], "%d", &quantity)
	}

	req := &proto.AddProductRequest{
		Region:  region,
		Product: product,
		Quantity: quantity,
	}

	resp, err := h.client.AddProduct(context.Background(), req)
	if err != nil {
		fmt.Println("Error al agregar producto:", err)
	} else {
		fmt.Printf("Producto agregado. Reloj vectorial: %v\n", resp.VectorClock)
	}
}

func (h *CommandHandler) renameProduct(parts []string) {
	if len(parts) < 4 {
		fmt.Println("Uso: RenombrarProducto <Region> <Producto> <NuevoNombre>")
		return
	}

	region := parts[1]
	oldProduct := parts[2]
	newProduct := parts[3]

	req := &proto.RenameProductRequest{
		Region:      region,
		OldProduct:  oldProduct,
		NewProduct:  newProduct,
	}

	resp, err := h.client.RenameProduct(context.Background(), req)
	if err != nil {
		fmt.Println("Error al renombrar producto:", err)
	} else {
		fmt.Printf("Producto renombrado. Reloj vectorial: %v\n", resp.VectorClock)
	}
}

func (h *CommandHandler) updateProduct(parts []string) {
	if len(parts) < 4 {
		fmt.Println("Uso: ActualizarValor <Region> <Producto> <NuevoValor>")
		return
	}

	region := parts[1]
	product := parts[2]
	var quantity int32
	fmt.Sscanf(parts[3], "%d", &quantity)

	req := &proto.UpdateProductRequest{
		Region:   region,
		Product:  product,
		Quantity: quantity,
	}

	resp, err := h.client.UpdateProduct(context.Background(), req)
	if err != nil {
		fmt.Println("Error al actualizar valor:", err)
	} else {
		fmt.Printf("Valor actualizado. Reloj vectorial: %v\n", resp.VectorClock)
	}
}

func (h *CommandHandler) deleteProduct(parts []string) {
	if len(parts) < 3 {
		fmt.Println("Uso: BorrarProducto <Region> <Producto>")
		return
	}

	region := parts[1]
	product := parts[2]

	req := &proto.DeleteProductRequest{
		Region:  region,
		Product: product,
	}

	resp, err := h.client.DeleteProduct(context.Background(), req)
	if err != nil {
		fmt.Println("Error al borrar producto:", err)
	} else {
		fmt.Printf("Producto borrado. Reloj vectorial: %v\n", resp.VectorClock)
	}
}
