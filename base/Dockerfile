FROM registry.fedoraproject.org/fedora:30
LABEL maintainer="dmisharo@redhat.com"

ENV SELENIUM_HOME=/home/selenium

WORKDIR $SELENIUM_HOME

RUN dnf -y update && \
    dnf install -y --setopt=tsflags=nodocs \
        bzip2 \
        dbus-glib \
        dejavu-sans-fonts \
        dejavu-serif-fonts \
        dumb-init \
        fluxbox \
        java-1.8.0-openjdk \
        libappindicator-gtk3 \
        libcurl \
        liberation-fonts \
        libXScrnSaver \
        net-tools \
        redhat-lsb \
        tigervnc-server \
        unzip \
        wget \
        xdg-utils \
        xdotool && \
    dnf clean all && \
    touch $SELENIUM_HOME/.Xauthority && \
    mkdir -p $SELENIUM_HOME/.cache/dconf && \
    mkdir -p $SELENIUM_HOME/.mozilla/plugins
