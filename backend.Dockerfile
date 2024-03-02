FROM golang:latest

WORKDIR /backend

COPY backend .

RUN go mod tidy

RUN go build -o backend .

CMD ["./backend", "start"]