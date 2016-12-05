package main

import (
	"fmt"

	"github.com/fjukstad/luftkvalitet"
)

func main() {
	aqis, err := luftkvalitet.GetAqis()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(aqis)
}
