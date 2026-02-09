package main

import (
    "log"
    "net/http"
    "os"
    "time"

    "github.com/albanosali/Kubewatch/backend/internal/http"
    k8sclient "github.com/albanosali/Kubewatch/backend/internal/k8s"
    promclient "github.com/albanosali/Kubewatch/backend/internal/prometheus"
    "github.com/gorilla/handlers"
)

func main() {
    promURL := os.Getenv("PROMETHEUS_URL")
    if promURL == "" {
        promURL = "http://prometheus-server.monitoring.svc.cluster.local:9090"
    }

    prom, err := promclient.NewClient(promURL, 10*time.Second)
    if err != nil {
        log.Fatalf("failed to create prometheus client: %v", err)
    }

    k8s, err := k8sclient.NewInClusterClient()
    if err != nil {
        log.Fatalf("failed to create k8s client: %v", err)
    }

    router := http.NewRouter(prom, k8s)

    // Enable CORS for frontend
    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"*"}),
    )(router)

    addr := ":8080"
    log.Printf("Kubewatch starting on %s (Prometheus: %s)", addr, promURL)
    log.Fatal(http.ListenAndServe(addr, corsHandler))
}
