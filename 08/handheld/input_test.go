package handheld

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

func TestParseInput(t *testing.T) {
	got := ParseInput(`nop +0
jmp -3
acc -99
acc +1`)
	want := Program{
		Commands: []Command{
			{Instruction: "nop", Arg: 0},
			{Instruction: "jmp", Arg: -3},
			{Instruction: "acc", Arg: -99},
			{Instruction: "acc", Arg: 1},
		},
		Cursor: 0,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestParseInputLineNop(t *testing.T) {
	got := parseInputLine(`nop +0`)

	want := Command{
		Instruction: "nop",
		Arg:         0,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestParseInputLineJmpMinus(t *testing.T) {
	got := parseInputLine(`jmp -583`)

	want := Command{
		Instruction: "jmp",
		Arg:         -583,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
