# Stage 1: Build
FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go mod init go-upload-api && go mod tidy
RUN go build -o upload-api ./cmd/main.go

# Stage 2: Runtime
FROM alpine
WORKDIR /app
COPY --from=builder /app/upload-api .
CMD ["./upload-api"]
