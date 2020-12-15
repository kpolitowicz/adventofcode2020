package main

import (
	"fmt"
	"github.com/kpolitowicz/adventofcode2020/recitation"
	// "io/ioutil"
	"os"
	// "strings"
)

func main() {
	game := recitation.ParseInput("9,6,0,10,18,2,1")

	switch os.Args[1] {
	case "1":
		game.TakeTurnsUntil(2020)
		fmt.Println(game.LastNumber())
	case "2":
		game.TakeTurnsUntil(30000000)
		fmt.Println(game.LastNumber())
	}
}
