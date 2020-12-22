package conwaycubes

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

func TestParseInput(t *testing.T) {
	got := ParseInput(`.#.
..#
###`)
	want := pocketUniverse{{
		".#.",
		"..#",
		"###",
	}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
