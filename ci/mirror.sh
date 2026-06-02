#!/bin/bash

REPO_PATH="${PROJECT_HOME}/cwc/"

cd "${REPO_PATH}"
git reset --hard origin/main
git push -f github main
git push -f pgitlab main
exit 0
