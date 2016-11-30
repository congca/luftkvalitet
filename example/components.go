package main

import (
	"fmt"

	"github.com/fjukstad/luftkvalitet"
)

func main() {
	c, err := luftkvalitet.GetComponents()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c)
}
