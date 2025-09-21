# base go image
FROM golang:1.24.2-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o authApp ./cmd/api

RUN chmod +x /app/authApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/authApp /app

CMD [ "/app/authApp" ]