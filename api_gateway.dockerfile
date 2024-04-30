FROM golang:1.21-alpine3.19 AS builder

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/api_gateway cmd/api_gateway
COPY pkg pkg
COPY templates templates

RUN go build -o gatewayApp ./cmd/api_gateway

FROM alpine:3.19

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/gatewayApp .
COPY --from=builder /app/pkg/config/envs/dev.env pkg/config/envs/dev.env
COPY templates templates

CMD ["sh","-c","echo $PORT && ./gatewayApp"]