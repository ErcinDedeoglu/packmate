FROM golang:1.23 AS builder
WORKDIR /app
COPY archiver/go.mod .
COPY archiver/archiver.go .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o archiver

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/archiver /usr/local/bin/
ENTRYPOINT ["archiver"]