FROM golang:1.17.8-alpine AS builder

ENV GO111MODULE=on
WORKDIR /build

COPY . .
RUN CGO_ENABLED=0 go build -o ca ./cmd/main.go

FROM ubuntu:20.04

WORKDIR /root

COPY --from=builder /build/ca .
COPY --from=builder /build/configs .
RUN chmod +x ca

# TLS service
CMD ["./ca", "tls", "-c", "configs/config.toml"]