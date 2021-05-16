FROM golang:1.16-alpine3.12 AS builder
RUN apk update
RUN apk add build-base
WORKDIR /go/src 
COPY . .
ENV GOPROXY=https://goproxy.cn
RUN go build -o /go/bin/kafkagoproxy main.go

FROM alpine:3.12
COPY --from=builder /go/bin/kafkagoproxy /app/
WORKDIR /app
CMD ["./kafkagoproxy"]