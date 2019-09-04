FROM alpine:latest
ENV APP_PATH=/opt/cloudfeet/backend
RUN apk --no-cache add curl ca-certificates bash
RUN mkdir -p /opt/cloudfeet
COPY cloudfeet $APP_PATH
RUN chmod +x $APP_PATH
CMD ["$APP_PATH"]
expose 8082


