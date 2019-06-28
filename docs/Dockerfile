FROM alpine:latest AS build

# Install build tools.
RUN apk add --update openssl make git

COPY ./install-hugo.sh /install-hugo.sh
RUN /install-hugo.sh

COPY . /go/src/go.nlx.io/nlx/docs
WORKDIR /go/src/go.nlx.io/nlx/docs

# Build docs.
RUN make

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

RUN apk add --update jq

# Copy nginx configuration
COPY ./docker/default.conf /etc/nginx/conf.d/default.conf
COPY ./docker/nginx.conf  /etc/nginx/nginx.conf

COPY --from=build /go/src/go.nlx.io/nlx/docs/public /usr/share/nginx/html

# Add non-privileged user
RUN adduser -D -u 1001 appuser

# Set ownership nginx.pid and cache folder in order to run nginx as non-root user
RUN touch /var/run/nginx.pid && \
    chown -R appuser /var/run/nginx.pid && \ 
    chown -R appuser /var/cache/nginx

# Set ownership of the html directory because we will add env.js here
RUN chown -R appuser /usr/share/nginx/html

USER appuser

COPY ./docker/startup.sh .
CMD ["./startup.sh"]

