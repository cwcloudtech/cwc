docker run --rm --privileged \
      -v $PWD:/go/src/gitlab.com/goreleaser/cwc \
      -w /go/src/gitlab.com/goreleaser/cwc \
      -v /var/run/docker.sock:/var/run/docker.sock \
      -e DOCKER_USERNAME -e DOCKER_PASSWORD -e DOCKER_REGISTRY  \
      -e GITLAB_TOKEN \
      goreleaser/goreleaser release --rm-dist