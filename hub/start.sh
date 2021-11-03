#!/usr/bin/env bash

# set -e: exit asap if a command exits with a non-zero status
set -e

if [ ! -z "$SE_OPTS" ]; then
  echo "Appending Selenium options: ${SE_OPTS}"
fi

if [ ! -z "$SE_HUB_HOST" ]; then
  HOST_CONFIG="--host ${SE_HUB_HOST}"
fi

if [ ! -z "$SE_HUB_PORT" ]; then
  PORT_CONFIG="--port ${SE_HUB_PORT}"
fi


java ${JAVA_OPTS} -jar ${SELENIUM_PATH} hub \
  --session-request-timeout ${SE_SESSION_REQUEST_TIMEOUT} \
  --session-retry-interval ${SE_SESSION_RETRY_INTERVAL} \
  --relax-checks ${SE_RELAX_CHECKS} \
  ${HOST_CONFIG} \
  ${PORT_CONFIG} \
  ${SE_OPTS}
