# Use go 1.x based on alpine image.
FROM golang:1.16.7-alpine AS build

# Install build tools.
RUN apk add --update gcc musl-dev

# Cache dependencies
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
ENV GO111MODULE on
WORKDIR /go/src/go.nlx.io/nlx
RUN go mod download

COPY auth-service /go/src/go.nlx.io/nlx/auth-service
WORKDIR /go/src/go.nlx.io/nlx/auth-service
RUN go build -o dist/bin/auth-service ./cmd/auth-service


FROM alpine:latest
COPY --from=build /go/src/go.nlx.io/nlx/auth-service/dist/bin/auth-service /usr/local/bin/auth-service
COPY --from=build /go/src/go.nlx.io/nlx/auth-service/users.csv /users.csv

# Add non-privileged user
RUN adduser -D -u 1001 appuser
USER appuser

CMD ["/usr/local/bin/auth-service"]
