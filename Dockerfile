FROM golang:alpine AS builder

WORKDIR /server

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /server/cmd

RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder server/cmd/main .

COPY --from=builder /server/internal/db/migrations /root/internal/db/migrations

COPY --from=builder /server/config/config.yaml /root/config/config.yaml

EXPOSE 8080

CMD [ "./main" ] 