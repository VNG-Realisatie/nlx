# Copyright © VNG Realisatie 2022
# Licensed under the EUPL

FROM alpine:3.18.5

RUN apk add --no-cache sniproxy

COPY ./entrypoint.sh /usr/share/sniproxy/entrypoint.sh

ENTRYPOINT ["/usr/share/sniproxy/entrypoint.sh"]
CMD ["sniproxy", "-f", "-c", "/etc/sniproxy-generated.conf"]
