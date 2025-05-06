FROM alpine:latest

RUN apk add --no-cache --update curl ca-certificates openssl git tar bash sqlite fontconfig \
    && adduser --disabled-password --home /home/container container

USER container
WORKDIR /home/container

ADD syscall-bot /home/container
RUN chmod +x /home/container/syscall-bot

RUN ls -la
RUN ls -la /home/container

CMD ["/bin/bash", "/home/container/syscall-bot"]
