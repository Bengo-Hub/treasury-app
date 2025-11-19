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

type ledgerAccount struct {
	Code          string `json:"code" example:"1000"`
	Name          string `json:"name" example:"Platform Cash"`
	Type          string `json:"type" example:"asset"`
	Currency      string `json:"currency" example:"KES"`
	Balance       string `json:"balance" example:"0"`
	Tenant        string `json:"tenant" example:"tenant-123"`
	AllowPostings bool   `json:"allowPostings" example:"true"`
}

type chartOfAccountsResponse struct {
	Accounts []ledgerAccount `json:"accounts"`
}

// ChartOfAccounts lists the configured ledger accounts.
// @Summary List chart of accounts
// @Description Returns the ledger chart of accounts for the requesting tenant.
// @Tags Ledger
// @Produce json
// @Param tenantID path string true "Tenant identifier"
// @Success 200 {object} chartOfAccountsResponse
// @Security bearerAuth
// @Router /{tenantID}/ledger/chart-of-accounts [get]
func (h *Ledger) ChartOfAccounts(w http.ResponseWriter, r *http.Request) {
	account := ledgerAccount{
		Code:          "1000",
		Name:          "Platform Cash",
		Type:          "asset",
		Currency:      "KES",
		Balance:       decimal.NewFromInt(0).String(),
		Tenant:        r.Header.Get("X-Tenant-ID"),
		AllowPostings: true,
	}

	respondJSON(w, http.StatusOK, chartOfAccountsResponse{
		Accounts: []ledgerAccount{account},
	})
}
