FROM golang:1.20-alpine

WORKDIR /code

RUN apk add htop

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
