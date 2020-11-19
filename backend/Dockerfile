FROM golang:1.15.5 AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build

FROM alpine:3.11.2

RUN apk add --no-cache ca-certificates openssl

COPY --from=builder /app/backend /go/bin/backend

EXPOSE 8080

ENTRYPOINT [ "/go/bin/backend" ]
