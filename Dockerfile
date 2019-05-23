FROM golang:alpine

RUN apk add --no-cache --update git

ENV GOPATH=/go

ADD . $GOPATH/src/crontalk

RUN go get -u github.com/golang/dep/cmd/dep && cd $GOPATH/src/crontalk/ && dep init && dep ensure -v

RUN go install -v $GOPATH/src/crontalk


FROM alpine:latest

WORKDIR app

COPY --from=0 /go/bin/crontalk /app

ENTRYPOINT ["./crontalk"]

CMD ["serve"]

EXPOSE 8008
