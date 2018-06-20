FROM node:alpine AS build

# First copy only package.json and yarn.lock to make the dependency fetching step optional.
COPY ./logdb-ui/package.json \
     ./logdb-ui/yarn.lock \
     /go/src/go.nlx.io/nlx/logdb-ui/

WORKDIR /go/src/go.nlx.io/nlx/logdb-ui
RUN yarn install

# Now copy the whole directory for the build step.
COPY ./* /go/src/go.nlx.io/nlx/logdb-ui/
RUN yarn run build

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine
COPY --from=build /go/src/go.nlx.io/nlx/logdb-ui/dist /usr/share/nginx/html
