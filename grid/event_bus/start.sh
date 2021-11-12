#!/usr/bin/env bash

# set -e: exit asap if a command exits with a non-zero status
set -e

if [ ! -z "$SE_EVENT_BUS_HOST" ]; then
  HOST_CONFIG="--host ${SE_EVENT_BUS_HOST}"
fi

if [ ! -z "$SE_EVENT_BUS_PORT" ]; then
  PORT_CONFIG="--port ${SE_EVENT_BUS_PORT}"
fi

if [ ! -z "$SE_OPTS" ]; then
  echo "Appending Selenium options: ${SE_OPTS}"
fi

java ${JAVA_OPTS} -jar ${SELENIUM_PATH} event-bus ${HOST_CONFIG} ${PORT_CONFIG} ${SE_OPTS}
