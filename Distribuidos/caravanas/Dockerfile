# Dockerfile
FROM golang:1.23.1

WORKDIR /Distribuidos/caravanas

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 50052

CMD ["./main"]
