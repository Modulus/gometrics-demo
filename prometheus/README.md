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

# Run on minikube with release-0.4

release-0.4 is compatible with kubernetes 16.5x
Start minikube

minikube start --memory=6000 --cpus=6 --kubernetes-version=v1.16.9 --bootstrapper=kubeadm --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook --extra-config=scheduler.address=0.0.0.0 --extra-config=controller-manager.address=0.0.0.0


# TODO: Add loki config?