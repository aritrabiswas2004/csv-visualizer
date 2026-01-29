#!/bin/bash

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

curl -LO https://github.com/you/csvview/releases/latest/download/csvview-$OS-$ARCH
chmod +x csvview-$OS-$ARCH

echo "Done."

# give usr instructions

echo "You can move to /usr/local or /usr/bin if you want to reuse binary everytime."
echo "Join our mailing list: https://groups.google.com/g/csv-coreutils-dev"