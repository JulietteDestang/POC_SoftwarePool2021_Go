package humanity

import (
	"SofwareGoDay1/data"
	"fmt"
	"strconv"
)

type Human struct {
	Name    string
	Age     int
	Country string
}

func NewHumanFromCSV(csv []string) (*Human, error) {
	temp, err := strconv.Atoi(csv[1])
	human := Human{csv[0], temp, csv[2]}
	return &human, err
}

func NewHumansFromCsvFile(path string) ([]*Human, error) {
	mydata, err := data.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	lendata := len(mydata)
	human := make([]*Human, lendata)
	for i, line := range mydata {
		content, err := data.SeparateFile(line)
		tab, err := NewHumanFromCSV(content)
		if err != nil {
			fmt.Println(err)
		}
		human[i], err = NewHumanFromCSV(content)
		fmt.Println(tab)
		if err != nil {
			fmt.Println(err)
		}
	}
	return human, err
}

func (h *Human) String() string {
	return fmt.Sprintf("%s, %d, %s\n", h.Name, h.Age, h.Country)
}
