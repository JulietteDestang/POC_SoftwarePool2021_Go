package data

import (
	"io/ioutil"
	"strings"
)

func ReadFile(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	mydata := strings.Split(string(data), "\n")
	return mydata, err
}
func SeparateFile(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	mydata := strings.Split(string(data), ",")
	return mydata, err
}
