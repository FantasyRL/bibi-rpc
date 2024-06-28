FROM golang:1.22 AS builder
LABEL authors="fanr"

ENV TZ Asia/Shanghai
ENV CGO_ENABLED=0
RUN go env -w GO111MODULE=on \
  && go env -w GOPROXY=https://goproxy.cn,direct \
  && go env -w GOOS=linux \
  && go env -w GOARCH=amd64

RUN mkdir -p /app
WORKDIR /app

ADD . /app
RUN go mod tidy
RUN make build-all

FROM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata

ENV TZ Asia/Shanghai
ENV service api
EXPOSE 10002
WORKDIR /app
#COPY --from=builder /app/output /app/output
COPY --from=builder /app/config /app/config
#COPY --from=builder /app/pkg /app/pkg
COPY --from=builder /app/cmd /app/cmd
#COPY --from=builder /app/kitex_gen /app/kitex_gen

#CMD ["sh","-c","./output/${service}/${service}"]
CMD ["sh","-c","sh cmd/${service}/output/bootstrap.sh"]



