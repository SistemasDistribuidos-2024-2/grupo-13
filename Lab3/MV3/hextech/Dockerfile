
FROM golang:1.22.9

WORKDIR /hextech

COPY . .

RUN go mod tidy

RUN go build -o server ./cmd/server/main.go

EXPOSE 5001

CMD ["./server", "-id", "2", "-port", "50056", "-peers", "dist049:50054,dist050:50055"]
