package handheld

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

func TestExecuteNextCommandNop(t *testing.T) {
	p := ParseInput(`nop +0
acc +1`)
	program := &p

	// nop does nothing
	// Cursor moves to the next command
	program.executeNextCommand()
	got := program.Cursor
	want := 1

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestExecuteNextCommandAcc(t *testing.T) {
	p := ParseInput(`acc +3
acc +1`)
	program := &p

	// acc add/substracts from Acumulator
	// Cursor moves to the next Command
	program.executeNextCommand()

	got := program.Accumulator
	want := 3
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

	got = program.Cursor
	want = 1
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestExecuteNextCommandJmp(t *testing.T) {
	p := ParseInput(`jmp +2
acc +3
acc +1`)
	program := &p

	// jmp jumps Cursor
	program.executeNextCommand()

	got := program.Cursor
	want := 2
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
