FROM alpine:latest AS build

# Install build tools.
RUN apk add --update openssl make git

# Install hugo.
RUN wget -O - "https://github.com/gohugoio/hugo/releases/download/v0.36/hugo_0.36_Linux-64bit.tar.gz" | tar --no-same-owner -C /usr/local/bin/ -xz hugo

# Build docs.
COPY . /go/src/go.nlx.io/nlx
WORKDIR /go/src/go.nlx.io/nlx/docs
RUN make


# Copy static docs to alpine-based nginx container.
FROM nginx:alpine
COPY --from=build /go/src/go.nlx.io/nlx/docs/public /usr/share/nginx/html
