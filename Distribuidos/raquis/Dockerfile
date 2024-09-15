# syntax=docker/dockerfile:1
FROM golang:1.23

WORKDIR /Distribuidos/raquis


COPY . .
RUN go mod tidy

RUN go build -o main .

EXPOSE 5672

CMD [ "./main" ]