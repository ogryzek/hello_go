FROM golang

ADD . /go/src/github.com/ogryzek/hello_go
RUN go install github.com/ogryzek/hello_go
ENTRYPOINT /go/bin/hello_go

EXPOSE 3000
