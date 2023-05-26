package callviacep

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func Call(cep string) (ViaCEP, error) {
	API_URL := "http://viacep.com.br/ws/" + cep + "/json/"

	req, err := http.NewRequest("GET", API_URL, nil)
	if err != nil {
		return ViaCEP{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ViaCEP{}, err
	}

	defer res.Body.Close()

	parsed, err := io.ReadAll(res.Body)
	if err != nil {
		return ViaCEP{}, err
	}

	var r ViaCEP

	err = json.Unmarshal(parsed, &r)
	if err != nil {
		return ViaCEP{}, err
	}

	return r, nil
}
