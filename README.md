# WHat is this
1. Demo app of golang running custom prometheus metrics
2. Locust for triggering metrics

# Run locally
docker-comopse up 

# Build your own

docker build -t yourrepo/metrics-demo .

docker build -t yourrepo/metrics-locust ./locust
