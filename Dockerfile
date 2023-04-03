# syntax=docker/dockerfile:1
FROM golang:latest

WORKDIR /app

COPY . /app

RUN go build .

EXPOSE 1337

CMD [ "./vm" ]
