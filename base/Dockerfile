FROM registry.fedoraproject.org/fedora:34
LABEL maintainer="dmisharo@redhat.com"

ENV SELENIUM_HOME=/home/selenium

WORKDIR $SELENIUM_HOME

RUN dnf install -y --nodocs \
        bzip2 \
        dbus-glib \
        dejavu-sans-fonts \
        dejavu-serif-fonts \
        dumb-init \
        fluxbox \
        java-1.8.0-openjdk \
        libappindicator-gtk3 \
        liberation-fonts \
        libXScrnSaver \
        net-tools \
        redhat-lsb \
        tigervnc-server \
        unzip \
        wget \
        xdg-utils \
        xdotool \
        xmodmap \
        xrdb \
        mesa-libgbm \
        vulkan-loader && \
    dnf clean all && \
    mkdir -p .cache/dconf .mozilla/plugins .vnc/ && \
    touch .Xauthority .vnc/config .Xmodmap
