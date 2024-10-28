# Dockerfile
FROM golang:1.23.1

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 50055

CMD ["./main"]
