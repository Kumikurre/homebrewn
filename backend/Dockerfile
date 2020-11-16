FROM golang:latest AS builder

RUN mkdir -p /go/src/builder/back

ADD ./main.go /go/src/builder/back
ADD ./go.mod /go/src/builder/back
ADD ./go.sum /go/src/builder/back

WORKDIR /go/src/builder/back

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /go/bin/main

FROM alpine:3.11.2

RUN apk add --no-cache ca-certificates openssl

COPY --from=builder /go/bin/main /go/bin/main

EXPOSE 8080

ENTRYPOINT [ "/go/bin/main" ]
