#!/bin/sh

if [[ $VNC_GEOMETRY ]]; then
    sed -i "s/geometry=.*/geometry=$VNC_GEOMETRY/g" "$HOME"/.vnc/config
fi

/usr/libexec/vncserver $DISPLAY > /dev/null 2>&1 &
java -jar $SELENIUM_PATH -port $SELENIUM_PORT 2>&1
