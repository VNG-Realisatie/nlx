FROM node:10.15.0-alpine AS build

ENV CI=true

# First copy package.json to make the dependency fetching step optional.
COPY ./package.json \
  ./package-lock.json \
  /go/src/go.nlx.io/nlx/insight-ui/

WORKDIR /go/src/go.nlx.io/nlx/insight-ui
RUN npm ci --no-progress --color=false --quiet

# Now copy the whole directory for the build step.
COPY . /go/src/go.nlx.io/nlx/insight-ui

# Run tests and build
RUN npm test
RUN npm run build

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY ./docker/default.conf /etc/nginx/conf.d/default.conf.template

COPY --from=build /go/src/go.nlx.io/nlx/insight-ui/build /usr/share/nginx/html

# The API base URL should be configurable at run-time. Because Nginx does not suppport using environment variables
# in the config file, we use envsubst to substitute the desired value for $BASE_API_URL in the config file.
# See https://github.com/docker-library/docs/tree/master/nginx#using-environment-variables-in-nginx-configuration
CMD envsubst < /etc/nginx/conf.d/default.conf.template '$BASE_API_URL' > /etc/nginx/conf.d/default.conf && cat /etc/nginx/conf.d/default.conf && exec nginx -g 'daemon off;'
