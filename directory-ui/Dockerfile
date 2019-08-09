FROM node:10.16.2-alpine AS build

RUN apk add git jq

ENV CI=true

# First copy only package.json and yarn.lock to make the dependency fetching step optional.
COPY directory-ui/package.json \
    directory-ui/package-lock.json \
    /go/src/go.nlx.io/nlx/directory-ui/

WORKDIR /go/src/go.nlx.io/nlx/directory-ui
RUN npm ci --no-progress --color=false --quiet

# Now copy the whole workdir for the build step.
COPY directory-ui /go/src/go.nlx.io/nlx/directory-ui

RUN npm run build

# Add file with version tag from git
COPY .git /go/src/go.nlx.io/nlx/directory-ui/.git
RUN ash -c 'echo "\"$(git describe --tags)\"" | jq "{tag: .}" > /tmp/version.json'

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY directory-ui/docker/default.conf /etc/nginx/conf.d/default.conf
COPY directory-ui/docker/nginx.conf /etc/nginx/nginx.conf

COPY --from=build /go/src/go.nlx.io/nlx/directory-ui/build /usr/share/nginx/html
COPY --from=build /tmp/version.json /usr/share/nginx/html/version.json

# Add non-privileged user
RUN adduser -D -u 1001 appuser

# Set ownership nginx.pid and cache folder in order to run nginx as non-root user
RUN touch /var/run/nginx.pid && \
    chown -R appuser /var/run/nginx.pid && \ 
    chown -R appuser /var/cache/nginx

USER appuser
