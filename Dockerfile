FROM golang:latest

WORKDIR /yyblog

COPY . .

WORKDIR /yyblog/backend

RUN go mod tidy

RUN go build -o yyblog .

CMD ["./yyblog", "start"]