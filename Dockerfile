FROM golang:1.11-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git

