#!/bin/bash
set -e

echo "Performing post-release cleanup..."
rm -rf dist || true
echo "✓ Cleanup completed"
