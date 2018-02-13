FROM golang:1.9.3 as builder
RUN mkdir -p /go/src/github.com/cloudposse/slack-notifier
WORKDIR /go/src/github.com/cloudposse/slack-notifier
COPY . .
RUN go get && CGO_ENABLED=0 go build -v -o "./dist/bin/slack-notifier" *.go


FROM alpine:3.6
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/cloudposse/slack-notifier/dist/bin/slack-notifier /usr/bin/slack-notifier
ENV PATH $PATH:/usr/bin
ENTRYPOINT ["slack-notifier"]
