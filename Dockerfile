FROM alpine:latest
WORKDIR /app
COPY /dist/main .
EXPOSE 8080
CMD ["./main"]