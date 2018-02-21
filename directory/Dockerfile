# Use go 1.x based on the latest alpine image.
FROM golang:1-alpine AS build

# Install build tools.
RUN apk add --update make git

# Install modd, which is used within docker-compose.
RUN go get github.com/cortesi/modd/cmd/modd

# Add code and build.
COPY . /go/src/github.com/VNG-Realisatie/nlx
WORKDIR /go/src/github.com/VNG-Realisatie/nlx/directory
RUN make


# Release binary on latest alpine image.
FROM alpine:latest
COPY --from=build /go/src/github.com/VNG-Realisatie/nlx/directory/dist/bin/nlx-directory /usr/local/bin/nlx-directory
