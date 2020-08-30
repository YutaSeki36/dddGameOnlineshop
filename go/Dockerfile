FROM golang:1.11.2-alpine3.8 AS build

WORKDIR /
RUN mkdir -p /go/src/github.com/dddGameonlineshop/
COPY . /go/src/github.com/dddGameonlineshop/
RUN apk update \
  && apk add --no-cache git \
  && apk add --no-cache make
RUN cd /go/src/github.com/dddGameonlineshop/ && make all

FROM alpine:3.8
COPY --from=build /go/src/github.com/dddGameonlineshop/ /usr/local/bin/
CMD ["app"]