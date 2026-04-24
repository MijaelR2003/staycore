FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o staycore-api .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/staycore-api .

EXPOSE 8080

CMD ["./staycore-api"]