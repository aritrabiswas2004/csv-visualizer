# This Dockerfile is untested as of 2026-1-29
# Changes are to come!

# ---------- build stage ----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o csvview .

# ---------- runtime stage ----------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/csvview /usr/local/bin/csvview

ENTRYPOINT ["csvview"]
