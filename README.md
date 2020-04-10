![Building and pushing to quay.io](https://github.com/RedHatQE/selenium-images/workflows/Building%20Red%20Hat%20QE%20selenium%20images/badge.svg)


Red Hat QE Selenium Images
==========================

Docker images for running UI tests. Images can be used locally and in openshift.


Deployment
==========

Images are built using [Github Actions](https://github.com/RedHatQE/selenium-images/actions) and pushed to quay.io registry: https://quay.io/repository/redhatqe/selenium-standalone


Usage
=====

In order to start a container use the following command:

`podman run -it --shm-size=2g -p 4444:4444 -p 5999:5999 quay.io/redhatqe/selenium-standalone`


Exposed ports
=============

* `4444`: standard selenium standalone server port
* `5999`: VNC port
