FROM golang:latest AS builder
ENV BUILD_DIR /go/src/github.com/haodiaodemingzi/cloudfeet
ENV APP cloudfeet-api
RUN mkdir -p $BUILD_DIR
ADD . $BUILD_DIR
WORKDIR $BUILD_DIR
RUN go get -v .
RUN CGO_ENABLED=0 go build -o $APP .

CMD ["$APP"]
expose 8082


