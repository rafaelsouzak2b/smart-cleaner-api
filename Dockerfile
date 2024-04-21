FROM golang:1.22.2 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod vendor
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o "dist/main"

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/dist/main .
EXPOSE 8080
CMD ["./main"]