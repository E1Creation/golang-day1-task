FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o app.exe

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app.exe .
EXPOSE 8080
CMD ["./app.exe"]
