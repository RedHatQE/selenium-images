# Red Hat QE Selenium Images

Fedora based container images for running UI tests. Images can be used locally and in OpenShift.

## Selenium Grid

Images in the `grid` directory are supposed to be use in the Selenium Grid. Please refer to the
[Selenium documentation](https://www.selenium.dev/documentation/grid/) for the details.

## Selenium Standalone

Standalone version includes Google Chrome and Mozilla Firefox browsers. A container is supposed to
be runnning either locally or within a Kubernetes pod.

In order to start a container locally use the following command:

`podman run -it --shm-size=2g -p 4444:4444 -p 5999:5999 quay.io/redhatqe/selenium-standalone`

* `4444`: standard selenium standalone server port
* `5999`: VNC port

To run a container within a pod use this manifest as an example:

```yaml
kind: Pod
apiVersion: v1
metadata:
  name: some-name
spec:
  containers:
    - resources:
        limits:
          cpu: '1'
          memory: 3Gi
        requests:
          cpu: '1'
          memory: 3Gi
      terminationMessagePath: /dev/termination-log
      name: selenium
      imagePullPolicy: Always
      volumeMounts:
        - name: shm
          mountPath: /dev/shm
      terminationMessagePolicy: File
      image: 'quay.io/redhatqe/selenium-standalone:latest'
  volumes:
    - name: shm
      emptyDir:
        medium: Memory
        sizeLimit: 2Gi
```

The standalone container image also starts a small HTTP server to allow selenium to be shut down via an HTTP GET to '/shutdown' on port 8000. This is useful
when running the selenium container as a sidecar in a kubernetes pod to have your testing container shut the selenium container down when tests are completed.
