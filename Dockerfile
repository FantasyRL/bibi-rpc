FROM golang:latest AS builder
LABEL authors="fanr"

ENV TZ Asia/Shanghai
RUN go env -w GO111MODULE=on \
  && go env -w GOPROXY=https://goproxy.cn,direct \
  && go env -w GOOS=linux \
  && go env -w GOARCH=amd64

RUN mkdir -p /app
WORKDIR /app

ADD . /app
RUN go mod tidy
RUN make build-all

