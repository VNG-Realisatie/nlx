FROM node:10.16.3 AS build

ENV CI=true

# First copy only package.json and yarn.lock to make the dependency fetching step optional.
COPY docs/website/package.json \
    docs/website/package-lock.json \
    /go/src/go.nlx.io/nlx/docs/website/

WORKDIR /go/src/go.nlx.io/nlx/docs/website
RUN npm ci --no-progress --color=false --quiet

# Now copy the whole workdir for the build step.
COPY docs /go/src/go.nlx.io/nlx/docs

RUN npm run build

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY docs/docker/default.conf /etc/nginx/conf.d/default.conf
COPY docs/docker/nginx.conf /etc/nginx/nginx.conf

COPY --from=build /go/src/go.nlx.io/nlx/docs/website/build/nlx-docs /usr/share/nginx/html

# Add non-privileged user
RUN adduser -D -u 1001 appuser

# Set ownership nginx.pid and cache folder in order to run nginx as non-root user
RUN touch /var/run/nginx.pid && \
    chown -R appuser /var/run/nginx.pid && \
    chown -R appuser /var/cache/nginx

USER appuser
