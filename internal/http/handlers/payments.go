package handlers

import "net/http"

// Payments exposes orchestration endpoints for payment intents and disbursements.
type Payments struct{}

func NewPayments() *Payments { return &Payments{} }

type paymentIntent struct {
	ID       string `json:"id" example:"pi_123"`
	Status   string `json:"status" example:"pending"`
	Amount   string `json:"amount" example:"0.00"`
	Currency string `json:"currency" example:"KES"`
}

type paymentIntentsResponse struct {
	Intents []paymentIntent `json:"intents"`
}

// Intents lists payment intents registered for the tenant.
// @Summary List payment intents
// @Description Returns the payment intents that have been created for the tenant.
// @Tags Payments
// @Produce json
// @Param tenantID path string true "Tenant identifier"
// @Success 200 {object} paymentIntentsResponse
// @Router /{tenantID}/payments/intents [get]
func (h *Payments) Intents(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, paymentIntentsResponse{
		Intents: []paymentIntent{},
	})
}
