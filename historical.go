package luftkvalitet

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type HistoricalResult struct {
	Area
	Station
	Eoi       string `json:"eoi"`
	Component string `json:"component"`
	Location
	Measurements []Measurement `json:"values"`
}

func GetHistorical(f Filter) ([]HistoricalResult, error) {
	u := endpoint + "aq/historical/" + strings.Join(f.Components, ",")
	u = addFilter(u, f)

	resp, err := http.Get(u)

	if err != nil {
		return []HistoricalResult{}, err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var results []HistoricalResult
	err = json.Unmarshal(body, &results)
	if err != nil {
		return []HistoricalResult{}, err
	}

	return results, nil
}
