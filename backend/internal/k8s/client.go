package k8s

import (
    "context"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
)

type Client struct {
    cs *kubernetes.Clientset
}

func NewInClusterClient() (*Client, error) {
    cfg, err := rest.InClusterConfig()
    if err != nil {
        return nil, err
    }
    cs, err := kubernetes.NewForConfig(cfg)
    if err != nil {
        return nil, err
    }
    return &Client{cs: cs}, nil
}

func (c *Client) ListNamespaces(ctx context.Context) ([]string, error) {
    nsList, err := c.cs.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
    if err != nil {
        return nil, err
    }
    res := make([]string, 0, len(nsList.Items))
    for _, ns := range nsList.Items {
        res = append(res, ns.Name)
    }
    return res, nil
}
