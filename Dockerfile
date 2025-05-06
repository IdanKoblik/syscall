FROM alpine:latest

RUN apk add --no-cache --update curl ca-certificates openssl git tar bash sqlite fontconfig \
    && adduser --disabled-password --home /home/container container

USER container

WORKDIR /home/container/

ADD syscall-bot /home/container/
RUN chmod +x /home/container/syscall-bot

RUN ls -la

CMD ["/bin/bash", "/app/syscall-bot"]
