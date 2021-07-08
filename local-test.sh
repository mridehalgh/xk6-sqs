#!/usr/bin/env bash
set -euo pipefail

if [[ ! -v GITHUB_RUN_ID ]]; then
  docker-compose up -f example/docker-compose.yml -d localstack
fi

timeout 22 sh -c 'until aws --endpoint-url=http://localhost:4566 sqs list-queues; do sleep 0.1 && echo "Sleeping"; done'
aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name dummy-k6-queue

go install github.com/k6io/xk6/cmd/xk6@latest
xk6 build \
    --with github.com/mridehalgh/xk6-sqs@latest=.

AWS_ENDPOINT=http://localhost:4566 ./k6 run example/localstack.js