FROM alpine:3 AS build
ARG IRMA_VERSION
ARG IRMA_CHECKSUM

RUN wget -O /usr/local/bin/irma https://github.com/privacybydesign/irmago/releases/download/v${IRMA_VERSION}/irma-master-linux-amd64 && \
    sha256sum /usr/local/bin/irma && \
    echo "${IRMA_CHECKSUM}  /usr/local/bin/irma" | sha256sum -c && \
    chmod 0755 /usr/local/bin/irma && \
    /usr/local/bin/irma version


FROM alpine:3
COPY --from=build /usr/local/bin/irma /usr/local/bin/irma

RUN apk --no-cache add ca-certificates
RUN mkdir -p /usr/local/share/irma/irma_configuration && /usr/local/bin/irma scheme download /usr/local/share/irma/irma_configuration

ENV IRMASERVER_SCHEMES_PATH /usr/local/share/irma/irma_configuration
ENTRYPOINT ["/usr/local/bin/irma"]
