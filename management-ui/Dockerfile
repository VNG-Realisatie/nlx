FROM node:10.16.3-alpine AS build

RUN apk add git jq

ENV CI=true

# Only copy package.json and package-lock.json so we can use Docker's caching mechanism.
COPY management-ui/package.json \
    management-ui/package-lock.json \
    /go/src/go.nlx.io/nlx/management-ui/

WORKDIR /go/src/go.nlx.io/nlx/management-ui
RUN npm ci --no-progress --color=false --quiet

# Now copy the whole workdir for the build step.
COPY management-ui /go/src/go.nlx.io/nlx/management-ui

RUN npm run build

# Add file with version tag from git
COPY .git /go/src/go.nlx.io/nlx/management-ui/.git
RUN ash -c 'echo "\"$(git describe --tags)\"" | jq "{tag: .}" > /tmp/version.json'

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY management-ui/docker/default.conf /etc/nginx/conf.d/default.conf
COPY management-ui/docker/nginx.conf /etc/nginx/nginx.conf

COPY --from=build /go/src/go.nlx.io/nlx/management-ui/build /usr/share/nginx/html
COPY --from=build /tmp/version.json /usr/share/nginx/html/version.json

# Add non-privileged user
RUN adduser -D -u 1001 appuser

# Set ownership nginx.pid and cache folder in order to run nginx as non-root user
RUN touch /var/run/nginx.pid && \
    chown -R appuser /var/run/nginx.pid && \
    chown -R appuser /var/cache/nginx

USER appuser
