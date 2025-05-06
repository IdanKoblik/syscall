FROM alpine:latest

RUN apk add --no-cache --update curl ca-certificates openssl git tar bash sqlite fontconfig \
    && adduser --disabled-password --home /home/container container

WORKDIR /app/

ADD syscall-bot /app/

RUN chmod +x /app/syscall-bot

RUN ls -la

USER container

WORKDIR /home/container/

CMD ["/bin/bash", "/app/syscall-bot"]
