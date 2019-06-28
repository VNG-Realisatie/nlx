#!/bin/sh

# the navigation URLs should be configurable at run-time.
jq -n 'env | {NAVBAR_ABOUT_PAGE_URL, NAVBAR_HOME_PAGE_URL,NAVBAR_DIRECTORY_URL}' > /usr/share/nginx/html/env.json
(echo "window._env = "; cat /usr/share/nginx/html/env.json) > /usr/share/nginx/html/env.js
rm /usr/share/nginx/html/env.json

exec nginx -g 'daemon off;'

