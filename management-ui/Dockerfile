FROM node:12.16.2-alpine AS build

RUN apk add --update git jq

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

# Create version.json
FROM alpine:3.11 AS version

RUN apk add --update jq

ARG GIT_TAG_NAME=undefined
ARG GIT_COMMIT_HASH=undefined
RUN jq -ncM --arg tag $GIT_TAG_NAME --arg commit $GIT_COMMIT_HASH  '{tag: $tag, commit: $commit}' | tee /version.json

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY docker/nginx.conf /etc/nginx/nginx.conf
COPY docker/default.conf /etc/nginx/conf.d/default.conf

# Add non-privileged user
RUN adduser -D -u 1001 appuser

# Set ownership nginx.pid and cache folder in order to run nginx as non-root user
RUN touch /var/run/nginx.pid && \
    chown -R appuser /var/run/nginx.pid && \
    chown -R appuser /var/cache/nginx

COPY --from=build /go/src/go.nlx.io/nlx/management-ui/build /usr/share/nginx/html
COPY --from=version /version.json /usr/share/nginx/html

USER appuser
