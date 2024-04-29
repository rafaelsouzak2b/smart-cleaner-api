FROM golang:1.21.4 AS builder
WORKDIR /app
COPY . .
RUN go clean -modcache
RUN go mod vendor
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o "dist/main"

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/dist/main .
EXPOSE 5555
CMD ["./main"]