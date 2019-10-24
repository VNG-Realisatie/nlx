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

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY management-ui/docker/default.conf /etc/nginx/conf.d/default.conf
COPY management-ui/docker/nginx.conf /etc/nginx/nginx.conf

COPY --from=build /go/src/go.nlx.io/nlx/management-ui/build /usr/share/nginx/html

# Add file with version tag from git
ARG GIT_TAG_NAME
ARG GIT_COMMIT_HASH

RUN ash -c 'echo "{\"tag\": \"$GIT_TAG_NAME\", \"commit\": \"$GIT_COMMIT_HASH\"}" > /usr/share/nginx/html/version.json'
