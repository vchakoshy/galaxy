#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /go/src/gitlab.fidibo.com/backend/galaxy
COPY go.mod $WORKDIR
COPY go.sum $WORKDIR

ENV GO111MODULE=on
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $GOPATH/bin/galaxy

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/galaxy /galaxy
EXPOSE 8080

ENV HUBBLE_REINDEXER="0"
ENV LISTEN_ADDRESS="0.0.0.0:8080"

CMD ["./galaxy", "api"]
LABEL Name=galaxy Version=1.0