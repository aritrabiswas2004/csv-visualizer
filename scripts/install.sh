#!/bin/bash

set -e

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
  x86_64) ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *)
    echo "Arch not supported: $ARCH" # In case someone still uses PowerPC in this day and age.
    exit 1
  ;;
esac

if grep -qi microsoft /proc/version; then
  echo "NOTE: It seems that you are running WSL2. The binary will be viz-linux-amd64."
fi

BINARY="viz-$OS-$ARCH"

curl -LO https://github.com/aritrabiswas2004/csv-visualizer/releases/latest/download/$BINARY
chmod +x $BINARY

echo "Binary successfully as $BINARY"

# /usr/local/bin or /usr/bin makes it difficult anyways to remove, and we don't
# even have a script for uninstall yet (not that it is required; this software
# is nothing short of pure art)

echo "You can move to /usr/local/bin or /usr/bin if you want to reuse binary globally."
echo "Join our mailing list: https://groups.google.com/g/csv-coreutils-dev"
