package conwaycubes

import (
	"reflect"
	"testing"
)

func TestCycleOnce(t *testing.T) {
	pu := pocketUniverse{{
		".#.",
		"..#",
		"###",
	}}

	got := pu.CycleOnce()
	want := pocketUniverse{
		{
			".....",
			".....",
			".#...",
			"...#.",
			"..#..",
		},
		{
			".....",
			".....",
			".#.#.",
			"..##.",
			"..#..",
		},
		{
			".....",
			".....",
			".#...",
			"...#.",
			"..#..",
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCountActiveNeighbors(t *testing.T) {
	pu := pocketUniverse{{
		".#.",
		"..#",
		"###",
	}}

	got := pu.countActiveNeighbors(0, 0, 2)
	want := 2
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// func TestCalcNextCycleCube(t *testing.T) {
// 	pu := pocketUniverse{{
// 		".#.",
// 		"..#",
// 		"###",
// 	}}

// 	got := pu.calcNextCycleCube()
// 	want := pocketUniverse{
// 		{
// 			".....",
// 			".....",
// 			".#...",
// 			"...#.",
// 			"..#..",
// 		},
// 		{
// 			".....",
// 			".....",
// 			".#.#.",
// 			"..##.",
// 			"..#..",
// 		},
// 		{
// 			".....",
// 			".....",
// 			".#...",
// 			"...#.",
// 			"..#..",
// 		},
// 	}
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %v want %v", got, want)
// 	}
// }
