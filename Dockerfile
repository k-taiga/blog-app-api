FROM golang:1.18

WORKDIR /app

COPY ./src .env ./

RUN go mod download

EXPOSE 8080
