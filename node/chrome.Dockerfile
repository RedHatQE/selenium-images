FROM quay.io/redhatqe/selenium-node-base:4.0.0

# Chrome versions
# https://www.ubuntuupdates.org/package/google_chrome/stable/main/base/google-chrome-stable
ARG CHROME_VERSION="95.0.4638.69"

RUN curl -L https://dl.google.com/linux/chrome/rpm/stable/x86_64/google-chrome-stable-${CHROME_VERSION}-1.x86_64.rpm \
         -o google-chrome-stable-x86_64.rpm && \
    rpm -i google-chrome-stable-x86_64.rpm && \
    rm -f google-chrome-stable-x86_64.rpm

# chrome and chrome driver versions should match in order to avoid incompatibility
RUN CHROME_VERSION=$(rpm -q --qf "%{VERSION}\n" google-chrome-stable | sed -Ee 's/^(.*)\..*/\1/') && \
    CHROME_DRIVER_VERSION=$(curl -s https://chromedriver.storage.googleapis.com/LATEST_RELEASE_${CHROME_VERSION}) && \
    curl -O https://chromedriver.storage.googleapis.com/${CHROME_DRIVER_VERSION}/chromedriver_linux64.zip && \
    unzip -d /opt/bin/ chromedriver_linux64.zip && \
    chmod +x /opt/bin/chromedriver && \
    rm -f chromedriver_linux64.zip

RUN chgrp -R 0 ${HOME} && \
    chmod -R g=u ${HOME}

USER 1001

# we use a custom init to start and stop container processes in the specific order
CMD ["init"]

