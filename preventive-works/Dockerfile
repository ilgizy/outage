FROM golang:1.17-alpine

WORKDIR /app

COPY . /app

RUN go build -o preventive-works ./cmd

CMD ["./preventive-works"]