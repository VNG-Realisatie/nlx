FROM node:alpine AS build

# First copy only package.json and yarn.lock to make the dependency fetching step optional.
COPY ./directory-ui/package.json \
     ./directory-ui/yarn.lock \
     /go/src/go.nlx.io/nlx/directory-ui/

WORKDIR /go/src/go.nlx.io/nlx/directory-ui
RUN yarn install

# Now copy the whole directory for the build step.
COPY . /go/src/go.nlx.io/nlx
RUN yarn run build

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY ./directory-ui/docker/default.conf /etc/nginx/conf.d/default.conf

COPY --from=build /go/src/go.nlx.io/nlx/directory-ui/build /usr/share/nginx/html
