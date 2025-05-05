# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Salin go mod files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Salin source code
COPY . .

# Build aplikasi
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Runtime
FROM alpine:3.21

WORKDIR /app

# Salin binary dari stage build
COPY --from=builder /app/main .

# Expose port yang digunakan aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]