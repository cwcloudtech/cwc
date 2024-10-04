#!/bin/bash
set -e

echo "Starting test process..."

#? Cleanup function with better error handling
cleanup() {
    echo "Cleaning up test environment..."
    
    #? Stop any running containers with our test image
    docker ps -a | grep cwc_tests && docker stop $(docker ps -a | grep cwc_tests | awk '{print $1}') || true
    
    #? Remove any existing test containers
    docker ps -a | grep cwc_tests && docker rm $(docker ps -a | grep cwc_tests | awk '{print $1}') || true
    
    #? Remove test image
    docker images | grep cwc_tests && docker rmi cwc_tests || true
    
    #? Cleanup dangling images and volumes
    docker system prune -f || true
}

#? Ensure cleanup runs even on failure
trap cleanup EXIT

echo "Building test image..."
docker build -f Dockerfile.test -t cwc_tests . || {
    echo "Failed to build test image"
    exit 1
}

echo "Running tests..."
docker run --rm \
    --name cwc_tests_run \
    -v "$(pwd):/app" \
    -w /app \
    cwc_tests \
    go test -v ./... || {
    echo "Tests failed"
    exit 1
}

echo "Tests completed successfully"
