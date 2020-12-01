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
	// fmt.Println(expenseReport)

	res := FindEntries(expenseReport)
	fmt.Println(res[0] * res[1])

	// fmt.Println(CalculateOrbits(cleaned))
	// fmt.Println(TotalTransfers(cleaned, "YOU", "SAN"))
}
