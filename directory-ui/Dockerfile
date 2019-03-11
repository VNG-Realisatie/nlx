FROM node:10.15.0-alpine AS build

RUN apk add git

ENV CI=true

# First copy only package.json and yarn.lock to make the dependency fetching step optional.
COPY directory-ui/package.json \
     directory-ui/package-lock.json \
    /go/src/go.nlx.io/nlx/directory-ui/
COPY .git /go/src/go.nlx.io/nlx/directory-ui/.git

WORKDIR /go/src/go.nlx.io/nlx/directory-ui
RUN npm ci --no-progress --color=false --quiet

# Now copy the whole workdir for the build step.
COPY directory-ui /go/src/go.nlx.io/nlx/directory-ui

RUN REACT_APP_GIT_COMMIT_HASH=$(git rev-parse HEAD) \
    REACT_APP_GIT_TAG_NAME=$(git describe --tags) \
    npm run build

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY directory-ui/docker/default.conf /etc/nginx/conf.d/default.conf

COPY --from=build /go/src/go.nlx.io/nlx/directory-ui/build /usr/share/nginx/html
