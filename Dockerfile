FROM golang:1.22-alpine AS builder

WORKDIR /build

COPY cmd ./cmd
COPY internal ./internal
COPY docs ./docs
COPY go.mod go.sum ./

RUN go mod download

WORKDIR /build/cmd/app

RUN go build -o /build/reglogauth .

FROM alpine:latest
WORKDIR /app
COPY configs/main.yml ./configs/
# Копируем бинарный файл из предыдущего этапа
COPY --from=builder /build/reglogauth /app/reglogauth

CMD ["/app/reglogauth"]
