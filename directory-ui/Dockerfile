FROM node:10.15.0-alpine AS build

ENV CI=true

# First copy only package.json and yarn.lock to make the dependency fetching step optional.
COPY package.json \
    ./package-lock.json \
    /go/src/go.nlx.io/nlx/directory-ui/

WORKDIR /go/src/go.nlx.io/nlx/directory-ui
RUN npm ci --no-progress --color=false --quiet

# Now copy the whole workdir for the build step.
COPY . /go/src/go.nlx.io/nlx/directory-ui
RUN npm run build
RUN npm test

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY docker/default.conf /etc/nginx/conf.d/default.conf

COPY --from=build /go/src/go.nlx.io/nlx/directory-ui/build /usr/share/nginx/html
