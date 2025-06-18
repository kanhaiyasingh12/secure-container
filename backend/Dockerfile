# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app
RUN apk --no-cache add git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags="-s -w" -o server

# Stage 2: Runtime
FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/server /
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/server"]

