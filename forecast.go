package luftkvalitet

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Result struct {
	Zone         string     `json:"zone"`
	Municipality string     `json:"municipality"`
	Area         string     `json:"area"`
	Today        []Forecast `json:"today"`
	Tomorrow     []Forecast `json:"tomorrow"`
}

type Forecast struct {
	Index        int    `json:"index"`
	Description  string `json:"description"`
	ForecastDate string `json:"forecastDate"`
	TimeOfDay    int    `json:"timeOfDay"`
}

func GetAllForecasts() ([]Result, error) {
	return getForecasts([]string{})
}

func GetForecasts(areas []string) ([]Result, error) {
	return getForecasts(areas)
}

func getForecasts(areas []string) ([]Result, error) {
	url := endpoint + "aq/forecast.json"
	if len(areas) > 1 {
		url = url + "+" + strings.Join(areas, "")
	}

	resp, err := http.Get(url)

	if err != nil {
		return []Result{}, err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var result []Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		return []Result{}, err
	}

	return result, nil

}
