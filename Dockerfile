FROM golang:1.21.2-alpine3.18 AS builder

COPY . /github.com/drewspitsin/auth/grpc/source/
WORKDIR /github.com/drewspitsin/auth/grpc/source/

RUN go mod download
RUN go build -o ./bin/auth cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder github.com/drewspitsin/auth/grpc/source/bin/auth .

CMD ["./auth"]