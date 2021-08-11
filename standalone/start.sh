#!/bin/sh

trap cleanup 1 2 3 6
cleanup() {
    # we would like to shutdown everything gracefully in the right order
    pkill --signal SIGTERM java && \
    pkill --signal SIGTERM fluxbox && \
    pkill --signal SIGTERM Xvnc
}

/usr/bin/Xvnc $DISPLAY \
              -alwaysshared \
              -depth 16 \
              -geometry $VNC_GEOMETRY \
              -securitytypes none \
              -auth $HOME/.Xauthority \
              -fp catalogue:/etc/X11/fontpath.d \
              -pn \
              -rfbport 59$(echo $DISPLAY| tr -d ":") \
              -rfbwait 30000 > /dev/null 2>&1 &

sleep 3

startfluxbox > /dev/null 2>&1 &

java -jar $SELENIUM_PATH -port $SELENIUM_PORT 2>&1
