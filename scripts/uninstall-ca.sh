#!/usr/bin/env bash

KEYCHAIN="/Library/Keychains/System.keychain"

function removeOldCertificates() {
  HAS_CADDY_CERTIFICATE=$(security find-certificate -Z -c Caddy "$KEYCHAIN" 2>&1)
  # security: SecKeychainSearchCopyNext: The specified item could not be found in the keychain.
  if [[ $HAS_CADDY_CERTIFICATE == *"SecKeychainSearchCopyNext"* ]]; then
    echo "Caddy certificate not found in keychain"
  else
    HASH=$(echo "$HAS_CADDY_CERTIFICATE" | grep -n "SHA-256 hash" | cut -d: -f3 | xargs)

    echo "Removing old Caddy certificate..."
    sudo security delete-certificate -t -Z "$HASH" "$KEYCHAIN"
  fi
}

removeOldCertificates
