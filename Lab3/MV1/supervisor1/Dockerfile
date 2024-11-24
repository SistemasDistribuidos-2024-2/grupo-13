# Dockerfile
FROM golang:1.23.3

WORKDIR /Lab3/MV1/supervisor1

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 50051

CMD ["./main", "1"]
