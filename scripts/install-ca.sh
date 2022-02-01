#!/usr/bin/env bash

KEYCHAIN="/Library/Keychains/System.keychain"
TEMP_FILE="/tmp/caddy-root.crt"
CERTIFICATE_FILE="/data/caddy/pki/authorities/local/root.crt"
CONTAINER_ID=$(docker-compose ps -q mercure)

docker cp "$CONTAINER_ID":"$CERTIFICATE_FILE" "$TEMP_FILE"

sudo security add-trusted-cert -d -k "$KEYCHAIN" "$TEMP_FILE"

rm "$TEMP_FILE"
