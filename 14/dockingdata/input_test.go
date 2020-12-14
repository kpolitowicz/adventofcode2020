package dockingdata

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

func TestParseInput(t *testing.T) {
	got := ParseInput(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`)
	want := Program{
		MaskCmd{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"},
		MemWriteCmd{8, 11},
		MemWriteCmd{7, 101},
		MemWriteCmd{8, 0},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
