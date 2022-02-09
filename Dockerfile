FROM golang:1.17.6-alpine3.15

ENV SRC_DIR=/go/src/github.com/gouser/money-boy/api

ENV GOBIN=/go/bin

WORKDIR $SRC_DIR

ADD ./api $SRC_DIR
RUN cd /go/src/;

RUN go get github.com/go-sql-driver/mysql \
    && go get -u github.com/gin-gonic/gin \
    && go get github.com/gorilla/mux \
    && go get -u github.com/jinzhu/gorm \
    && go get gopkg.in/ini.v1

ENTRYPOINT [ "go", "run", "main.go" ]