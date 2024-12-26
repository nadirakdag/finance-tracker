FROM golang:1.23.2-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN apk add --no-cache gcc musl-dev
RUN CGO_ENABLED=1 GOOS=linux go build -o server ./cmd/server

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/server .
COPY --from=builder /app/configs/config.yaml ./configs/config.yaml

EXPOSE 8080
CMD ["./server"]