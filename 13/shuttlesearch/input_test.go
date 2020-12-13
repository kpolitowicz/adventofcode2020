package shuttlesearch

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

func TestParseInput(t *testing.T) {
	got := ParseInput(`939
7,13,x,x,59,x,31,19`)
	want := BusSchedule{7, 13, 0, 0, 59, 0, 31, 19}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
