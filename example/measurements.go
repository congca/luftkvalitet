package main

import (
	"fmt"

	"github.com/fjukstad/luftkvalitet"
)

func main() {

	m, err := luftkvalitet.GetMeasurements(luftkvalitet.Filter{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m, len(m))

	areas := []string{"Tromsø"}

	m, err = luftkvalitet.GetMeasurements(luftkvalitet.Filter{Areas: areas})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(m)

	stations := []string{"Hansjordnesbukta"}
	m, err = luftkvalitet.GetMeasurements(luftkvalitet.Filter{
		Stations: stations})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(m)

	components := []string{"NO2"}
	m, err = luftkvalitet.GetMeasurements(luftkvalitet.Filter{Areas: areas,
		Stations:   stations,
		Components: components})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(m)

	within := luftkvalitet.Point{Location: luftkvalitet.Location{Latitude: 69.667081, Longitude: 18.9595023}, Radius: 10}
	m, err = luftkvalitet.GetMeasurements(luftkvalitet.Filter{
		Areas:  areas,
		Within: within})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(m)

	nearest := luftkvalitet.Point{Location: luftkvalitet.Location{Latitude: 69.667081, Longitude: 18.9595023}, Radius: 2}
	m, err = luftkvalitet.GetMeasurements(luftkvalitet.Filter{
		Nearest: nearest})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(m)

}
