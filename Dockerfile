FROM golang:1.24.3-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /home/container

COPY --from=builder /app/main .

ENTRYPOINT ["/home/container/main"]

