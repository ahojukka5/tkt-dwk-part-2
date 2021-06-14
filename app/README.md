# Workflow to update app

Build new image:

```bash
docker build -t ahojukka5/dwk-app:<tag>
```

Push it to docker hub:

```bash
docker push ahojukka5/dwk-app:<tag>
```

Change tag in `manifests/deployment.yaml` and apply changes:

```bash
kubectl apply -f manifests/deployment.yaml
```
