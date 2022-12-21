FROM golang:1.17-alpine3.13 AS builder

RUN go version

COPY . /test-task/
WORKDIR /test-task/

RUN go mod tidy
RUN GOOS=linux go build -o ./.bin/service ./cmd/mmrp/main.go

FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update
RUN apk upgrade
WORKDIR /root/

COPY --from=0 /test-task/.bin/service .

EXPOSE 80

CMD ["./service"]