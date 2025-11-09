package storage

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bengobox/treasury-app/internal/config"
)

type HealthChecker struct {
	client   *http.Client
	endpoint string
}

func NewHealthChecker(cfg config.StorageConfig) *HealthChecker {
	return &HealthChecker{
		client:   &http.Client{Timeout: 5 * time.Second},
		endpoint: cfg.Endpoint,
	}
}

func (h *HealthChecker) Check(ctx context.Context) error {
	if h == nil || h.endpoint == "" {
		return nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodHead, h.endpoint, nil)
	if err != nil {
		return fmt.Errorf("storage health request: %w", err)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return fmt.Errorf("storage health request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("storage endpoint unhealthy: status %d", resp.StatusCode)
	}

	return nil
}
