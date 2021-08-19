FROM docker.io/library/golang:1.17 AS builder

COPY init.go .

RUN go build init.go

FROM quay.io/redhatqe/selenium-base:latest
LABEL maintainer="dmisharo@redhat.com"

# Firefox releases
# https://download-installer.cdn.mozilla.net/pub/firefox/releases/
# Chrome versions
# https://www.ubuntuupdates.org/package/google_chrome/stable/main/base/google-chrome-stable
ARG FIREFOX_VERSION
ARG GECKODRIVER_VERSION
ARG CHROME_VERSION

ENV SELENIUM_MAJOR_VERSION=3 \
    SELENIUM_MINOR_VERSION=141 \
    SELENIUM_PATCH_VERSION=59 \
    SELENIUM_PORT=4444 \
    VNC_PORT=5999 \
    FIREFOX_VERSION=${FIREFOX_VERSION:-"78.11.0esr"} \
    CHROME_VERSION=${CHROME_VERSION:-"91.0.4472.114"} \
    GECKODRIVER_VERSION=${GECKODRIVER_VERSION:-"v0.29.1"} \
    DISPLAY=:99 \
    DBUS_SESSION_BUS_ADDRESS=/dev/null \
    HOME=${SELENIUM_HOME} \
    VNC_GEOMETRY=${VNC_GEOMETRY:-"1600x900"}

ENV SELENIUM_VERSION=${SELENIUM_MAJOR_VERSION}.${SELENIUM_MINOR_VERSION}.${SELENIUM_PATCH_VERSION} \
    SELENIUM_PATH=${SELENIUM_HOME}/selenium-server/selenium-server-standalone.jar \
    PATH=${SELENIUM_HOME}/firefox:/opt/google/chrome:${PATH}

# selenium server port
EXPOSE $SELENIUM_PORT

# vnc port
EXPOSE $VNC_PORT

# selenium server
RUN mkdir -p ${SELENIUM_HOME}/selenium-server && \
    curl -L http://selenium-release.storage.googleapis.com/${SELENIUM_MAJOR_VERSION}.${SELENIUM_MINOR_VERSION}/selenium-server-standalone-${SELENIUM_VERSION}.jar \
         -o ${SELENIUM_PATH}

# chrome
RUN curl -L https://dl.google.com/linux/chrome/rpm/stable/x86_64/google-chrome-stable-${CHROME_VERSION}-1.x86_64.rpm \
         -o google-chrome-stable-x86_64.rpm && \
    # curl -LO https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm && \
    rpm -i google-chrome-stable-x86_64.rpm && \
    rm -f google-chrome-stable-x86_64.rpm

# chrome and chrome driver versions should match in order to avoid incompatibility
RUN CHROME_VERSION=$(rpm -q --qf "%{VERSION}\n" google-chrome-stable | sed -Ee 's/^(.*)\..*/\1/') && \
    CHROME_DRIVER_VERSION=$(curl -s https://chromedriver.storage.googleapis.com/LATEST_RELEASE_${CHROME_VERSION}) && \
    curl -O https://chromedriver.storage.googleapis.com/${CHROME_DRIVER_VERSION}/chromedriver_linux64.zip && \
    unzip -d /usr/bin/ chromedriver_linux64.zip && \
    chmod a+x /usr/bin/chromedriver && \
    rm -f chromedriver_linux64.zip

# firefox
RUN curl -LO https://download-installer.cdn.mozilla.net/pub/firefox/releases/${FIREFOX_VERSION}/linux-x86_64/en-US/firefox-${FIREFOX_VERSION}.tar.bz2 && \
    tar -C . -xjvf firefox-${FIREFOX_VERSION}.tar.bz2 && \
    rm -f firefox-${FIREFOX_VERSION}.tar.bz2

# gecko for FF
RUN curl -LO https://github.com/mozilla/geckodriver/releases/download/${GECKODRIVER_VERSION}/geckodriver-${GECKODRIVER_VERSION}-linux64.tar.gz && \
    tar -C /usr/bin/ -xvf geckodriver-${GECKODRIVER_VERSION}-linux64.tar.gz && \
    rm -f geckodriver-${GECKODRIVER_VERSION}-linux64.tar.gz

COPY --from=builder /go/init /

RUN chmod +x /init && \
    chgrp -R 0 ${SELENIUM_HOME} && \
    chmod -R g=u ${SELENIUM_HOME}

USER 1001

CMD ["/init"]
