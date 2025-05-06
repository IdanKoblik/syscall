FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o syscall-bot

FROM alpine:latest

RUN apk add --no-cache ca-certificates

RUN adduser -D container
USER container

WORKDIR /home/container

COPY --from=builder /app/syscall-bot /app/syscall-bot

RUN chmod +x /app/syscall-bot
CMD ["/app/syscall-bot"]

