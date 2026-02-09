package http

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/albanosali/Kubewatch/backend/internal/issues"
    "github.com/gorilla/mux"
)

func (s *Server) handleNamespaces(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()
    
    ns, err := s.K8s.ListNamespaces(ctx)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    writeJSON(w, http.StatusOK, ns)
}

func (s *Server) handleIssues(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
    defer cancel()
    
    iss, err := issues.BuildIssues(ctx, s.Prom)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    writeJSON(w, http.StatusOK, iss)
}

func (s *Server) handleWorkloads(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
    defer cancel()
    
    metrics, err := s.Prom.GetWorkloadMetrics(ctx)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    writeJSON(w, http.StatusOK, metrics)
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(v)
}
