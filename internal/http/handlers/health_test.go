package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	testing "testing"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type fakeDB struct{ err error }

func (f fakeDB) Ping(context.Context) error { return f.err }

type fakeStorage struct{ err error }

func (f fakeStorage) Check(context.Context) error { return f.err }

func TestHealth_Liveness(t *testing.T) {
	h := NewHealth(zap.NewNop(), nil, nil, nil, nil)
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	h.Liveness(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestHealth_Readiness_WithIssues(t *testing.T) {
	h := NewHealth(zap.NewNop(), fakeDB{err: context.DeadlineExceeded}, redis.NewClient(&redis.Options{Addr: "localhost:0"}), nil, fakeStorage{err: context.DeadlineExceeded})
	req := httptest.NewRequest(http.MethodGet, "/readyz", nil)
	rec := httptest.NewRecorder()

	h.Readiness(rec, req)

	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected 503, got %d", rec.Code)
	}
}
