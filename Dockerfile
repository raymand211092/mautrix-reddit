# Build stage
FROM golang:1.21-alpine AS builder

# Instalar dependencias de build
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /build

# Copiar archivos de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fuente
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s" \
    -o mautrix-reddit \
    ./cmd/mautrix-reddit

# Runtime stage
FROM alpine:latest

# Instalar ca-certificates para HTTPS
RUN apk add --no-cache ca-certificates tzdata

# Crear usuario no-root
RUN addgroup -g 1000 mautrix && \
    adduser -D -u 1000 -G mautrix mautrix && \
    mkdir -p /data && \
    chown mautrix:mautrix /data

# Copiar binario
COPY --from=builder /build/mautrix-reddit /usr/local/bin/

# Copiar ejemplo de configuración
COPY --from=builder /build/example-config.yaml /opt/mautrix-reddit/

USER mautrix
WORKDIR /data

VOLUME ["/data"]

ENTRYPOINT ["/usr/local/bin/mautrix-reddit"]
CMD ["-c", "/data/config.yaml"]
