FROM alpine:3.18

RUN apk add --no-cache htop

EXPOSE 9112
