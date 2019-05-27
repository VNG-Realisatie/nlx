FROM node:10.15.3-alpine AS build

RUN apk add git jq

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

# Add file with version identifier from git
COPY .git /go/src/go.nlx.io/nlx/insight-ui/.git
RUN ash -c 'echo "\"$(git describe --tags)\"" | jq "{tag: .}" > /tmp/version.json'

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY insight-ui/docker/default.conf /etc/nginx/conf.d/default.conf.template
COPY insight-ui/docker/nginx.conf /etc/nginx/nginx.conf

COPY --from=build /go/src/go.nlx.io/nlx/insight-ui/build /usr/share/nginx/html
COPY --from=build /tmp/version.json /usr/share/nginx/html/version.json

# Add non-privileged user
RUN adduser -D -u 1001 appuser

# Set ownership nginx.pid and cache folder in order to run nginx as non-root user
RUN touch /var/run/nginx.pid && \
    chown -R appuser /var/run/nginx.pid && \ 
    chown -R appuser /var/cache/nginx

# Set ownership of nginx/conf.d/default.conf because we will edit this file once we run nginx as non-root user
RUN touch /etc/nginx/conf.d/default.conf && \
    chown -R appuser /etc/nginx/conf.d/default.conf

RUN chown -R appuser /usr/share/nginx/html

USER appuser

COPY insight-ui/docker/startup.sh .
CMD ["./startup.sh"]

