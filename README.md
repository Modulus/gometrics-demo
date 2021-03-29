# WHat is this
1. Demo app of golang running custom prometheus metrics
2. Locust for triggering metrics

# Run locally
docker-comopse up 

# Build your own

docker build -t yourrepo/metrics-demo .

docker build -t yourrepo/metrics-locust ./locust


# Kubernetes kind

## Nginx ingress controller
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/kind/deploy.yaml
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s

## Ambassador
kubectl apply -f https://github.com/datawire/ambassador-operator/releases/latest/download/ambassador-operator-crds.yaml
kubectl apply -n ambassador -f https://github.com/datawire/ambassador-operator/releases/latest/download/ambassador-operator-kind.yaml
kubectl wait --timeout=180s -n ambassador --for=condition=deployed ambassadorinstallations/ambassador

kubectl annotate ingress example-ingress kubernetes.io/ingress.class=ambassador


## Kube-prometheus
kubectl apply -k ./prometheus

### Metrics
http_request_duration_milliseconds{namespace="monitoring", service="grafana"}