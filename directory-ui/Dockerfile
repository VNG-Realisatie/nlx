FROM node:alpine AS build

# First copy only package.json en package-lock.json to make the dependency fetching step optional.
COPY ./directory-ui/package.json \
     ./directory-ui/package-lock.json \
     /go/src/go.nlx.io/nlx/directory-ui/

WORKDIR /go/src/go.nlx.io/nlx/directory-ui
RUN npm install

# Now copy the whole directory for the build step.
COPY ./* /go/src/go.nlx.io/nlx/directory-ui/
RUN npm run build

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine
COPY --from=build /go/src/go.nlx.io/nlx/directory-ui/dist /usr/share/nginx/html
