# syntax=docker/dockerfile:1
FROM golang:1.23.1

WORKDIR /Distribuidos/konzu


COPY . .
RUN go mod tidy

RUN go build -o main .

EXPOSE 50051
EXPOSE 5672
EXPOSE 50053

CMD [ "./main" ]