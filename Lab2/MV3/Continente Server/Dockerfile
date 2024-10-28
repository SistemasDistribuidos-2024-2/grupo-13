# Usar la imagen oficial de Go 1.23.2 con Alpine
FROM golang:1.23.2-alpine

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos del proyecto al contenedor
COPY . .

# Ejecutar go mod tidy para gestionar dependencias
RUN go mod tidy

# Compilar el proyecto
RUN go build -o primary_node cmd/client/main.go

# Exponer el puerto utilizado por el Regional Server
EXPOSE 50056

# Comando de inicio del contenedor
CMD ["./regional_server"]
