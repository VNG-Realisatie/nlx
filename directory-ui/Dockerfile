FROM node:18.5.0-alpine AS build

ENV CI=true
# Don't inlne scripts, troublesome with CSP headers
# See https://create-react-app.dev/docs/advanced-configuration/
ENV INLINE_RUNTIME_CHUNK=false

# First copy only package.json and yarn.lock to make the dependency fetching step optional.
COPY package.json \
    package-lock.json \
    /go/src/go.nlx.io/nlx/directory-ui/

WORKDIR /go/src/go.nlx.io/nlx/directory-ui
RUN npm ci --no-progress --color=false --quiet

# Now copy the whole workdir for the build step.
COPY . /go/src/go.nlx.io/nlx/directory-ui

RUN npm run build

# Create version.json
FROM alpine:3.16.0 AS version

RUN apk add --update jq

ARG GIT_TAG_NAME=undefined
ARG GIT_COMMIT_HASH=undefined
RUN jq -ncM --arg tag $GIT_TAG_NAME --arg commit $GIT_COMMIT_HASH  '{tag: $tag, commit: $commit}' | tee /version.json

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Add non-privileged user
# Note: `USER appuser` is not present in Dockerfile, because nginx.conf includes `user appuser;`
RUN adduser -D -u 1001 appuser

# Copy nginx configuration
COPY docker/nginx.conf /etc/nginx/nginx.conf
COPY docker/default.conf /etc/nginx/conf.d/default.conf

COPY --from=build /go/src/go.nlx.io/nlx/directory-ui/build /usr/share/nginx/html
COPY --from=version /version.json /usr/share/nginx/html
