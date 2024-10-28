# Dockerfile
FROM golang:1.23.2

WORKDIR /Lab2/MV1/Data Node 1

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 50051

CMD ["./main", "1"]
