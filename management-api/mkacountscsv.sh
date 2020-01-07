#!/usr/bin/env bash
set -e

NAMES=${@:-admin readonly}

echo "id,name,password_hash,role"

for NAME in admin readonly
do
  echo "$(uuid),${NAME},\"$(go run cmd/hashpasswd/main.go <<< ${NAME})\",${NAME}"
done
