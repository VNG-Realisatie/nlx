FROM node:12.16.1-alpine AS build

ENV CI=true
# Don't inlne scripts, troublesome with CSP headers
# See https://create-react-app.dev/docs/advanced-configuration/
ENV INLINE_RUNTIME_CHUNK=false

# First copy package.json to make the dependency fetching step optional.
COPY package.json \
  package-lock.json \
  /go/src/go.nlx.io/nlx/insight-ui/

WORKDIR /go/src/go.nlx.io/nlx/insight-ui
RUN npm ci --no-progress --color=false --quiet

# Now copy the whole directory for the build step.
COPY . /go/src/go.nlx.io/nlx/insight-ui

# Run tests and build
RUN npm run build

# Create version.json
FROM alpine:3.11 AS version

RUN apk add --update jq

ARG GIT_TAG_NAME=undefined
ARG GIT_COMMIT_HASH=undefined
RUN jq -ncM --arg tag $GIT_TAG_NAME --arg commit $GIT_COMMIT_HASH  '{tag: $tag, commit: $commit}' | tee /version.json

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

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

# Copy nginx configuration
COPY docker/nginx.conf /etc/nginx/nginx.conf
COPY docker/default.conf /etc/nginx/conf.d/default.conf.template
COPY docker/startup.sh .
CMD ["./startup.sh"]

COPY --from=build /go/src/go.nlx.io/nlx/insight-ui/build /usr/share/nginx/html
COPY --from=version /version.json /usr/share/nginx/html

USER appuser
