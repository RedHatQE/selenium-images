#!/bin/sh

trap cleanup 1 2 3 6

cleanup() {
# we would like to shutdown everything gracefully in the right order
    pkill java
    pkill fluxbox
    pkill Xvnc
}

vncserver $DISPLAY -noxstartup \
                   -securitytypes none \
                   -geometry "${VNC_GEOMETRY:-1600x900}" \
                   -depth 16 \
                   -alwaysshared &

sleep 3
xsetroot -solid grey
startfluxbox -display $DISPLAY > /dev/null 2>&1 &
java -jar $SELENIUM_PATH -port $SELENIUM_PORT 2>&1
