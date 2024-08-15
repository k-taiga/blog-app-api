FROM golang:1.17

WORKDIR /app

COPY ./src/go.mod ./

RUN go mod download

COPY ./src ./

EXPOSE 8080
