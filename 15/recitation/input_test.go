package recitation

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

func TestTakeTurnsUntilExample1(t *testing.T) {
	game := &MemoryGame{
		Numbers: []int{0, 3, 6},
		LastSpoken: map[int]int{
			0: 0,
			3: 1,
			6: 2,
		},
	}
	game.TakeTurnsUntil(2020)

	got := game.LastNumber()
	want := 436
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestTakeTurnsUntilExample7(t *testing.T) {
	game := ParseInput("3,1,2")
	game.TakeTurnsUntil(2020)

	got := game.LastNumber()
	want := 1836
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestTakeTurnsUntilExample2(t *testing.T) {
	game := ParseInput("1,3,2")
	game.TakeTurnsUntil(2020)

	got := game.LastNumber()
	want := 1
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNextTurnNotSpokenBefore(t *testing.T) {
	game := &MemoryGame{
		Numbers: []int{0, 3, 6},
		LastSpoken: map[int]int{
			0: 0,
			3: 1,
			6: 2,
		},
	}
	game.NextTurn()

	want := &MemoryGame{
		Numbers: []int{0, 3, 6, 0},
		LastSpoken: map[int]int{
			0: 0,
			3: 1,
			6: 2,
		},
	}
	if !reflect.DeepEqual(game, want) {
		t.Errorf("got %v want %v", game, want)
	}
}

func TestNextTurnSpokenBefore(t *testing.T) {
	game := &MemoryGame{
		Numbers: []int{0, 3, 6, 0},
		LastSpoken: map[int]int{
			0: 0,
			3: 1,
			6: 2,
		},
	}
	game.NextTurn()

	want := &MemoryGame{
		Numbers: []int{0, 3, 6, 0, 3},
		LastSpoken: map[int]int{
			0: 3,
			3: 1,
			6: 2,
		},
	}
	if !reflect.DeepEqual(game, want) {
		t.Errorf("got %v want %v", game, want)
	}
}

func TestParseInput(t *testing.T) {
	got := ParseInput("0,3,6")
	want := &MemoryGame{
		Numbers: []int{0, 3, 6},
		LastSpoken: map[int]int{
			0: 0,
			3: 1,
			6: 2,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
