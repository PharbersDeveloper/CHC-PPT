FROM golang:alpine

RUN apk add --no-cache git mercurial

LABEL CHC-PPT.version="1.0.1" maintainer="Pharber"

ENV BM_HOME /go/bin

RUN go get github.com/alfredyang1986/blackmirror && \
go get github.com/alfredyang1986/BmServiceDef && \
go get github.com/PharbersDeveloper/CHC-PPT

RUN go install -v github.com/PharbersDeveloper/CHC-PPT

ADD resource /go/bin/resource

WORKDIR /go/bin

ENTRYPOINT ["CHC-PPT"]
