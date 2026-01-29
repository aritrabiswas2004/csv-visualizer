#!/bin/bash

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -r | sed 's/.*-\(.*\)/\1/')

curl -LO https://github.com/aritrabiswas2004/csv-visualizer/releases/latest/download/viz-$OS-$ARCH
chmod +x viz-$OS-$ARCH

echo "Binary downloaded."

# give usr instructions

echo "You can move to /usr/local or /usr/bin if you want to reuse binary everytime."
echo "Join our mailing list: https://groups.google.com/g/csv-coreutils-dev"