# syntax=docker/dockerfile:1
FROM golang:1.22.1  

COPY . /app
WORKDIR /app
RUN go mod download

WORKDIR /app/cmd/money
RUN CGO_ENABLED=0 GOOS=linux go build .
EXPOSE 8080

# Run
CMD [ "/app/cmd/money/money"]
