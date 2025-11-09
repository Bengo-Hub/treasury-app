package handlers

import "net/http"

// Payments exposes orchestration endpoints for payment intents and disbursements.
type Payments struct{}

func NewPayments() *Payments { return &Payments{} }

func (h *Payments) Intents(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]any{
		"intents": []any{},
	})
}
