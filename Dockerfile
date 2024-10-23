FROM golang:alpine AS builder

WORKDIR /

COPY . .

RUN go build -o ./bin/bot cmd/main.go

FROM alpine:latest

COPY --from=0 /configs ./configs/
COPY --from=0 /bin/bot .

CMD ["./bot", "-c", "configs/config.yml"]
