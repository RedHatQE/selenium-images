#!/bin/bash

Xvnc ${DISPLAY} \
     -alwaysshared \
     -depth 16 \
     -geometry ${VNC_GEOMETRY} \
     -securitytypes none \
     -auth ${HOME}/.Xauthority \
     -fp catalogue:/etc/X11/fontpath.d \
     -pn \
     -rfbport 59$(echo ${DISPLAY} | tr -d ":") \
     -rfbwait 30000 > /dev/null 2>&1 &

sleep 3

startfluxbox > /dev/null 2>&1 &

java -jar ${SELENIUM_PATH} -port ${SELENIUM_PORT} 2>&1 &

# a hack preventing fluxbox core dump
cat
