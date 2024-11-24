# Dockerfile
FROM golang:1.23.3

WORKDIR /Lab3/MV4/broker

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 50050

CMD ["./main"]
