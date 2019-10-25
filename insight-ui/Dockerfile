FROM node:10.16.3-alpine AS build

ENV CI=true

# First copy package.json to make the dependency fetching step optional.
COPY insight-ui/package.json \
  insight-ui/package-lock.json \
  /go/src/go.nlx.io/nlx/insight-ui/

WORKDIR /go/src/go.nlx.io/nlx/insight-ui
RUN npm ci --no-progress --color=false --quiet

# Now copy the whole directory for the build step.
COPY insight-ui /go/src/go.nlx.io/nlx/insight-ui

# Run tests and build
RUN npm run build

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY insight-ui/docker/default.conf /etc/nginx/conf.d/default.conf.template
COPY insight-ui/docker/nginx.conf /etc/nginx/nginx.conf
COPY insight-ui/docker/startup.sh .

COPY --from=build /go/src/go.nlx.io/nlx/insight-ui/build /usr/share/nginx/html

# Add file with version tag from git
ARG GIT_TAG_NAME=undefined
ARG GIT_COMMIT_HASH=undefined
RUN ash -c 'echo "{\"tag\": \"$GIT_TAG_NAME\", \"commit\": \"$GIT_COMMIT_HASH\"}" > /usr/share/nginx/html/version.json'

# Add non-privileged user
RUN adduser -D -u 1001 appuser

# Set ownership to run nginx as non-root user
RUN touch /var/run/nginx.pid && \
    touch /etc/nginx/conf.d/default.conf && \
    touch /usr/share/nginx/html/env.js && \
    chown -R appuser /var/cache/nginx && \
    chown -R appuser /var/run/nginx.pid && \
    chown -R appuser /etc/nginx/conf.d/default.conf && \
    chown -R appuser /usr/share/nginx/html/env.js

USER appuser
CMD ["./startup.sh"]
