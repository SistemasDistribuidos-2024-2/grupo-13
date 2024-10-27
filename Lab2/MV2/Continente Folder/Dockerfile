# Usar la última imagen oficial de Golang
FROM golang:latest AS builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos go.mod y go.sum para descargar dependencias
COPY go.mod go.sum ./

# Descargar las dependencias del proyecto
RUN go mod download

# Copiar el resto de los archivos del proyecto
COPY . .

# Compilar el proyecto
RUN go build -o regional_server cmd/client/main.go

# Usar una imagen más ligera para ejecutar el binario
FROM alpine:latest

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar el binario desde la etapa de construcción
COPY --from=builder /app/regional_server .

# Exponer el puerto utilizado por el Regional Server
EXPOSE 50056

# Comando de inicio del contenedor
CMD ["./regional_server"]
