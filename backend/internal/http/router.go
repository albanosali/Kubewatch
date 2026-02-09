package http

import (
    "net/http"

    "github.com/albanosali/Kubewatch/backend/internal/k8s"
    promclient "github.com/albanosali/Kubewatch/backend/internal/prometheus"
    "github.com/gorilla/mux"
)

type Server struct {
    Prom *promclient.Client
    K8s  *k8s.Client
}

func NewRouter(prom *promclient.Client, k8s *k8s.Client) http.Handler {
    s := &Server{Prom: prom, K8s: k8s}
    
    r := mux.NewRouter()
    r.HandleFunc("/api/namespaces", s.handleNamespaces).Methods("GET")
    r.HandleFunc("/api/issues", s.handleIssues).Methods("GET")
    r.HandleFunc("/api/workloads", s.handleWorkloads).Methods("GET")
    
    // Health check
    r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }).Methods("GET")
    
    return r
}
