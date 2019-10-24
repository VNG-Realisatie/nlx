FROM node:10.16.3-alpine AS build

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

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY directory-ui/docker/default.conf /etc/nginx/conf.d/default.conf
COPY directory-ui/docker/nginx.conf /etc/nginx/nginx.conf

COPY --from=build /go/src/go.nlx.io/nlx/directory-ui/build /usr/share/nginx/html

# Add file with version tag from git
ARG GIT_TAG_NAME
ARG GIT_COMMIT_HASH

RUN ash -c 'echo "{\"tag\": \"$GIT_TAG_NAME\", \"commit\": \"$GIT_COMMIT_HASH\"}" > /usr/share/nginx/html/version.json'
