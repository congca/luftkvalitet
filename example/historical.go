package main

import (
	"fmt"
	"time"

	"github.com/fjukstad/luftkvalitet"
)

func main() {
	fromTime := time.Date(2016, time.December, 1, 0, 0, 0, 0, time.UTC)
	toTime := time.Date(2016, time.December, 4, 0, 0, 0, 0, time.UTC)

	f := luftkvalitet.Filter{
		Areas:      []string{"Tromsø"},
		Components: []string{"PM10"},
		FromTime:   fromTime,
		ToTime:     toTime,
	}

	res, err := luftkvalitet.GetHistorical(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
