FROM golang:1.11-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git

# Force the go compiler to use modules
ENV GO111MODULE=on
ENV CGO_ENABLED=0
