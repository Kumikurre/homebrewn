
FROM golang:1.15.5 AS base

ENV GO111MODULE=on
WORKDIR /app
COPY . .

FROM base AS dev

RUN go get -u github.com/cosmtrek/air
ENTRYPOINT [ "air" ]

FROM base AS builder

RUN CGO_ENABLED=0 go build

FROM alpine:3.11.2 as prod

RUN apk add --no-cache ca-certificates openssl
COPY --from=builder /app/backend /go/bin/backend
EXPOSE 8080
ENTRYPOINT [ "/go/bin/backend" ]
