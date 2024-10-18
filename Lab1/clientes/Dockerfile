# Dockerfile
FROM golang:1.23

WORKDIR /Distribuidos/clientes

COPY . .
RUN go mod tidy

RUN go build -o main .

CMD ["./main"]
