FROM golang:1.21.6

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go mod download