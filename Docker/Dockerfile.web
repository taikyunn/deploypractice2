FROM golang:1.15.2

ENV GO111MODULE=on

RUN go get github.com/gin-gonic/gin && \
    go get github.com/go-sql-driver/mysql && \
    go get github.com/jinzhu/gorm

# 移動したい時はcdではなくWORKDIRを設定
WORKDIR /Docker/go
CMD go run main.go
