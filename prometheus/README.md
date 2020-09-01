# What is this?
Files from the manifests folder in https://github.com/prometheus-operator/kube-prometheus

## Why?
To create our own kustomize setup of the kube-prometheus config, with kustomize patches where needed.

## More
Adding custom dashboards and alerts later down the line

## Installing
```
kubectl kustomize
kubectl apply -k .
```


# TODO: Add loki config?