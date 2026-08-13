# Etapa 1: compila o binario Go
FROM golang:1.26-alpine AS builder

WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server .

# Etapa 2: imagem final, so com o binario (sem toolchain do Go)
#
# O binario fica em /usr/local/bin, FORA do WORKDIR /app. Isso e proposital:
# o volume nomeado do docker-compose (backend_data) e montado em /app pra
# persistir o tasks.json entre reinicializacoes do container. Se o binario
# tambem vivesse em /app, um volume ja existente (criado numa build antiga)
# "venceria" e continuaria servindo o binario antigo mesmo apos reconstruir
# a imagem - foi exatamente esse bug que apareceu ao testar.
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /server /usr/local/bin/server

EXPOSE 3001

CMD ["/usr/local/bin/server"]
