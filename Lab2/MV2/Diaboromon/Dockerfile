# Dockerfile
FROM golang:1.23.2

WORKDIR /Lab2/MV2/Diaboromon

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 50055

CMD ["./main"]
