package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type dbPinger interface {
	Ping(context.Context) error
}

type storageHealth interface {
	Check(ctx context.Context) error
}

// Health exposes liveness/readiness endpoints for the treasury service.
type Health struct {
	log     *zap.Logger
	db      dbPinger
	cache   *redis.Client
	events  *nats.Conn
	storage storageHealth
}

func NewHealth(log *zap.Logger, db dbPinger, cache *redis.Client, events *nats.Conn, storage storageHealth) *Health {
	return &Health{log: log, db: db, cache: cache, events: events, storage: storage}
}

func (h *Health) Liveness(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"service": "treasury-app",
	})
}

func (h *Health) Readiness(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	issues := map[string]string{}

	if h.db != nil {
		if err := h.db.Ping(ctx); err != nil {
			issues["postgres"] = err.Error()
		}
	}

	if h.cache != nil {
		if err := h.cache.Ping(ctx).Err(); err != nil {
			issues["redis"] = err.Error()
		}
	}

	if h.events != nil && !h.events.IsConnected() {
		issues["nats"] = "not connected"
	}

	if h.storage != nil {
		if err := h.storage.Check(ctx); err != nil {
			issues["storage"] = err.Error()
		}
	}

	status := http.StatusOK
	if len(issues) > 0 {
		status = http.StatusServiceUnavailable
	}

	respondJSON(w, status, map[string]any{
		"status":       http.StatusText(status),
		"dependencies": issues,
	})
}

func (h *Health) Metrics(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}
