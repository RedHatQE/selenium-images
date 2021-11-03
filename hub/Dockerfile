FROM quay.io/redhatqe/selenium-base:4.0.0

LABEL maintainer="dmisharo@redhat.com"

EXPOSE 4442
EXPOSE 4443
EXPOSE 4444

ENV SE_SESSION_REQUEST_TIMEOUT=300 \
    SE_SESSION_RETRY_INTERVAL=5 \
    SE_RELAX_CHECKS=true

COPY start.sh /opt/selenium/

USER 1001

CMD ["/opt/selenium/start.sh"]
