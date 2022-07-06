FROM golang:1.17-alpine

EXPOSE 8080

WORKDIR /app

COPY . /app

RUN go build -o preventive-works.exe ./cmd

CMD ["./preventive-works.exe"]