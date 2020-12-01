package main

import (
	"reflect"
	"testing"
)

// the expected numbers are the lowest and highest in the sorted array
// this is the example from the exercise
func TestFindEntriesEasy(t *testing.T) {
	got := FindEntries([]int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	})
	want := []int{
		299,
		1721,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// the expected numbers are the 2nd and the last
func TestFindEntriesAfterFirst(t *testing.T) {
	got := FindEntries([]int{
		1,
		1001,
		1019,
	})
	want := []int{
		1001,
		1019,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// the expected numbers are the 2nd and 3rd
func TestFindEntriesBeforeLast(t *testing.T) {
	got := FindEntries([]int{
		1008,
		1009,
		1011,
		1500,
	})
	want := []int{
		1009,
		1011,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
