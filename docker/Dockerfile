FROM golang:1.15 as Compile

COPY . /go/src/github.com/slzcc/algae/

WORKDIR /go/src/github.com/slzcc/algae

RUN go mod download && go build

FROM ubuntu:20.04

COPY --from=Compile /go/src/github.com/slzcc/algae/algae /data/algae

CMD ["/data/algae"]
