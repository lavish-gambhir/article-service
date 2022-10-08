#### Build stage
FROM golang:1.19-alpine as build
LABEL maintainer = "lavish_gambhir@hotmail.com"

RUN apk update && apk add --no-cache build-base

WORKDIR /go/src/
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

COPY . ./

RUN go build -o /zara cmd/main.go

#### Deploy stage
FROM alpine:latest

WORKDIR /
COPY --from=build /zara /zara

EXPOSE 9090

ENTRYPOINT ["/zara"]