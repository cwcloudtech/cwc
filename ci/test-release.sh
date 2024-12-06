#!/bin/bash
set -e

echo "Starting test release process..."

#? Ensure clean state before testing
rm -rf dist || true
mkdir -p dist

docker login "${CI_REGISTRY}" --username "${CI_REGISTRY_USER}" --password "${CI_REGISTRY_PASSWORD}"

if ! docker run --rm --privileged \
  -v "$PWD:/go/src/gitlab.com/goreleaser/cwc" \
  -w "/go/src/gitlab.com/goreleaser/cwc" \
  -v "/var/run/docker.sock:/var/run/docker.sock" \
  -e DOCKER_USERNAME="${CI_REGISTRY_USER}" \
  -e DOCKER_PASSWORD="${CI_REGISTRY_PASSWORD}" \
  -e DOCKER_REGISTRY="${CI_REGISTRY}" \
  -e GITLAB_TOKEN \
  goreleaser/goreleaser release --snapshot --clean --verbose; then
  echo "❌ Test release failed"
  exit 1
fi

if [ ! -d "dist" ] || [ -z "$(ls -A dist)" ]; then
  echo "❌ No artifacts were created in dist directory"
  exit 1
fi

echo "✅ Test release completed successfully"
echo "📦 Generated artifacts:"
ls -la dist/
