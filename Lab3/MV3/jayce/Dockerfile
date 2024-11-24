# Dockerfile
FROM golang:1.23.3

WORKDIR /Lab3/MV3/jayce

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 50053

CMD ["./main"]
