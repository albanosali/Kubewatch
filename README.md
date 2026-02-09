# Kubewatch
# Kubewatch - Kubernetes Health Visualizer

Open-source web UI that shows OOM kills, resource pressure, and workload metrics from Prometheus + Kubernetes API.

[![Screenshot](screenshots/example.png)](screenshots/example.png)

## 🚀 Quick Start

```bash
# Deploy with Helm
helm repo add kubewatch https://albanosali.github.io/helm-charts
helm install kubewatch kubewatch/kubewatch --set prometheusUrl=http://prometheus:9090

# Port-forward to view
kubectl port-forward svc/kubewatch 8080:80
# Open http://localhost:8080
