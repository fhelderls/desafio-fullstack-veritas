# Etapa 1: instala dependencias e gera os assets estaticos de producao
FROM node:24-alpine AS builder

WORKDIR /build

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
RUN npm run build

# Etapa 2: serve os assets com nginx
FROM nginx:alpine

COPY --from=builder /build/dist /usr/share/nginx/html

EXPOSE 80
