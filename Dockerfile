FROM golang:1.24

RUN apt-get update && apt-get install -y tree

LABEL maintainer="me@idank.dev"

WORKDIR /app/

COPY *.go /app/
COPY go.mod go.sum /app/

RUN go build -o /app/syscall-bot

RUN useradd -m container
USER container

WORKDIR /home/container/

RUN tree
CMD ["./syscall-bot"]

