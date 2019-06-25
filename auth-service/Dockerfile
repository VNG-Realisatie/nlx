# Use go 1.x based on the latest alpine image.
FROM nlxio/go_base AS build

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
