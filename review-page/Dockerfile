# Use go 1.x based on alpine image.
FROM golang:1.16.2-alpine AS build

WORKDIR /go/src/go.nlx.io/nlx

ENV GO111MODULE=off

# Add code that we use for building review-page
COPY . /go/src/go.nlx.io/nlx/review-page

WORKDIR /go/src/go.nlx.io/nlx/review-page

RUN go build \
    -o dist/bin/review-page

# Release binary on latest alpine image.
FROM alpine:latest

COPY --from=build /go/src/go.nlx.io/nlx/review-page/dist/bin/review-page /usr/local/bin/review-page
COPY --from=build /go/src/go.nlx.io/nlx/review-page/templates/sites.html /templates/sites.html

# Add non-privileged user
RUN adduser -D -u 1001 appuser
USER appuser

CMD ["/usr/local/bin/review-page"]
