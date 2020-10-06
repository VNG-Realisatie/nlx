#!/bin/sh
set -e

if [[ -z "${MANAGEMENT_API_BASE_URL}" ]]; then
    cp /etc/nginx/templates/default.conf.template /etc/nginx/conf.d/default.conf
else
    cat /etc/nginx/templates/reverse-proxy.conf.template | sed "s~\${MANAGEMENT_API_BASE_URL}~${MANAGEMENT_API_BASE_URL}~g" > /etc/nginx/conf.d/default.conf
fi

nginx -g "daemon off;"
