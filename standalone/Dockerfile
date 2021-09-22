FROM docker.io/library/golang:1.17 AS builder

COPY init.go .

RUN go build init.go

FROM quay.io/redhatqe/selenium-base:latest
LABEL maintainer="dmisharo@redhat.com"

# Firefox releases
# https://download-installer.cdn.mozilla.net/pub/firefox/releases/
ARG FIREFOX_VERSION="91.1.0esr"
# Gecko driver releases
# https://github.com/mozilla/geckodriver/releases
ARG GECKODRIVER_VERSION="v0.30.0"
# Chrome versions
# https://www.ubuntuupdates.org/package/google_chrome/stable/main/base/google-chrome-stable
ARG CHROME_VERSION="94.0.4606.54"

ARG SELENIUM_MAJOR_VERSION=3

ARG SELENIUM_MINOR_VERSION=141

ARG SELENIUM_PATCH_VERSION=59

ENV SELENIUM_PORT=4444 \
    VNC_PORT=5999 \
    DISPLAY=:99 \
    DBUS_SESSION_BUS_ADDRESS=/dev/null \
    HOME=${SELENIUM_HOME} \
    VNC_GEOMETRY=${VNC_GEOMETRY:-"1600x900"} \
    SELENIUM_VERSION=${SELENIUM_MAJOR_VERSION}.${SELENIUM_MINOR_VERSION}.${SELENIUM_PATCH_VERSION} \
    SELENIUM_PATH=${SELENIUM_HOME}/selenium-server/selenium-server-standalone.jar \
    PATH=${SELENIUM_HOME}/firefox:/opt/google/chrome:${PATH}

EXPOSE ${SELENIUM_PORT}

EXPOSE ${VNC_PORT}

RUN mkdir -p .cache/dconf .mozilla/plugins .vnc/ .fluxbox/ && \
    echo "session.screen0.toolbar.visible: false" > .fluxbox/init && \
    touch .Xauthority .vnc/config

RUN mkdir -p ${SELENIUM_HOME}/selenium-server && \
    curl -L http://selenium-release.storage.googleapis.com/${SELENIUM_MAJOR_VERSION}.${SELENIUM_MINOR_VERSION}/selenium-server-standalone-${SELENIUM_VERSION}.jar \
         -o ${SELENIUM_PATH}

RUN curl -L https://dl.google.com/linux/chrome/rpm/stable/x86_64/google-chrome-stable-${CHROME_VERSION}-1.x86_64.rpm \
         -o google-chrome-stable-x86_64.rpm && \
    rpm -i google-chrome-stable-x86_64.rpm && \
    rm -f google-chrome-stable-x86_64.rpm

# chrome and chrome driver versions should match in order to avoid incompatibility
RUN CHROME_VERSION=$(rpm -q --qf "%{VERSION}\n" google-chrome-stable | sed -Ee 's/^(.*)\..*/\1/') && \
    CHROME_DRIVER_VERSION=$(curl -s https://chromedriver.storage.googleapis.com/LATEST_RELEASE_${CHROME_VERSION}) && \
    curl -O https://chromedriver.storage.googleapis.com/${CHROME_DRIVER_VERSION}/chromedriver_linux64.zip && \
    unzip -d /usr/bin/ chromedriver_linux64.zip && \
    chmod +x /usr/bin/chromedriver && \
    rm -f chromedriver_linux64.zip

RUN curl -LO https://download-installer.cdn.mozilla.net/pub/firefox/releases/${FIREFOX_VERSION}/linux-x86_64/en-US/firefox-${FIREFOX_VERSION}.tar.bz2 && \
    tar -C . -xjvf firefox-${FIREFOX_VERSION}.tar.bz2 && \
    rm -f firefox-${FIREFOX_VERSION}.tar.bz2

RUN curl -LO https://github.com/mozilla/geckodriver/releases/download/${GECKODRIVER_VERSION}/geckodriver-${GECKODRIVER_VERSION}-linux64.tar.gz && \
    tar -C /usr/bin/ -xvf geckodriver-${GECKODRIVER_VERSION}-linux64.tar.gz && \
    rm -f geckodriver-${GECKODRIVER_VERSION}-linux64.tar.gz

COPY --from=builder /go/init /usr/bin/

RUN chmod +x /usr/bin/init && \
    chgrp -R 0 ${SELENIUM_HOME} && \
    chmod -R g=u ${SELENIUM_HOME}

USER 1001

# we use a custom init to start and stop container processes in the specific order
CMD ["init"]
