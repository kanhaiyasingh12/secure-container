# Stage 1: Build React app
FROM node:18-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Stage 2: Serve with NGINX (patched image)
FROM nginx:stable-alpine as runtime

# Upgrade system packages to patch CVEs
RUN apk update && apk upgrade --no-cache

# Replace default config
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copy build artifacts
COPY --from=builder /app/build /usr/share/nginx/html

