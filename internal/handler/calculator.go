package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bubble-hook/lightnet-assignment/internal/services"
	"github.com/bubble-hook/lightnet-assignment/shared"
)

// Calculate Calculate
func Calculate(operator string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := shared.CalculateRequest{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&request); err != nil {
			shared.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}
		c := services.Calculator{}
		availableOperator := map[string]func(shared.CalculateRequest) float64{"sum": c.Sum, "mul": c.Mul, "div": c.Div, "sub": c.Sub}
		if calculateFunc, ok := availableOperator[operator]; ok {
			shared.RespondJSON(w, http.StatusOK, shared.CalculateResponse{
				CalculateRequest: request,
				Result:           calculateFunc(request),
			})
		} else {
			shared.RespondError(w, http.StatusBadRequest, fmt.Sprintf("%v %v", operator, " unavailable"))
		}
	}
}
