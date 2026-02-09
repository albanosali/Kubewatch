package issues

type IssueType string
type Severity string

const (
    IssueTypeOOM          IssueType = "OOM"
    IssueTypeResource     IssueType = "ResourcePressure"
    IssueTypeHighCPU      IssueType = "HighCPU"
    IssueTypeHighMemory   IssueType = "HighMemory"
    
    SeverityLow           Severity  = "LOW"
    SeverityMedium        Severity  = "MEDIUM"
    SeverityHigh          Severity  = "HIGH"
    SeverityCritical      Severity  = "CRITICAL"
)

type Issue struct {
    Namespace  string    `json:"namespace"`
    Workload   string    `json:"workload"`
    Pod        string    `json:"pod"`
    Type       IssueType `json:"type"`
    Severity   Severity  `json:"severity"`
    Message    string    `json:"message"`
    Suggestion string    `json:"suggestion"`
}
