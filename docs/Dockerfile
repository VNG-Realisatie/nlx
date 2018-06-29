FROM alpine:latest AS build

# Install build tools.
RUN apk add --update openssl make git

COPY . /go/src/go.nlx.io/nlx/docs
WORKDIR /go/src/go.nlx.io/nlx/docs

RUN ./install-hugo.sh

# Build docs.
RUN make

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY ./docker/default.conf /etc/nginx/conf.d/default.conf

COPY --from=build /go/src/go.nlx.io/nlx/docs/public /usr/share/nginx/html
