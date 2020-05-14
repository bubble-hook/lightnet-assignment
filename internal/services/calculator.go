package services

import (
	"math"

	"github.com/bubble-hook/lightnet-assignment/shared"
)

// Calculator Calculator
type Calculator struct {
}

func (c Calculator) round(raw float64) float64 {
	//round to 2  decimal places (round to nearest)
	return math.Round(raw*100) / 100
}

// Sum sum all value in CalculateRequest
func (c Calculator) Sum(reqeust shared.CalculateRequest) float64 {
	return c.round(reqeust.A + reqeust.B)
}

// Mul multiply value a by value b
func (c Calculator) Mul(reqeust shared.CalculateRequest) float64 {
	return c.round(reqeust.A * reqeust.B)
}

// Div divide value a by value b
func (c Calculator) Div(reqeust shared.CalculateRequest) float64 {
	if reqeust.B == 0 {
		return 0
	}
	return c.round(reqeust.A / reqeust.B)
}

// Sub subtract value a by value b
func (c Calculator) Sub(reqeust shared.CalculateRequest) float64 {
	return c.round(reqeust.A - reqeust.B)
}
