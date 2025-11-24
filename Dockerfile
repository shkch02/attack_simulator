FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY main.go .
COPY testpwd.txt .
RUN go build -o attacker main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/attacker .
CMD ["mkdir /tmp"]
COPY --from=builder /app/testpwd.txt /tmp/test.txt

CMD ["./attacker"]
