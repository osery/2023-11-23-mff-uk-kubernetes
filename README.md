# Kubernetes Talk Demo

Source codes for demos to the Kubernetes talk at MFF-UK on the 23rd of November 2023.

## Containers

Build single stage:

```bash
docker build . -f Dockerfile.multi -t osery/webserver:single
```

Build multi stage:

```bash
docker build . -f Dockerfile.multi -t osery/webserver:multi
```

Check container sizes:

```bash
docker image ps | grep osery/webserver
```

Run container:

```bash
docker run --rm -p 8080:8080 osery/webserver:multi
```

Build for multple platforms using docker buildx plugin.

```bash
docker buildx build . -f Dockerfile.multi --platform=linux/amd64,linux/arm64 -t osery/webserver
```

## Deployments

```bash
kubectl rollout undo deployment/webserver -n mff
```