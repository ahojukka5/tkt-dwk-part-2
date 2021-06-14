# DevOps with Kubernetes

Build app and run using Kubernetes:

```bash
docker build -t ahojukka5/dwk-app:0.1 app
```

```bash
docker push ahojukka5/dwk-app:0.1
```

```bash
kubectl create deployment dwk-app --image=ahojukka5/dwk-app:0.1
```

```bash
kubectl get deployments
```

```bash
kubectl get pods
```

```bash
kubectl logs -f dwk-app-75b7576c45-tqjqf
```
