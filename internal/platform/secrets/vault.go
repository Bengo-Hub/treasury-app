package secrets

import "context"

// Provider defines the interface required for secret retrieval.
type Provider interface {
	Get(ctx context.Context, path string) (map[string]any, error)
}

type noopProvider struct{}

func (noopProvider) Get(context.Context, string) (map[string]any, error) {
	return map[string]any{}, nil
}

// NewNoop creates a placeholder secrets provider until Vault integration is wired.
func NewNoop() Provider {
	return noopProvider{}
}
