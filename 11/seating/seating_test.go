package seating

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

func TestSimulate(t *testing.T) {
	seatMap := ParseInput(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)
	got := seatMap.Simulate()
	want := FerrySeating{
		"#.#L.L#.##",
		"#LLL#LL.L#",
		"L.#.L..#..",
		"#L##.##.L#",
		"#.#L.LL.LL",
		"#.#L#L#.##",
		"..L.L.....",
		"#L#L##L#L#",
		"#.LLLLLL.L",
		"#.#L#L#.##",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNextRoundAfterRound0(t *testing.T) {
	seatMap := ParseInput(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)
	got := seatMap.NextRound()
	want := FerrySeating{
		"#.##.##.##",
		"#######.##",
		"#.#.#..#..",
		"####.##.##",
		"#.##.##.##",
		"#.#####.##",
		"..#.#.....",
		"##########",
		"#.######.#",
		"#.#####.##",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNextRoundAfter2Rounds(t *testing.T) {
	seatMap := ParseInput(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)
	got := seatMap.NextRound()
	got = got.NextRound()
	want := FerrySeating{
		"#.LL.L#.##",
		"#LLLLLL.L#",
		"L.L.L..L..",
		"#LLL.LL.L#",
		"#.LL.LL.LL",
		"#.LLLL#.##",
		"..L.L.....",
		"#LLLLLLLL#",
		"#.LLLLLL.L",
		"#.#LLLL.##",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCountOccupied(t *testing.T) {
	seatMap := ParseInput(`#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##`)
	got := seatMap.CountOccupied()
	want := 37

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
