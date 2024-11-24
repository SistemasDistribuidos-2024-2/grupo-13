
FROM golang:1.22.9

WORKDIR /hextech

COPY . .

RUN go mod tidy

RUN go build -o server ./cmd/server/main.go

EXPOSE 5001

CMD ["./server", "-id", "0", "-port", "50054", "-peers", "hextech2_container:50055,hextech3_container:50056"]
