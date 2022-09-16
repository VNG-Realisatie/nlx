#!/usr/bin/env bash
# Copyright © VNG Realisatie 2022
# Licensed under the EUPL

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

find "$BASE_DIR" -name "*key.pem" -type f -exec chmod o-wr {} \;
