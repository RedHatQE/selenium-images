FROM quay.io/redhatqe/selenium-base:4.0.0

LABEL maintainer="dmisharo@redhat.com"

ENV SE_SESSION_REQUEST_TIMEOUT=300 \
    SE_SESSION_RETRY_INTERVAL=5

EXPOSE 4444

COPY start.sh /opt/selenium/

USER 1001

CMD ["/opt/selenium/start.sh"]
