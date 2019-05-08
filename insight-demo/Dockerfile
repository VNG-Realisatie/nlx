FROM node:alpine AS build

# First copy only package.json and yarn.lock to make the dependency fetching step optional.
COPY ./package.json \
    ./package-lock.json \
    /go/src/go.nlx.io/nlx/insight-demo/

WORKDIR /go/src/go.nlx.io/nlx/insight-demo
RUN npm install

# Now copy the whole directory for the build step.
COPY . /go/src/go.nlx.io/nlx/insight-demo
RUN npm run build

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine

# Copy nginx configuration
COPY ./docker/default.conf /etc/nginx/conf.d/default.conf
COPY ./docker/nginx.conf /etc/nginx/nginx.conf

COPY --from=build /go/src/go.nlx.io/nlx/insight-demo/build /usr/share/nginx/html

# Add non-privileged user
RUN adduser -D -u 1001 appuser

# Set ownership nginx.pid and cache folder in order to run nginx as non-root user
RUN touch /var/run/nginx.pid && \
    chown -R appuser /var/run/nginx.pid && \ 
    chown -R appuser /var/cache/nginx

USER appuser
