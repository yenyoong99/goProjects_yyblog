FROM golang:latest

WORKDIR /yyblog

COPY . .

RUN go mod tidy

RUN go build -o yyblog .

CMD ["./yyblog", "start"]