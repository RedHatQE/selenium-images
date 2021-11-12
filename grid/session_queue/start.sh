#!/usr/bin/env bash

# set -e: exit asap if a command exits with a non-zero status
set -e

if [[ -z "${SE_EVENT_BUS_HOST}" ]]; then
  echo "SE_EVENT_BUS_HOST not set, exiting!" 1>&2
  exit 1
fi

if [[ -z "${SE_EVENT_BUS_PUBLISH_PORT}" ]]; then
  echo "SE_EVENT_BUS_PUBLISH_PORT not set, exiting!" 1>&2
  exit 1
fi

if [[ -z "${SE_EVENT_BUS_SUBSCRIBE_PORT}" ]]; then
  echo "SE_EVENT_BUS_SUBSCRIBE_PORT not set, exiting!" 1>&2
  exit 1
fi

if [ ! -z "$SE_OPTS" ]; then
  echo "Appending Selenium options: ${SE_OPTS}"
fi

if [ ! -z "$SE_SESSION_QUEUE_HOST" ]; then
  HOST_CONFIG="--host ${SE_SESSION_QUEUE_HOST}"
fi

if [ ! -z "$SE_SESSION_QUEUE_PORT" ]; then
  PORT_CONFIG="--port ${SE_SESSION_QUEUE_PORT}"
fi

java ${JAVA_OPTS} -jar ${SELENIUM_PATH} sessionqueue \
  --publish-events tcp://"${SE_EVENT_BUS_HOST}":${SE_EVENT_BUS_PUBLISH_PORT} \
  --subscribe-events tcp://"${SE_EVENT_BUS_HOST}":${SE_EVENT_BUS_SUBSCRIBE_PORT} \
  --session-request-timeout ${SE_SESSION_REQUEST_TIMEOUT} \
  --session-retry-interval ${SE_SESSION_RETRY_INTERVAL} \
  ${HOST_CONFIG} \
  ${PORT_CONFIG} \
  ${SE_OPTS}
