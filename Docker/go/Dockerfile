FROM golang:1.15.2

ENV GO111MODULE=on

RUN go get github.com/gin-gonic/gin
ENV PORT=${PORT}
ENTRYPOINT ["/main web"]
