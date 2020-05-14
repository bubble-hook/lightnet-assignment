package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bubble-hook/lightnet-assignment/shared"
)

func PrintError(operator string, value struct {
	aValue float64
	bValue float64
}, err error) {
	fmt.Printf("\n[ERR][%v][a=%v,b=%v] : %v", operator, value.aValue, value.bValue, err.Error())
}

func main() {

	tasks := map[string][]struct {
		aValue float64
		bValue float64
	}{}

	tasks["sum"] = []struct {
		aValue float64
		bValue float64
	}{
		{aValue: 1, bValue: 2},
		{aValue: 3, bValue: 4},
		{aValue: 1, bValue: 3},
	}

	tasks["mul"] = []struct {
		aValue float64
		bValue float64
	}{
		{aValue: 1, bValue: 2},
		{aValue: 3, bValue: 4},
		{aValue: 1, bValue: 3},
	}

	tasks["div"] = []struct {
		aValue float64
		bValue float64
	}{
		{aValue: 1, bValue: 2},
		{aValue: 3, bValue: 4},
		{aValue: 1, bValue: 3},
	}

	tasks["sub"] = []struct {
		aValue float64
		bValue float64
	}{
		{aValue: 1, bValue: 2},
		{aValue: 3, bValue: 4},
		{aValue: 1, bValue: 3},
	}

	for operator, values := range tasks {
		for _, value := range values {
			reqeustPayload, err := json.Marshal(shared.CalculateRequest{
				A: value.aValue,
				B: value.bValue,
			})
			if err != nil {
				PrintError(operator, value, err)
				continue
			}

			req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", "http://localhost:9001/calculator.", operator), bytes.NewBuffer(reqeustPayload))
			if err != nil {
				PrintError(operator, value, err)
				continue
			}
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{Timeout: time.Second * 10}
			resp, err := client.Do(req)
			if err != nil {
				PrintError(operator, value, err)
				return
			}
			defer resp.Body.Close()

			response := shared.CalculateResponse{}

			decoder := json.NewDecoder(resp.Body)

			if err := decoder.Decode(&response); err != nil {
				PrintError(operator, value, err)
				return
			}

			fmt.Printf("\n[Success][%v][a=%v,b=%v] : %v", operator, value.aValue, value.bValue, response.Result)
		}

		fmt.Println("")
	}

}
