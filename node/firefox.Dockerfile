FROM quay.io/redhatqe/selenium-node-base:4.0.0

LABEL maintainer="dmisharo@redhat.com"

# Firefox releases
# https://download-installer.cdn.mozilla.net/pub/firefox/releases/
ARG FIREFOX_VERSION="91.3.0esr"
# Gecko driver releases
# https://github.com/mozilla/geckodriver/releases
ARG GECKODRIVER_VERSION="v0.30.0"

RUN curl -LO https://download-installer.cdn.mozilla.net/pub/firefox/releases/${FIREFOX_VERSION}/linux-x86_64/en-US/firefox-${FIREFOX_VERSION}.tar.bz2 && \
    tar -C /opt/ -xjvf firefox-${FIREFOX_VERSION}.tar.bz2 && \
    rm -f firefox-${FIREFOX_VERSION}.tar.bz2

RUN curl -LO https://github.com/mozilla/geckodriver/releases/download/${GECKODRIVER_VERSION}/geckodriver-${GECKODRIVER_VERSION}-linux64.tar.gz && \
    tar -C /opt/bin/ -xvf geckodriver-${GECKODRIVER_VERSION}-linux64.tar.gz && \
    rm -f geckodriver-${GECKODRIVER_VERSION}-linux64.tar.gz

RUN chgrp -R 0 ${HOME} && \
    chmod -R g=u ${HOME}

USER 1001

# we use a custom init to start and stop container processes in the specific order
CMD ["init"]
