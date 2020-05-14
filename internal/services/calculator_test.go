package services

import (
	"testing"

	"github.com/bubble-hook/lightnet-assignment/shared"
)

func Test_calculator_Sum(t *testing.T) {
	type args struct {
		reqeust shared.CalculateRequest
	}

	tests := []struct {
		name string

		args args
		want float64
	}{
		{"a=1,b=1", args{reqeust: shared.CalculateRequest{A: 1, B: 1}}, 2},
		{"a=2,b=3", args{reqeust: shared.CalculateRequest{A: 2, B: 3}}, 5},
		{"a=2.33,b=3.33", args{reqeust: shared.CalculateRequest{A: 2.33, B: 3.33}}, 5.66},
		{"a=2.1234,b=3.33", args{reqeust: shared.CalculateRequest{A: 2.1234, B: 3.33}}, 5.45},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Calculator{}
			if got := c.Sum(tt.args.reqeust); got != tt.want {
				t.Errorf("calculator.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculator_Mul(t *testing.T) {
	type args struct {
		reqeust shared.CalculateRequest
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"a=1,b=1", args{reqeust: shared.CalculateRequest{A: 1, B: 1}}, 1},
		{"a=2,b=3", args{reqeust: shared.CalculateRequest{A: 2, B: 3}}, 6},
		{"a=2.33,b=3.33", args{reqeust: shared.CalculateRequest{A: 2.33, B: 3.33}}, 7.76},
		{"a=2.1234,b=3.33", args{reqeust: shared.CalculateRequest{A: 2.1234, B: 3.33}}, 7.07},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Calculator{}
			if got := c.Mul(tt.args.reqeust); got != tt.want {
				t.Errorf("calculator.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculator_Sub(t *testing.T) {
	type args struct {
		reqeust shared.CalculateRequest
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"a=1,b=1", args{reqeust: shared.CalculateRequest{A: 1, B: 1}}, 0},
		{"a=2,b=3", args{reqeust: shared.CalculateRequest{A: 2, B: 3}}, -1},
		{"a=2.33,b=3.33", args{reqeust: shared.CalculateRequest{A: 2.33, B: 3.33}}, -1},
		{"a=2.1234,b=3.33", args{reqeust: shared.CalculateRequest{A: 2.1234, B: 3.33}}, -1.21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Calculator{}
			if got := c.Sub(tt.args.reqeust); got != tt.want {
				t.Errorf("calculator.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculator_Div(t *testing.T) {
	type args struct {
		reqeust shared.CalculateRequest
	}
	tests := []struct {
		name string

		args args
		want float64
	}{
		{"a=1,b=1", args{reqeust: shared.CalculateRequest{A: 1, B: 1}}, 1},
		{"a=2,b=3", args{reqeust: shared.CalculateRequest{A: 2, B: 3}}, 0.67},
		{"a=2.33,b=3.33", args{reqeust: shared.CalculateRequest{A: 2.33, B: 3.33}}, 0.70},
		{"a=2.1234,b=3.33", args{reqeust: shared.CalculateRequest{A: 2.1234, B: 3.33}}, 0.64},
		{"a=0,b=1", args{reqeust: shared.CalculateRequest{A: 0, B: 1}}, 0},
		{"a=1,b=0", args{reqeust: shared.CalculateRequest{A: 1, B: 0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Calculator{}
			if got := c.Div(tt.args.reqeust); got != tt.want {
				t.Errorf("calculator.Div() = %v, want %v", got, tt.want)
			}
		})
	}
}
