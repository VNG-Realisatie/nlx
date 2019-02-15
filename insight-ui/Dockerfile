FROM node:10.15.0-alpine AS build

RUN apk add git

ENV CI=true

# First copy package.json to make the dependency fetching step optional.
COPY insight-ui/package.json \
  insight-ui/package-lock.json \
  /go/src/go.nlx.io/nlx/insight-ui/
COPY .git /go/src/go.nlx.io/nlx/insight-ui/.git

WORKDIR /go/src/go.nlx.io/nlx/insight-ui
RUN npm ci --no-progress --color=false --quiet

# Now copy the whole directory for the build step.
COPY insight-ui /go/src/go.nlx.io/nlx/insight-ui

# Run tests and build
RUN npm test
RUN REACT_APP_GIT_COMMIT_HASH=$(git rev-parse HEAD) && \
    REACT_APP_GIT_TAG_NAME=$(git describe --tags) && \
    npm run build

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY insight-ui/docker/default.conf /etc/nginx/conf.d/default.conf.template

COPY --from=build /go/src/go.nlx.io/nlx/insight-ui/build /usr/share/nginx/html

# The API base URL should be configurable at run-time. Because Nginx does not suppport using environment variables
# in the config file, we use envsubst to substitute the desired value for $DIRECTORY_INSPECTION_API_URL in the config file.
# See https://github.com/docker-library/docs/tree/master/nginx#using-environment-variables-in-nginx-configuration
CMD envsubst < /etc/nginx/conf.d/default.conf.template '$DIRECTORY_INSPECTION_API_URL' > /etc/nginx/conf.d/default.conf && exec nginx -g 'daemon off;'
