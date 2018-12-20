# create image from the official Go image
FROM golang:1.10.3-alpine3.8

RUN apk add --update tzdata \
    bash wget curl git;

# define work directory
ADD . /go/src/github.com/addressbook
WORKDIR /go/src/github.com/addressbook

# serve the app
CMD go test -v ./... && go run main.go
EXPOSE 8080
