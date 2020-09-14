#!/bin/sh

#
# Copyright © VNG Realisatie 2020
# Licensed under the EUPL
#

echo "window._env = { \"MANAGEMENT_API_BASE_URL\": \"${MANAGEMENT_API_BASE_URL}\", \"OIDC_BASE_URL\": \"${OIDC_BASE_URL}\" }" > /usr/share/nginx/html/env.js

exec nginx -g 'daemon off;'
