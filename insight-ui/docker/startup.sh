#!/bin/sh

# The API base URL should be configurable at run-time. Because Nginx does not suppport using environment variables
# in the config file, we use envsubst to substitute the desired value for $DIRECTORY_INSPECTION_API_URL in the config file.
# See https://github.com/docker-library/docs/tree/master/nginx#using-environment-variables-in-nginx-configuration
envsubst < /etc/nginx/conf.d/default.conf.template '$DIRECTORY_INSPECTION_API_URL' > /etc/nginx/conf.d/default.conf

echo "window._env = { \"NAVBAR_HOME_PAGE_URL\": \"${NAVBAR_HOME_PAGE_URL}\" }" > /usr/share/nginx/html/env.js

exec nginx -g 'daemon off;'

