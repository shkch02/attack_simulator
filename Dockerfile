FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go build -o attacker main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/attacker .
CMD ["./attacker"]
