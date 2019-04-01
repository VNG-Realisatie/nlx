FROM node:10.15.0-alpine AS build

RUN apk add git

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
RUN ash -c 'git describe --tags > /tmp/version'

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY txlog-ui/docker/default.conf /etc/nginx/conf.d/default.conf

COPY --from=build /go/src/go.nlx.io/nlx/txlog-ui/build /usr/share/nginx/html
COPY --from=build /tmp/version /usr/share/nginx/html/version
