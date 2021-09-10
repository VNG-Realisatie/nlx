# Use go 1.x based on alpine image.
FROM golang:1.17.1-alpine AS build

WORKDIR /go/src/go.nlx.io/nlx

ENV GO111MODULE=off

# Add code that we use for building apps-overview
COPY . /go/src/go.nlx.io/nlx/apps-overview

WORKDIR /go/src/go.nlx.io/nlx/apps-overview

RUN go build \
    -o dist/bin/apps-overview

FROM alpine:3.14.2

COPY --from=build /go/src/go.nlx.io/nlx/apps-overview/dist/bin/apps-overview /usr/local/bin/apps-overview
COPY --from=build /go/src/go.nlx.io/nlx/apps-overview/templates /templates

# Add non-privileged user
RUN adduser -D -u 1001 appuser
USER appuser

CMD ["/usr/local/bin/apps-overview"]
