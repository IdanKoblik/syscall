FROM alpine:latest

RUN apk add --no-cache --update curl ca-certificates openssl git tar bash sqlite fontconfig \
    && adduser --disabled-password --home /home/container container

WORKDIR /home/container/

ADD syscall-bot /home/container/
RUN chown container:container /home/container/syscall-bot \
    && chmod +x /home/container/syscall-bot

USER container

RUN ls -la
RUN pwd

CMD ["syscall-bot"]

