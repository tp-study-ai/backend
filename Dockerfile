FROM golang:alpine as build

COPY . /project

WORKDIR /project

RUN apk add make git && make build

EXPOSE 8000

CMD ./main