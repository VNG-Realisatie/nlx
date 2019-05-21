FROM node:10.15.3-alpine AS build

RUN apk add git jq

# First copy only package.json to make the dependency fetching step optional.
COPY txlog-ui/package.json \
    txlog-ui/package-lock.json \
    /go/src/go.nlx.io/nlx/txlog-ui/

WORKDIR /go/src/go.nlx.io/nlx/txlog-ui

ARG REACT_APP_API_BASE_URL
ENV CI=true

RUN npm ci --no-progress --color=false --quiet

# Now copy the whole directory for the build step.
COPY txlog-ui /go/src/go.nlx.io/nlx/txlog-ui

RUN npm run build

# Add file with version identifier from git
COPY .git /go/src/go.nlx.io/nlx/txlog-ui/.git
RUN ash -c 'echo "\"$(git describe --tags)\"" | jq "{tag: .}" > /tmp/version.json'

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY txlog-ui/docker/default.conf /etc/nginx/conf.d/default.conf
COPY txlog-ui/docker/nginx.conf /etc/nginx/nginx.conf

COPY --from=build /go/src/go.nlx.io/nlx/txlog-ui/build /usr/share/nginx/html
COPY --from=build /tmp/version.json /usr/share/nginx/html/version.json

# Add non-privileged user
RUN adduser -D -u 1001 appuser

# Set ownership nginx.pid and cache folder in order to run nginx as non-root user
RUN touch /var/run/nginx.pid && \
    chown -R appuser /var/run/nginx.pid && \ 
    chown -R appuser /var/cache/nginx

USER appuser
