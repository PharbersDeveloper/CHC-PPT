FROM golang:alpine

RUN apk add --no-cache git mercurial bash gcc g++ make pkgconfig

ENV BM_HOME /go/bin
ENV BM_KAFKA_CONF_HOME /go/bin/resource/kafkaconfig.json
ENV PKG_CONFIG_PATH /usr/lib/pkgconfig

RUN git clone https://github.com/edenhill/librdkafka.git $GOPATH/librdkafka

WORKDIR /go/librdkafka
RUN ./configure --prefix /usr  && \
make && \
make install

# 以LABEL行的变动(version的变动)来划分(变动以上)使用cache和(变动以下)不使用cache
LABEL CHC-PPT.version="1.0.8" maintainer="developer@pharbers.com"

RUN go get github.com/alfredyang1986/blackmirror && \
go get github.com/alfredyang1986/BmServiceDef && \
go get github.com/PharbersDeveloper/CHC-PPT

ADD resource /go/bin/resource

RUN go install -v github.com/PharbersDeveloper/CHC-PPT

WORKDIR /go/bin

ENTRYPOINT ["CHC-PPT"]
