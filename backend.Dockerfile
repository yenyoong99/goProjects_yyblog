FROM golang:latest

WORKDIR /backend

COPY . .

RUN go mod tidy

RUN go build -o backend .

CMD ["./backend", "start"]