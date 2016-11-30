package luftkvalitet

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var endpoint = "https://api.nilu.no/"

type Measurement struct {
	Zone         string    `json:"zone"`
	Municipality string    `json:"municipality"`
	Area         string    `json:"area"`
	Station      string    `json:"station"`
	Eoi          string    `json:"eoi"`
	Component    string    `json:"component"`
	FromTime     time.Time `json:"fromTime"`
	ToTime       time.Time `json:"toTime"`
	Value        float64   `json:"value"`
	Unit         string    `json:"unit"`
	Index        int       `json:"index"`
	Color        string    `json:"color"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
}

type Point struct {
	Lat    string
	Long   string
	Radius string
}

type Filter struct {
	Areas      []string
	Stations   []string
	Components []string
	Within     Point
	Nearest    Point
}

func GetMeasurements(f Filter) ([]Measurement, error) {
	u := endpoint + url.QueryEscape("aq/utd.json?")

	if len(f.Areas) > 0 {
		query := url.QueryEscape("areas=" + strings.Join(f.Areas, ";"))
		u = u + query
	}

	if len(f.Stations) > 0 {
		query := url.QueryEscape("&stations=" + strings.Join(f.Stations, ";"))
		u = u + query
	}

	if len(f.Components) > 0 {
		query := url.QueryEscape("&components=" + strings.Join(f.Components,
			";"))
		u = u + query
	}

	if f.Within.Lat != "" && f.Within.Long != "" && f.Within.Radius != "" {
		query := url.QueryEscape("&within=" +
			strings.Join([]string{f.Within.Lat, f.Within.Long, f.Within.Radius},
				";"))
		u = u + query
	}

	if f.Nearest.Lat != "" && f.Nearest.Long != "" && f.Nearest.Radius != "" {
		query := url.QueryEscape("&nearest=" + strings.Join([]string{f.Nearest.Lat, f.Nearest.Long, f.Nearest.Radius}, ";"))
		u = u + query
	}

	resp, err := http.Get(u)

	if err != nil {
		return []Measurement{}, err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var measurements []Measurement
	err = json.Unmarshal(body, &measurements)
	if err != nil {
		return []Measurement{}, err
	}

	return measurements, nil

}
