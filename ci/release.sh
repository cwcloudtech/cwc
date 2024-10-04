#!/bin/bash
set -e

echo "Starting release process for version $CI_COMMIT_TAG..."

#? Verify this is a new version
if curl --silent "https://gitlab.com/api/v4/projects/$CI_PROJECT_ID/releases" | grep -q "\"tag_name\":\"$CI_COMMIT_TAG\""; then
  echo "❌ Release $CI_COMMIT_TAG already exists!"
  exit 1
fi

if ! docker run --rm --privileged \
  -v "$PWD:/go/src/gitlab.com/goreleaser/cwc" \
  -w "/go/src/gitlab.com/goreleaser/cwc" \
  -v "/var/run/docker.sock:/var/run/docker.sock" \
  -e DOCKER_USERNAME \
  -e DOCKER_PASSWORD \
  -e DOCKER_REGISTRY \
  -e GITLAB_TOKEN \
  goreleaser/goreleaser release --clean; then
  echo "❌ Release failed"
  exit 1
fi

echo "✅ Release completed successfully"
