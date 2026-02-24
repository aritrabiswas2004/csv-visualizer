#!/bin/bash

set -e

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -r | sed 's/.*-\(.*\)/\1/')

if [[ $OS -eq "WSL2" ]]; then
  curl -LO https://github.com/aritrabiswas2004/csv-visualizer/releases/latest/download/viz-linux-$ARCH
  chmod +x viz-$OS-$ARCH
else
  curl -LO https://github.com/aritrabiswas2004/csv-visualizer/releases/latest/download/viz-$OS-$ARCH
  chmod +x viz-$OS-$ARCH
fi

echo "Binary successfully downloaded."

# /usr/local/bin or /usr/bin makes it difficult anyways to remove, and we don't
# even have a script for uninstall yet (not that it is required; this software
# is nothing short of pure art)

echo "You can move to /usr/local or /usr/bin if you want to reuse binary everytime."
echo "Join our mailing list: https://groups.google.com/g/csv-coreutils-dev"
