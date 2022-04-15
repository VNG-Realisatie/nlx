#!/bin/sh
HOST_IP=$(getent hosts host.docker.internal | awk '{print $1}' | head -1) # This is needed to resolve the correct host machine ip on both macOS, Linux and Windows

sed "s@HOST_IP@${HOST_IP}@" /etc/sniproxy.conf > /etc/sniproxy-generated.conf

exec "$@"
