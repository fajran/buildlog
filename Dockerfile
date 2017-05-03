FROM alpine:3.5

RUN apk add --no-cache ca-certificates \
 && mkdir -p /opt/buildlog/migrations

ADD bin/buildlog.static /opt/buildlog/buildlog
ADD migrations/ /opt/buildlog/migrations/

ENV DB_MIGRATIONS_SOURCE_URI file:///opt/buildlog/migrations
ENTRYPOINT /opt/buildlog/buildlog
EXPOSE 8080

