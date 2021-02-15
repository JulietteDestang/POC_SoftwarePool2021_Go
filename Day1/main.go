package main

import (
	"SofwareGoDay1/data"
	"SofwareGoDay1/humanity"
	"fmt"
)

func main() {
	mydata, err := data.ReadFile("Go_CSV/example.csv")
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range mydata {
		content, error := data.SeparateFile(line)
		tab, err := humanity.NewHumanFromCSV(content)
		if error != nil {
			fmt.Println(err)
		}
		fmt.Println(tab)
	}
	content, err := humanity.NewHumansFromCsvFile("Go_CSV/big_csv.csv")
	fmt.Printf("%v\n", content)
}
