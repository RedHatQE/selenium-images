FROM docker.io/library/golang:1.17 AS builder

COPY init.go .

RUN go build init.go && chmod +x /go/init

FROM registry.fedoraproject.org/fedora-minimal:34
LABEL maintainer="dmisharo@redhat.com"

# Firefox releases
# https://download-installer.cdn.mozilla.net/pub/firefox/releases/
ARG FIREFOX_VERSION="91.9.0esr"
# Gecko driver releases
# https://github.com/mozilla/geckodriver/releases
ARG GECKODRIVER_VERSION="v0.31.0"
# Chrome versions
# https://www.ubuntuupdates.org/package/google_chrome/stable/main/base/google-chrome-stable
ARG CHROME_VERSION="101.0.4951.54"

ARG SELENIUM_MAJOR_VERSION=4

ARG SELENIUM_MINOR_VERSION=1

ARG SELENIUM_PATCH_VERSION=4

ENV SELENIUM_HOME=/home/selenium

ENV SELENIUM_PORT=4444 \
    SELENIUM_SESSION_TIMEOUT=1800 \
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

WORKDIR ${SELENIUM_HOME}

RUN PACKAGES="\
        alsa-lib \
        at-spi2-atk \
        at-spi2-core \
        atk \
        avahi-libs \
        bzip2 \
        cairo \
        cairo-gobject \
        cups-libs \
        dbus-glib \
        dbus-libs \
        expat \
        fluxbox \
        fontconfig \
        freetype \
        fribidi \
        gdk-pixbuf2 \
        graphite2 \
        gtk3 \
        harfbuzz \
        imlib2 \
        java-1.8.0-openjdk-headless \
        libcloudproviders \
        libdatrie \
        libdrm \
        libepoxy \
        liberation-fonts \
        liberation-fonts-common \
        liberation-mono-fonts \
        liberation-sans-fonts \
        liberation-serif-fonts \
        libfontenc \
        libglvnd \
        libglvnd-glx \
        libICE \
        libjpeg-turbo \
        libpng \
        libSM \
        libthai \
        libwayland-client \
        libwayland-cursor \
        libwayland-egl \
        libwayland-server \
        libwebp \
        libX11 \
        libX11-common \
        libX11-xcb \
        libXau \
        libxcb \
        libXcomposite \
        libXcursor \
        libXdamage \
        libXdmcp \
        libXext \
        libXfixes \
        libXfont2 \
        libXft \
        libXi \
        libXinerama \
        libxkbcommon \
        libxkbfile \
        libXpm \
        libXrandr \
        libXrender \
        libxshmfence \
        libXt \
        mesa-libgbm \
        nspr \
        nss \
        nss-softokn \
        nss-softokn-freebl \
        nss-util \
        pango \
        pixman \
        tar \
        tigervnc-server-minimal \
        tzdata-java \
        unzip \
        vulkan-loader \
        wget \
        xdg-utils \
        xkbcomp \
        xkeyboard-config" && \
    microdnf download -y --archlist=x86_64,noarch ${PACKAGES} && \
    rpm -Uvh --nodeps *.rpm && \
    rm -f *.rpm && \
    microdnf clean all

RUN mkdir -p .cache/dconf .mozilla/plugins .vnc/ .fluxbox/ && \
    echo "session.screen0.toolbar.visible: false" > .fluxbox/init && \
    touch .Xauthority .vnc/config

RUN mkdir -p ${SELENIUM_HOME}/selenium-server && \
    curl -L https://github.com/SeleniumHQ/selenium/releases/download/selenium-${SELENIUM_MAJOR_VERSION}.${SELENIUM_MINOR_VERSION}.0/selenium-server-${SELENIUM_VERSION}.jar \
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

COPY --from=builder /go/init /init

RUN chown -R 1001:0 ${SELENIUM_HOME} && \
    chmod -R g=u ${SELENIUM_HOME}

USER 1001
# we use a custom init to start and stop container processes in the specific order
CMD ["/init"]
