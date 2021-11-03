FROM registry.fedoraproject.org/fedora-minimal:35
LABEL maintainer="dmisharo@redhat.com"

ARG SELENIUM_MAJOR_VERSION=4

ARG SELENIUM_MINOR_VERSION=0

ARG SELENIUM_PATCH_VERSION=0

ENV SELENIUM_PATH=/opt/selenium/selenium-server.jar

RUN PACKAGES="\
        nspr \
        nss \
        nss-softokn \
        nss-softokn-freebl \
        nss-util \
        java-1.8.0-openjdk-headless \
        tzdata-java" && \
    microdnf download -y --archlist=x86_64,noarch ${PACKAGES} && \
    rpm -Uvh --nodeps *.rpm && \
    rm -f *.rpm && \
    microdnf clean all

RUN mkdir -p /opt/selenium/ && \
    SELENIUM_VERSION=${SELENIUM_MAJOR_VERSION}.${SELENIUM_MINOR_VERSION}.${SELENIUM_PATCH_VERSION} && \
    curl -L https://github.com/SeleniumHQ/selenium/releases/download/selenium-${SELENIUM_VERSION}/selenium-server-${SELENIUM_VERSION}.jar \
         -o ${SELENIUM_PATH}
