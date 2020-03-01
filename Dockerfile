FROM golang:1.13-buster

WORKDIR /go/src/app

COPY . .
RUN go get
RUN go install

ENV SHELLMOCK_ENABLE=1
