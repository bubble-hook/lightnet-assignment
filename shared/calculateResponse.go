package shared

// CalculateResponse CalculateResponse
//
type CalculateResponse struct {
	CalculateRequest
	Result float64 `json:"result"`
}
