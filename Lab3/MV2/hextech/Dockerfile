
FROM golang:1.22.9

WORKDIR /hextech

COPY . .

RUN go mod tidy

RUN go build -o server ./cmd/server/main.go

EXPOSE 5001

CMD ["./server", "-id", "1", "-port", "50055", "-peers", "dist049:50054,dist001:50056"]
