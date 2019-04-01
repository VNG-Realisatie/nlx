FROM node:10.15.0-alpine AS build

RUN apk add git

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

# Add file with version identifier from git
COPY .git /go/src/go.nlx.io/nlx/directory-ui/.git
RUN ash -c 'git describe --tags > /tmp/version'

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY directory-ui/docker/default.conf /etc/nginx/conf.d/default.conf

COPY --from=build /go/src/go.nlx.io/nlx/directory-ui/build /usr/share/nginx/html
COPY --from=build /tmp/version /usr/share/nginx/html/version
