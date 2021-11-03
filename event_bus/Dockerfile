FROM quay.io/redhatqe/selenium-base:4.0.0

LABEL maintainer="dmisharo@redhat.com"

EXPOSE 5553

COPY start.sh /opt/selenium/

USER 1001

CMD ["/opt/selenium/start.sh"]
