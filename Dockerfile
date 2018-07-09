FROM golang:1.10-alpine

RUN apk add --update git

WORKDIR /go/src/github.com/henriqueholanda/widgets-api

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep ensure

RUN go build -o widgets-api

EXPOSE 80

CMD ["./widgets-api"]
