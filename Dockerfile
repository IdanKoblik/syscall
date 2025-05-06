FROM golang:1.24

LABEL maintainer="me@idank.dev"

WORKDIR /app

COPY *.go ./
COPY go.mod go.sum ./

RUN go build -o syscall-bot

RUN useradd -m container
USER container

WORKDIR /home/container/

CMD ["./syscall-bot"]

