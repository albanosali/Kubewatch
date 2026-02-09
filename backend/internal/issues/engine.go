package issues

import (
    "context"

    "github.com/albanosali/Kubewatch/backend/internal/prometheus"
)

type MetricsProvider interface {
    GetWorkloadMetrics(ctx context.Context) ([]prometheus.WorkloadMetrics, error)
}

func BuildIssues(ctx context.Context, mp MetricsProvider) ([]Issue, error) {
    metrics, err := mp.GetWorkloadMetrics(ctx)
    if err != nil {
        return nil, err
    }

    issues := make([]Issue, 0)

    for _, m := range metrics {
        // OOM incidents
        if m.OOMKills > 0 {
            issues = append(issues, Issue{
                Namespace:  m.Namespace,
                Workload:   m.Workload,
                Pod:        m.Pod,
                Type:       IssueTypeOOM,
                Severity:   SeverityCritical,
                Message:    fmt.Sprintf("Pod had %d OOM kills in last hour", int64(m.OOMKills)),
                Suggestion: "Increase memory limits or optimize memory usage. Check application logs for memory leaks.",
            })
        }

        // High memory usage
        if m.MemRequest > 0 && m.MemUsagePct > 90 {
            issues = append(issues, Issue{
                Namespace:  m.Namespace,
                Workload:   m.Workload,
                Pod:        m.Pod,
                Type:       IssueTypeHighMemory,
                Severity:   SeverityHigh,
                Message:    fmt.Sprintf("Memory usage at %.1f%% of request (%.1f/%.1f MiB)", 
                    m.MemUsagePct, m.MemUsage/1024/1024, m.MemRequest/1024/1024),
                Suggestion: fmt.Sprintf("Increase memory request to %.1f MiB or optimize application", 
                    m.MemRequest*1.2/1024/1024),
            })
        }

        // High CPU usage
        if m.CPURequest > 0 && m.CPUUsagePct > 90 {
            issues = append(issues, Issue{
                Namespace:  m.Namespace,
                Workload:   m.Workload,
                Pod:        m.Pod,
                Type:       IssueTypeHighCPU,
                Severity:   SeverityMedium,
                Message:    fmt.Sprintf("CPU usage at %.1f%% of request (%.3f/%.3f cores)", 
                    m.CPUUsagePct, m.CPUUsage, m.CPURequest),
                Suggestion: fmt.Sprintf("Increase CPU request to %.3f cores or optimize CPU usage", 
                    m.CPURequest*1.2),
            })
        }
    }

    return issues, nil
}
