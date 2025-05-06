FROM alpine:latest

RUN apk add --no-cache --update curl ca-certificates openssl git tar bash sqlite fontconfig \
    && adduser --disabled-password --home /home/container container

ADD syscall-bot /
RUN chmod +x syscall-bot

USER container
ENV USER=container HOME=/home/container


CMD ["/bin/bash", "/syscall-bot"]
