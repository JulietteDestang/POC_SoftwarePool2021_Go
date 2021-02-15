package main

import (
	"SofwareGoDay1/data"
	"fmt"
)

func main() {
	mydata, err := data.SeparateFile("Go_CSV/example.csv")
	if err != nil {
		fmt.Println(err)
	}
	for i := range mydata {
		fmt.Println(mydata[i])
	}
}
