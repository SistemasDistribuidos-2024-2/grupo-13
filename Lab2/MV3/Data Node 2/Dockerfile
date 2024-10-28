# Dockerfile
FROM golang:1.23.2

WORKDIR /Lab2/MV3/Data Node 2

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 50052

CMD ["./main", "2"]
