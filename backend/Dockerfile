
# Dcokerfile for go-gin-wrapper
FROM golang:1.15.4

RUN mkdir /backend
COPY . ./backend

RUN go get -u github.com/gin-gonic/gin

RUN go build -o main ./backend/

EXPOSE 8080
CMD ["main"]
