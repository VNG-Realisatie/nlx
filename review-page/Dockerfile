# Use go 1.x based on alpine image.
FROM golang:1.13.8-alpine AS build

WORKDIR /go/src/go.nlx.io/nlx

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
