package callapicep

import (
	"encoding/json"
	"io"
	"net/http"
)

type ApiCEP struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

func Call(cep string) (ApiCEP, error) {
	API_URL := "https://cdn.apicep.com/file/apicep/" + cep + ".json"

	req, err := http.NewRequest("GET", API_URL, nil)
	if err != nil {
		return ApiCEP{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ApiCEP{}, err
	}

	defer res.Body.Close()

	parsed, err := io.ReadAll(res.Body)
	if err != nil {
		return ApiCEP{}, err
	}

	var r ApiCEP

	err = json.Unmarshal(parsed, &r)
	if err != nil {
		return ApiCEP{}, err
	}

	return r, nil
}
