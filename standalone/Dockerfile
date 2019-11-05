FROM quay.io/redhatqe/selenium-base
LABEL maintainer="dmisharo@redhat.com"

ARG FIREFOX_VERSION
ARG GECKODRIVER_VERSION

ENV SELENIUM_MAJOR_VERSION=3 \
    SELENIUM_MINOR_VERSION=141 \
    SELENIUM_PATCH_VERSION=59 \
    SELENIUM_PORT=4444 \
    VNC_PORT=5999 \
    FIREFOX_VERSION=${FIREFOX_VERSION:-"68.2.0esr"} \
    GECKODRIVER_VERSION=${GECKODRIVER_VERSION:-"v0.26.0"} \
    DISPLAY=:99 \
    HOME=$SELENIUM_HOME

ENV SELENIUM_VERSION=$SELENIUM_MAJOR_VERSION.$SELENIUM_MINOR_VERSION.$SELENIUM_PATCH_VERSION \
    SELENIUM_PATH=$SELENIUM_HOME/selenium-server/selenium-server-standalone.jar \
    PATH=$SELENIUM_HOME/firefox:/opt/google/chrome:$PATH

# selenium server port
EXPOSE $SELENIUM_PORT

# vnc port
EXPOSE $VNC_PORT

# selenium server
ADD http://selenium-release.storage.googleapis.com/$SELENIUM_MAJOR_VERSION.$SELENIUM_MINOR_VERSION/selenium-server-standalone-$SELENIUM_VERSION.jar $SELENIUM_PATH

# chrome
RUN curl -LO https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm && \
    rpm -i google-chrome-stable_current_x86_64.rpm && \
    rm -f google-chrome-stable_current_x86_64.rpm

# chrome and chrome driver versions should match in order to avoid incompatibility
RUN CHROME_VERSION=$(rpm -q --qf "%{VERSION}\n" google-chrome-stable | sed -Ee 's/^(.*)\..*/\1/') && \
    CHROME_DRIVER_VERSION=$(curl -s https://chromedriver.storage.googleapis.com/LATEST_RELEASE_$CHROME_VERSION) && \
    curl -O https://chromedriver.storage.googleapis.com/$CHROME_DRIVER_VERSION/chromedriver_linux64.zip && \
    unzip -d /usr/bin/ chromedriver_linux64.zip && \
    chmod a+x /usr/bin/chromedriver && \
    rm -f chromedriver_linux64.zip

# firefox
RUN curl -LO https://download-installer.cdn.mozilla.net/pub/firefox/releases/$FIREFOX_VERSION/linux-x86_64/en-US/firefox-$FIREFOX_VERSION.tar.bz2 && \
    tar -C . -xjvf firefox-$FIREFOX_VERSION.tar.bz2 && \
    rm -f firefox-$FIREFOX_VERSION.tar.bz2

# gecko for FF
RUN curl -LO https://github.com/mozilla/geckodriver/releases/download/$GECKODRIVER_VERSION/geckodriver-$GECKODRIVER_VERSION-linux64.tar.gz && \
    tar -C /usr/bin/ -xvf geckodriver-$GECKODRIVER_VERSION-linux64.tar.gz && \
    rm -f geckodriver-$GECKODRIVER_VERSION-linux64.tar.gz

# Add the xstartup file into the image and add config.
COPY ./xstartup ./config .vnc/

RUN chmod a+x $SELENIUM_HOME/.vnc/xstartup && \
    chgrp -R 0 $SELENIUM_HOME && \
    chmod -R g=u $SELENIUM_HOME

USER 1001

ENTRYPOINT /usr/bin/dumb-init vncserver $DISPLAY -fg -Log *:stderr:100
