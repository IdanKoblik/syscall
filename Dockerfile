FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o syscall-bot

FROM alpine:latest

RUN apk add --no-cache ca-certificates tree

RUN adduser -D container

WORKDIR /home/container

COPY --from=builder /app/syscall-bot ./syscall-bot

RUN tree /

USER container

CMD ["./syscall-bot"]

