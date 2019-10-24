FROM node:10.16.3 AS build

ENV CI=true

# First copy only package.json and yarn.lock to make the dependency fetching step optional.
COPY website/package.json \
    website/package-lock.json \
    /go/src/go.nlx.io/nlx/docs/website/

WORKDIR /go/src/go.nlx.io/nlx/docs/website
RUN npm ci --no-progress --color=false --quiet

# Now copy the whole workdir for the build step.
COPY . /go/src/go.nlx.io/nlx/docs

RUN npm run build

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine
EXPOSE 8080

# Copy nginx configuration
COPY docker/default.conf /etc/nginx/conf.d/default.conf
COPY docker/nginx.conf /etc/nginx/nginx.conf

COPY --from=build /go/src/go.nlx.io/nlx/docs/website/build/nlx-docs /usr/share/nginx/html

# Add file with version tag from git
ARG GIT_TAG_NAME
ARG GIT_COMMIT_HASH

RUN ash -c 'echo "{\"tag\": \"$GIT_TAG_NAME\", \"commit\": \"$GIT_COMMIT_HASH\"}" > /usr/share/nginx/html/version.json'
