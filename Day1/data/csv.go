package data

import (
	"errors"
	"io/ioutil"
	"strings"
)

func ReadFile(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	mydata := strings.Split(string(data), "\n")
	return mydata, err
}

func SeparateFile(line string) ([]string, error) {
	mydata := strings.Split(line, ",")
	if len(mydata) != 3 {
		return nil, errors.New("error format")
	}
	return mydata, nil
}
