FROM alpine:3.16.0

RUN apk add --no-cache sniproxy

COPY ./entrypoint.sh /usr/share/sniproxy/entrypoint.sh

ENTRYPOINT ["/usr/share/sniproxy/entrypoint.sh"]
CMD ["sniproxy", "-f", "-c", "/etc/sniproxy-generated.conf"]
