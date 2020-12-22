package conwaycubes

import (
	"reflect"
	"testing"
)

func TestCountActiveHyper(t *testing.T) {
	hu := hyperUniverse{{{
		".#.",
		"..#",
		"###",
	}}}

	// 6 cycles
	lastHu := hu.CycleOnce().CycleOnce().CycleOnce().CycleOnce().CycleOnce().CycleOnce()

	got := lastHu.CountActive()
	want := 848
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCountActive(t *testing.T) {
	pu := pocketUniverse{{
		".#.",
		"..#",
		"###",
	}}

	// 6 cycles
	lastPu := pu.CycleOnce().CycleOnce().CycleOnce().CycleOnce().CycleOnce().CycleOnce()

	got := lastPu.CountActive()
	want := 112
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

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
