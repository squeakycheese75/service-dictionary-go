FROM golang:1.15.6-alpine AS build

LABEL MAINTAINER "Jamie Wooltorton <james_wooltorton@hotmail.com>"
LABEL SOURCE "https://github.com/squeakycheese75/service-dictionary-go"

WORKDIR /github.com/squeakycheese75/service-dictionary-go/api/
COPY ./api/. /github.com/squeakycheese75/service-dictionary-go/api/
COPY go.* /github.com/squeakycheese75/service-dictionary-go/

RUN apk update && apk add git && apk add build-base

RUN CGO_ENABLED=1 go build -o /bin/sdapi

EXPOSE 8080

ENTRYPOINT ["/bin/sdapi"]