FROM golang:1.24.0 AS base

WORKDIR /weather-family-service

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build cmd/main.go

CMD [ "./main" ]