FROM golang:1.17-alpine as build

COPY . /www

WORKDIR /www

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct

RUN  mkdir .target && go get -v && go build -o .target/bin

FROM alpine:latest

ENV TZ  "Asia/Shanghai"

COPY --from=build /www/.target/bin /www/.target/bin
COPY --from=build /www/conf /www/.target/conf

RUN chmod +x /www/.target/bin

WORKDIR /www/.target

ENTRYPOINT ["/bin/sh", "-l", "-c" ]

CMD [ "./bin" ]