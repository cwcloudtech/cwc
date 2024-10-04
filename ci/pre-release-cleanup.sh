#!/bin/bash
set -e

echo "Cleaning up any existing release artifacts..."
if [ -d "dist" ]; then
  rm -rf dist
  echo "✓ Cleaned up dist directory"
fi
