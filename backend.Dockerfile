FROM golang:latest

WORKDIR /yyblog_backend

COPY . .

RUN go mod tidy

RUN go build -o yyblog_backend .

CMD ["./yyblog_backend", "start"]