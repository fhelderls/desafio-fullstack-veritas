# Etapa 1: compila o binario Go
FROM golang:1.26-alpine AS builder

WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Etapa 2: imagem final, so com o binario (sem toolchain do Go)
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /build/server .

EXPOSE 3001

CMD ["./server"]
