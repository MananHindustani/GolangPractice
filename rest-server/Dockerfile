## We specify the base image we need for our
## go application

FROM golang:alpine

MAINTAINER Maintainer

ENV GIN_MODE=release
ENV PORT=3004
ENV GO111MODULE=auto
## We create directory within our

RUN mkdir -p go/src/github.com/sharmamanan796/onlineBank/

## We copy everything in the root directory
## into our directory

WORKDIR  /go/src/github.com/sharmamanan796/onlineBank/

COPY . .

RUN apk update && apk add --no-cache git

RUN go get -u github.com/gin-gonic/gin

RUN go get -u github.com/lib/pq

RUN go mod init main

ENV GOBIN /go/bin

ENV GOROOT /usr/local/go

ENV GOPATH /go

ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH

## Our start command which kicks off
## our newly created binary executable
## CMD ["/go/src/github.com/sharmamanan796/onlineBank/main"]
