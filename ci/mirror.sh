#!/bin/bash

REPO_PATH="${PROJECT_HOME}/cwc/"

cd "${REPO_PATH}"
git reset --hard origin/main
git push github main 
git push pgitlab main
exit 0
