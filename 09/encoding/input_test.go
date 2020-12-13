package encoding

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
	want := XmasData{35, 20, 15, 25}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
