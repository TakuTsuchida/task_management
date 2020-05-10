FROM golang:1.13.1-alpine

WORKDIR /go/src

ENV GO111MODULE=on

RUN apk update && \
    go get github.com/pilu/fresh

EXPOSE 8082

CMD ["fresh"]