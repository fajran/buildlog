FROM golang:1.8.1

ADD . /go/src/github.com/fajran/buildlog/
RUN cd /go/src/github.com/fajran/buildlog \
 && go get ./... \
 && make

ENV DB_MIGRATIONS_SOURCE_URI file:///go/src/github.com/fajran/buildlog/migrations
ENTRYPOINT /go/bin/buildlog
EXPOSE 8080

