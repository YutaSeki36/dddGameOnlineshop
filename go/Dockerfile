FROM golang:1.14.9-alpine3.11 AS build

WORKDIR /
RUN mkdir -p /go/src/github.com/dddGameOnlineshop/go
COPY . /go/src/github.com/dddGameOnlineshop/go
RUN apk update \
  && apk add --no-cache git \
  && apk add --no-cache make \
  && apk add --no-cache gcc
RUN cd /go/src/github.com/dddGameOnlineshop/go && export GO111MODULE=on && export CGO_ENABLED=0 && make all

FROM alpine:3.8
COPY --from=build /go/src/github.com/dddGameOnlineshop/go /usr/local/bin/
CMD ["app"]