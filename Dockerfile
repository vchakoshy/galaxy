#build stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY bin/app /app
EXPOSE 8080

ENV HUBBLE_REINDEXER="1"

CMD ["./app", "api"]
LABEL Name=galaxy Version=1.0