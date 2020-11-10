
# Dcokerfile for go-gin-wrapper
FROM golang:1.15.4

WORKDIR /homebrewn
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080
CMD ["go" , "run", "main.go"]
