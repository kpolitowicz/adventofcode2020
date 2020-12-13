package adapters

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

func TestParseInput(t *testing.T) {
	got := ParseInput(`35
20
15
25`)
	want := JoltageAdapters{35, 20, 15, 25}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCountJoltageDiffs(t *testing.T) {
	joltages := ParseInput(`16
10
15
5
1
11
7
19
6
12
4`)
	got := joltages.CountJoltageDiffs()
	want := map[int]int{
		1: 7,
		3: 5,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
