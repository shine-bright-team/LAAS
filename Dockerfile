FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o laas -buildvcs=false

FROM debian:buster-slim

WORKDIR /app

COPY --from=builder /app/laas .

EXPOSE 8000

ENTRYPOINT [ "/app/laas" ]