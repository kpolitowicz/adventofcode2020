package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	got := ParseInput(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)

	want := SlopeMap{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCountTrees3and1(t *testing.T) {
	slopeMap := SlopeMap{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	got := slopeMap.CountTrees(3, 1)

	if got != 7 {
		t.Errorf("expected to be 7 trees, but was %v", got)
	}
}

func TestCountTrees1and1(t *testing.T) {
	slopeMap := SlopeMap{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	got := slopeMap.CountTrees(1, 1)
	want := 2

	if got != want {
		t.Errorf("expected to be %v trees, but was %v", want, got)
	}
}

func TestCountTrees1and2(t *testing.T) {
	slopeMap := SlopeMap{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	got := slopeMap.CountTrees(1, 2)
	want := 2

	if got != want {
		t.Errorf("expected to be %v trees, but was %v", want, got)
	}
}
