FROM artifactory.mars.haw-hamburg.de:5000/golang:1.8-alpine

RUN apk add --update git glide
RUN git config --global http.https://gopkg.in.followRedirects true

WORKDIR /go/src/jwt-test

ADD . .

RUN go get github.com/codegangsta/gin
RUN go-wrapper download
RUN go-wrapper install

EXPOSE 8080

ENTRYPOINT sh entrypoint.sh
#ENTRYPOINT tail -f /dev/null
