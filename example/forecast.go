package main

import "fmt"
import "github.com/fjukstad/luftkvalitet"

func main() {
	all, err := luftkvalitet.GetForecasts(luftkvalitet.Filter{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(all)

	f := luftkvalitet.Filter{Areas: []string{"Bergen"}}
	berg1, err := luftkvalitet.GetForecasts(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(berg1)

	f = luftkvalitet.Filter{Areas: []string{"Bergen", "Tromsø"}}
	berg1tromso, err := luftkvalitet.GetForecasts(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(berg1tromso)
}
