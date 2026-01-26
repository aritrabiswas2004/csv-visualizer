#!/bin/bash

# Install script but only does minimal
# Does not move binary to /usr/local/bin
# Assumes you have git and make
# POSIX compliant (macOS & Linux)

if command -v git > /dev/null 2>&1; then
  git clone https://github.com/aritrabiswas2004/csv-visualizer.git > /dev/null
  cd csv-visualizer
else
  echo "error: this installation needs git to work (for now)."
fi

if command -v make > /dev/null 2>&1; then
  make build
  echo "=== Build completed successfully ==="
  ./viz -h
  echo "=== Above table is a temporary test.csv ==="
else
  echo "error: no make detected... (needs make for now)"
fi
