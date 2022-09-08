# syntax=docker/dockerfile:1

FROM golang:1.19.1-bullseye
WORKDIR /app
COPY fibserver/go.mod ./
RUN go mod download
COPY fibserver/*.go ./
RUN go build -o /fibserver
EXPOSE 8080
CMD [ "/fibserver" ]