package monstermessages

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	got := ParseInput(`35
20
15
25`)
	want := JoltageAdapters{0, 15, 20, 25, 35, 38}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
