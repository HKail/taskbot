FROM golang:alpine AS builder

ENV GOPROXY="https://goproxy.cn"

WORKDIR /build

COPY . .
RUN go build -o taskbot app/cmd/taskbot/main.go

FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
        apk update && \
        apk add tzdata ca-certificates && \
        cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
        rm -rf /var/cache/apk/*

ADD conf /build/conf
WORKDIR /build
COPY --from=builder /build/taskbot /build/taskbot

CMD ["./taskbot", "-config", "./conf/prod/config.yml"]