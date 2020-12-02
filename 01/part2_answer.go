package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	cleaned := strings.Trim(string(dat), "\n")
	expenseReport := ConvertInputToListOfNumbers(cleaned)

	res := FindEntries3(expenseReport)
	fmt.Println(expenseReport)
	fmt.Println(res)
	fmt.Println(res[0] + res[1] + res[2])
	fmt.Println(res[0] * res[1] * res[2])
}
