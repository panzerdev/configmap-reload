FROM golang:1.7.3-alpine
MAINTAINER DaWanda <dev@dawandamail.com>

RUN mkdir -p "$GOPATH/src/github.com/panzerdev/configmap-reload"
WORKDIR $GOPATH/src/github.com/panzerdev/configmap-reload
ADD / .

RUN go install

ENTRYPOINT ["/go/bin//configmap-reload"]