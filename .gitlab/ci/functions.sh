#/usr/bin/env sh

start_section() {
    echo -e "section_start:$(date +%s):${1}\r\e[0K${2}"
}

stop_section() {
    echo -e "section_end:$(date +%s):${1}\r\e[0K"
}

install_npm_dependencies() {
    start_section npm_dependencies "Installing NPM dependencies"
    npm ci --cache ${NPM_CACHE_DIR} --prefer-offline --no-progress --color=false --quiet
    stop_section npm_dependencies
}

wait_for_http() {
    start_section http "Waiting for http"
    ${CI_PROJECT_DIR}/wait-for-http.sh "${1}"
    stop_section http
}
