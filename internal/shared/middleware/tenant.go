package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type tenantKey struct{}

// Tenant extracts tenant metadata from headers or URL params and places it in the request context.
func Tenant(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenant := strings.TrimSpace(r.Header.Get("X-Tenant-ID"))
		if tenant == "" {
			if routeCtx := chi.RouteContext(r.Context()); routeCtx != nil {
				if param := strings.TrimSpace(routeCtx.URLParam("tenantID")); param != "" {
					tenant = param
				}
			}
		}

		if tenant != "" {
			r = r.WithContext(context.WithValue(r.Context(), tenantKey{}, tenant))
		}

		next.ServeHTTP(w, r)
	})
}

// TenantFromContext returns the tenant identifier.
func TenantFromContext(ctx context.Context) string {
	if v, ok := ctx.Value(tenantKey{}).(string); ok {
		return v
	}
	return ""
}
