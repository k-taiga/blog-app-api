FROM golang:1.18

WORKDIR /app

COPY ./src/go.mod ./src/go.sum ./

RUN go mod download

COPY ./src ./

EXPOSE 8080
