# Dockerfile
FROM golang:1.23.2

WORKDIR /Lab2/MV4/Nodo Tai

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 50058

CMD ["./main"]
