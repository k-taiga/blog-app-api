FROM golang:1.18

WORKDIR /app

COPY ./src ./

RUN go mod download

EXPOSE 8080
