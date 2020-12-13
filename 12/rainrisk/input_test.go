package rainrisk

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

func TestParseInput(t *testing.T) {
	got := ParseInput(`F10
N3
F7
R90
F11`)
	want := NavigationData{
		{'F', 10},
		{'N', 3},
		{'F', 7},
		{'R', 90},
		{'F', 11},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
