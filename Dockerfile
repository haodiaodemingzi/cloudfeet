FROM golang:latest

# set go env
ENV BUILD_DIR /opt/cloudfeet-api
ENV APP cloudfeet-api
ENV GO111MODULE on

RUN mkdir -p $BUILD_DIR
WORKDIR $BUILD_DIR

# ADD . $BUILD_DIR
COPY go.mod $BUILD_DIR
COPY go.sum $BUILD_DIR
RUN go mod download
#RUN go test ./...
ADD . $BUILD_DIR

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /usr/bin/$APP
RUN chmod +x /usr/bin/$APP

CMD ["cloudfeet-api"]
EXPOSE 8082
