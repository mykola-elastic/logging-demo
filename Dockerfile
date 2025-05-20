FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o taskprocessor

FROM alpine:3.21.3

WORKDIR /root/

COPY --from=builder /app/taskprocessor .

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

CMD ["./taskprocessor"]
