package recitation

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

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
