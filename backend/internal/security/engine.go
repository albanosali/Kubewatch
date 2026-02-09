package security

type SecurityIssue struct {
    Namespace string `json:"namespace"`
    Type      string `json:"type"`
    Severity  string `json:"severity"`
    Message   string `json:"message"`
}

func GetSecurityPosture() []SecurityIssue {
    // Mock CIS benchmark results
    return []SecurityIssue{
        {Namespace: "default", Type: "RBAC", Severity: "HIGH", Message: "ServiceAccount runs as root"},
        {Namespace: "prod", Type: "Network", Severity: "MEDIUM", Message: "No NetworkPolicy found"},
    }
}
