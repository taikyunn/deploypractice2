FROM golang:1.15.2

ENV GO111MODULE=on

RUN go get github.com/gin-gonic/gin && \
    go get github.com/go-sql-driver/mysql && \
    go get github.com/jinzhu/gorm

CMD cd app && go run main.go
