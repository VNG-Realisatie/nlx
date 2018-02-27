FROM node:alpine AS build

# Install node dependencies
COPY ./directory-ui/package.json \
     ./directory-ui/package-lock.json \
     /go/src/github.com/VNG-Realisatie/nlx/

WORKDIR /go/src/github.com/VNG-Realisatie/nlx/directory-ui
RUN npm install

# Build a new version
COPY ./directory-ui /go/src/github.com/VNG-Realisatie/nlx/
RUN npm run build

# Copy static docs to alpine-based nginx container.
FROM nginx:alpine
COPY --from=build /go/src/github.com/VNG-Realisatie/nlx/directory-ui/dist /usr/share/nginx/html
