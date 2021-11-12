FROM docker.io/library/golang:1.17 AS builder

COPY init.go .

RUN go build init.go

FROM quay.io/redhatqe/selenium-base:4.0.0
LABEL maintainer="dmisharo@redhat.com"

ENV VNC_PORT=5999 \
    DISPLAY=:99 \
    DBUS_SESSION_BUS_ADDRESS=/dev/null \
    HOME=/home/selenium \
    VNC_GEOMETRY=${VNC_GEOMETRY:-"1600x900"} \
    PATH=/opt/bin/:/opt/firefox/:/opt/google/chrome/:${PATH}

EXPOSE ${VNC_PORT}

WORKDIR ${HOME}

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
        pango \
        pixman \
        tar \
        tigervnc-server-minimal \
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

COPY --from=builder /go/init /opt/bin/

RUN chmod +x /opt/bin/init
