FROM golang:1.22.2-bullseye as builder
ENV GO111MODULE=on
ENV CGO_ENABLED=0
WORKDIR /usr/src/
COPY . /usr/src
RUN go build -v -o "bin/slack-notifier" *.go

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/src/bin/* /usr/bin/
ENV PATH $PATH:/usr/bin
ENTRYPOINT ["slack-notifier"]
