#!/usr/bin/env bash
set -euo pipefail

if [ -n "${GITHUB_RUN_ID-unset}" ]; then
  docker-compose -f example/docker-compose.yml up -d localstack
fi

export AWS_ACCESS_KEY_ID=foo
export AWS_SECRET_ACCESS_KEY=bar
export AWS_REGION=eu-west-1
export AWS_ENDPOINT=http://localhost:4566

QUEUE_NAME=dummy-k6-queue

timeout 22 sh -c 'until aws --endpoint-url=http://localhost:4566 sqs list-queues; do sleep 0.1 && echo "Sleeping"; done'
QUEUE_URL=$(aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name $QUEUE_NAME | jq -r '.QueueUrl')
aws --endpoint-url=http://localhost:4566 sqs purge-queue --queue-url $QUEUE_URL

go install github.com/k6io/xk6/cmd/xk6@latest
xk6 build \
    --with github.com/mridehalgh/xk6-sqs@latest=.

./k6 run example/localstack.js

echo $QUEUE_URL

JSON=$(aws --endpoint-url=http://localhost:4566 \
    --output=json \
    sqs receive-message \
    --queue-url $QUEUE_URL) \
    || die "failed to receive-message from SQS queue '$QUEUE_NAME'"

test $(echo $JSON | jq '.Messages | length' ) -eq 1 && echo "PASS" || echo "FAIL" && exit 1
