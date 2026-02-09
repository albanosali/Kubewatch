package prometheus

import (
    "context"
    "fmt"
    "strconv"
    "time"

    v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type WorkloadMetrics struct {
    Namespace   string  `json:"namespace"`
    Pod         string  `json:"pod"`
    Workload    string  `json:"workload"`
    CPUUsage    float64 `json:"cpuUsageCores"`
    CPURequest  float64 `json:"cpuRequestCores"`
    MemUsage    float64 `json:"memUsageBytes"`
    MemRequest  float64 `json:"memRequestBytes"`
    OOMKills    float64 `json:"oomKills"`
    CPUUsagePct float64 `json:"cpuUsagePct"`
    MemUsagePct float64 `json:"memUsagePct"`
}
type LatencyMetrics struct {
    Namespace string  `json:"namespace"`
    Service   string  `json:"service"`
    P50       float64 `json:"p50_ms"`
    P95       float64 `json:"p95_ms"`
    P99       float64 `json:"p99_ms"`
    ErrorRate float64 `json:"error_rate_pct"`
}

func (c *Client) GetLatencyMetrics(ctx context.Context) ([]LatencyMetrics, error) {
    // Mock data for now - replace with real PromQL for http_request_duration_seconds
    return []LatencyMetrics{
        {Namespace: "default", Service: "api", P50: 23, P95: 150, P99: 450, ErrorRate: 0.5},
        {Namespace: "prod", Service: "payments", P50: 45, P95: 320, P99: 1200, ErrorRate: 2.1},
    }, nil
}
func (c *Client) GetWorkloadMetrics(ctx context.Context) ([]WorkloadMetrics, error) {
    ts := time.Now()

    // Get OOM kills
    oomQuery := `sum(increase(container_oom_events_total{container!=""}[1h])) by (namespace, pod)`
    oomVal, err := c.Query(ctx, oomQuery, ts)
    if err != nil {
        return nil, fmt.Errorf("OOM query failed: %w", err)
    }

    // Get CPU usage vs requests
    cpuUsageQuery := `sum(rate(container_cpu_usage_seconds_total{container!=""}[5m])) by (namespace, pod)`
    cpuReqQuery := `sum(kube_pod_container_resource_requests{resource="cpu"}) by (namespace, pod)`
    
    cpuUsageVal, _ := c.Query(ctx, cpuUsageQuery, ts)
    cpuReqVal, _ := c.Query(ctx, cpuReqQuery, ts)

    // Get Memory usage vs requests
    memUsageQuery := `sum(container_memory_working_set_bytes{container!=""} ) by (namespace, pod)`
    memReqQuery := `sum(kube_pod_container_resource_requests{resource="memory"}) by (namespace, pod)`
    
    memUsageVal, _ := c.Query(ctx, memUsageQuery, ts)
    memReqVal, _ := c.Query(ctx, memReqQuery, ts)

    // For v0.1, return mock data to test UI
    // TODO: Parse actual Prometheus vector results
    mockMetrics := []WorkloadMetrics{
        {
            Namespace:   "default",
            Pod:         "api-abc123",
            Workload:    "api",
            CPUUsage:    0.25,
            CPURequest:  0.5,
            MemUsage:    256 * 1024 * 1024,
            MemRequest:  512 * 1024 * 1024,
            OOMKills:    2,
            CPUUsagePct: 50,
            MemUsagePct: 50,
        },
        {
            Namespace:   "prod",
            Pod:         "payment-def456",
            Workload:    "payments",
            CPUUsage:    0.8,
            CPURequest:  1.0,
            MemUsage:    900 * 1024 * 1024,
            MemRequest:  1024 * 1024 * 1024,
            OOMKills:    0,
            CPUUsagePct: 80,
            MemUsagePct: 87.89,
        },
    }

    return mockMetrics, nil
}
