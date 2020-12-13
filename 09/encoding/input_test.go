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

func TestFindBAdNumber(t *testing.T) {
	sequence := ParseInput(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`)
	got := sequence.FindBadNumber(5)
	want := uint64(127)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGoodNumber(t *testing.T) {
	sequence := ParseInput(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`)
	got := sequence.goodNumber(5, 5)
	want := true

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
