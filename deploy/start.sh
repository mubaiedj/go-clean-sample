#!/usr/bin/env bash
set -e
source .env

echo "MY IP LOCAL: ${MY_IP}"

docker run -d \
  --name=cockroachdb \
  -p 26257:26257 -p 8180:8080 \
  -v "${PWD}/cockroach-data/cockroachdb:/cockroach/cockroach-data" \
  cockroachdb/cockroach:v19.2.2 start \
  --insecure