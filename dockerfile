# --- Builder stage ---
FROM golang:alpine AS builder
RUN apk update && apk add ca-certificates tzdata git
WORKDIR /app
COPY . .
WORKDIR /app/cmd/ui
RUN CGO_ENABLED=0 go build -o /go/bin/ui -ldflags '-extldflags "-static"'

# --- Final stage ---
FROM scratch
WORKDIR /app
COPY --from=builder /go/bin/ui /app/ui
COPY --from=builder /app/cmd/ui/static /app/static
COPY --from=builder /app/cmd/ui/view /app/view
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080
ENV TZ=Europe/Moscow

ENTRYPOINT ["/app/ui"]