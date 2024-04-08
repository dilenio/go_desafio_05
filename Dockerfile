FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o loadtester

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/loadtester .
ENTRYPOINT ["./loadtester"]