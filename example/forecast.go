package main

import "fmt"
import "github.com/fjukstad/luftkvalitet"

func main() {
	all, err := luftkvalitet.GetAllForecasts()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(all)

	berg1, err := luftkvalitet.GetForecasts([]string{"Bergen"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(berg1)
}
