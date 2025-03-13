FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway ./cmd/api-gateway/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/api-gateway .
COPY gateway-conf.yml .

EXPOSE 8082

CMD ["./api-gateway"]