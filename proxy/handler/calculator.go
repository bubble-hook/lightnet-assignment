package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/bubble-hook/lightnet-assignment/shared"
)

// Calculate Calculate
func Calculate(w http.ResponseWriter, r *http.Request) {

	request := shared.CalculateRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		shared.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	internalServiceEnpoint := os.Getenv("SERVICE_ENDPOINT")
	reqeustPayload, err := json.Marshal(request)
	if err != nil {
		shared.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", internalServiceEnpoint, r.RequestURI), bytes.NewBuffer(reqeustPayload))
	if err != nil {
		shared.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		shared.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer resp.Body.Close()

	response := shared.CalculateResponse{}

	decoder = json.NewDecoder(resp.Body)

	if err := decoder.Decode(&response); err != nil {
		shared.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	shared.RespondJSON(w, http.StatusOK, response)

}
