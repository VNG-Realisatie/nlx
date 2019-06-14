FROM nlxio/common AS common

# Use go 1.x based on the latest alpine image.
FROM nlxio/go_base AS build

# Add code and build.
COPY --from=common /go/src/go.nlx.io/nlx/common/logoptions  /go/src/go.nlx.io/nlx/common/logoptions
COPY . /go/src/go.nlx.io/nlx/ca-certportal
WORKDIR /go/src/go.nlx.io/nlx/ca-certportal
RUN go build -o dist/bin/nlx-ca-certportal ./cmd/nlx-ca-certportal


# Release binary on latest alpine image.
FROM alpine:latest

COPY --from=build /go/src/go.nlx.io/nlx/ca-certportal/dist/bin/nlx-ca-certportal /usr/local/bin/nlx-ca-certportal
COPY --from=build /go/src/go.nlx.io/nlx/ca-certportal/public /var/lib/nlx-ca-certportal/public

WORKDIR /var/lib/nlx-ca-certportal

# Add non-priveleged user
RUN adduser -D -u 1001 appuser
USER appuser

CMD ["/usr/local/bin/nlx-ca-certportal"]
