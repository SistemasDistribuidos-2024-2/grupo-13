# Dockerfile
FROM golang:1.23.3

WORKDIR /Lab3/MV2/supervisor2

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 50052

CMD ["./main", "2"]
