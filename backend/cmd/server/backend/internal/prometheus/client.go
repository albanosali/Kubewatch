package prometheus

import (
    "context"
    "fmt"
    "net/http"
    "time"

    promapi "github.com/prometheus/client_golang/api"
    v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type Client struct {
    api v1.API
}

func NewClient(url string, timeout time.Duration) (*Client, error) {
    cfg := promapi.Config{
        Address:      url,
        RoundTripper: &http.Transport{Proxy: http.ProxyFromEnvironment},
    }
    c, err := promapi.NewClient(cfg)
    if err != nil {
        return nil, fmt.Errorf("create prometheus client: %w", err)
    }
    api := v1.NewAPI(c)
    return &Client{api: api}, nil
}

func (c *Client) Query(ctx context.Context, promQL string, ts time.Time) (v1.Value, error) {
    res, warnings, err := c.api.Query(ctx, promQL, ts)
    if len(warnings) > 0 {
        fmt.Printf("Prometheus warnings: %v
", warnings)
    }
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (c *Client) QueryRange(ctx context.Context, promQL string, start, end time.Time, step time.Duration) (v1.Range, error) {
    res, warnings, err := c.api.QueryRange(ctx, promQL, v1.Range{Start: start, End: end, Step: step})
    if len(warnings) > 0 {
        fmt.Printf("Prometheus warnings: %v
", warnings)
    }
    if err != nil {
        return v1.Range{}, err
    }
    return res, nil
}
