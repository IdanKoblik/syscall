FROM golang:1.24

LABEL maintainer="me@idank.dev"

WORKDIR /app/

COPY *.go /app/
COPY go.mod go.sum /app/

RUN go build -o /app/syscall-bot

RUN useradd -m container
USER container

WORKDIR /home/container/

CMD ["./app/syscall-bot"]

