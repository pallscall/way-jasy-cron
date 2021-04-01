FROM golang:latest

MAINTAINER WayJasy 306698601@qq.com

ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn

RUN mkdir -p /go/app

WORKDIR /go/app

ADD . /go/app

RUN cd /go/app/cron/cmd && go build main.go

EXPOSE 8000
EXPOSE 9000


ENTRYPOINT ["./cron/cmd/main", "-conf=./cron/configs"]

