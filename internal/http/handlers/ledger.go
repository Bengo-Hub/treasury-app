package handlers

import (
	"net/http"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

// Ledger exposes read-only endpoints for chart of accounts and balances.
type Ledger struct {
	log *zap.Logger
}

func NewLedger(log *zap.Logger) *Ledger {
	return &Ledger{log: log}
}

func (h *Ledger) ChartOfAccounts(w http.ResponseWriter, r *http.Request) {
	payload := map[string]any{
		"accounts": []map[string]any{
			{
				"code":          "1000",
				"name":          "Platform Cash",
				"type":          "asset",
				"currency":      "KES",
				"balance":       decimal.NewFromInt(0),
				"tenant":        r.Header.Get("X-Tenant-ID"),
				"allowPostings": true,
			},
		},
	}

	respondJSON(w, http.StatusOK, payload)
}
