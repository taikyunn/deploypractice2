FROM golang:1.15.2

ENV GO111MODULE=on

RUN go get github.com/gin-gonic/gin && \
    go get github.com/go-sql-driver/mysql && \
    go get github.com/jinzhu/gorm

RUN apk update \
    && apk add --no-cache git curl

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /main ./cmd

FROM alpine:3.12

COPY --from=builder /main .

ENV PORT=${PORT}
ENTRYPOINT ["/main web"]
