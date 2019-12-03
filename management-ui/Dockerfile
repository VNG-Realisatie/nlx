FROM node:10.17.0-alpine AS build

RUN apk add git jq

ENV CI=true

# Only copy package.json and package-lock.json so we can use Docker's caching mechanism.
COPY package.json \
    package-lock.json \
    /go/src/go.nlx.io/nlx/management-ui/

WORKDIR /go/src/go.nlx.io/nlx/management-ui
RUN npm ci --no-progress --color=false --quiet

# Now copy the whole workdir for the build step.
COPY . /go/src/go.nlx.io/nlx/management-ui

RUN npm run build

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY docker/default.conf /etc/nginx/conf.d/default.conf
COPY docker/nginx.conf /etc/nginx/nginx.conf

COPY --from=build /go/src/go.nlx.io/nlx/management-ui/build /usr/share/nginx/html

# Add file with version tag from git
ARG GIT_TAG_NAME=undefined
ARG GIT_COMMIT_HASH=undefined
RUN ash -c 'echo "{\"tag\": \"$GIT_TAG_NAME\", \"commit\": \"$GIT_COMMIT_HASH\"}" > /usr/share/nginx/html/version.json'

# Add non-privileged user
RUN adduser -D -u 1001 appuser

# Set ownership nginx.pid and cache folder in order to run nginx as non-root user
RUN touch /var/run/nginx.pid && \
    chown -R appuser /var/run/nginx.pid && \
    chown -R appuser /var/cache/nginx

USER appuser
