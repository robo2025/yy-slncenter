FROM golang:alpine as builder

COPY src/ /opt/src
COPY Makefile /opt
WORKDIR /opt

RUN apk add --update make upx git
RUN go get -u github.com/gin-gonic/gin \
    && go get -u github.com/gin-contrib/cors \
    && go get -u github.com/sirupsen/logrus \
    && go get -u github.com/fatih/color \
    && go get -u github.com/spf13/cobra \
    && go get -u github.com/go-sql-driver/mysql \
    && go get -u github.com/jinzhu/gorm \
    && go get -u github.com/dgrijalva/jwt-go
RUN make docker


FROM alpine

COPY --from=builder /opt/slncenter /usr/local/bin/

# Refer: http://blog.cloud66.com/x509-error-when-using-https-inside-a-docker-container/
RUN apk add --no-cache --update ca-certificates

EXPOSE 9000
CMD [ "slncenter", "-l", ":9000"]
