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
	got := seatMap.Simulate(FerrySeating.TransformRow)
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

func TestSimulateBetterModel(t *testing.T) {
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
	got := seatMap.Simulate(FerrySeating.TransformRowBetterModel)
	want := FerrySeating{
		"#.L#.L#.L#",
		"#LLLLLL.LL",
		"L.L.L..#..",
		"##L#.#L.L#",
		"L.L#.LL.L#",
		"#.LLLL#.LL",
		"..#.L.....",
		"LLL###LLL#",
		"#.LLLLL#.L",
		"#.L#LL#.L#",
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
	got := seatMap.NextRound(FerrySeating.TransformRow)
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
	got := seatMap.NextRound(FerrySeating.TransformRow)
	got = got.NextRound(FerrySeating.TransformRow)
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

func TestNextRoundAfter2RoundsBetterModel(t *testing.T) {
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
	got := seatMap.NextRound(FerrySeating.TransformRowBetterModel)
	got = got.NextRound(FerrySeating.TransformRowBetterModel)
	want := FerrySeating{
		"#.LL.LL.L#",
		"#LLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLL#",
		"#.LLLLLL.L",
		"#.LLLLL.L#",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNextRoundAfter3RoundsBetterModel(t *testing.T) {
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
	got := seatMap.NextRound(FerrySeating.TransformRowBetterModel)
	got = got.NextRound(FerrySeating.TransformRowBetterModel)
	got = got.NextRound(FerrySeating.TransformRowBetterModel)
	want := FerrySeating{
		"#.L#.##.L#",
		"#L#####.LL",
		"L.#.#..#..",
		"##L#.##.##",
		"#.##.#L.##",
		"#.#####.#L",
		"..#.#.....",
		"LLL####LL#",
		"#.L#####.L",
		"#.L####.L#",
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

func TestCountOccupiedInSight8(t *testing.T) {
	seatMap := ParseInput(`.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`)
	got := seatMap.countOccupiedInSight(4, 3)
	want := 8

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
