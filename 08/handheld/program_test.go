package handheld

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

func TestExecuteCorrectProgram(t *testing.T) {
	p := ParseInput(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1`)
	program := &p

	program.Execute()
	got := program.Accumulator
	want := 2

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestExecuteInfiniteProgram(t *testing.T) {
	p := ParseInput(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)
	program := &p

	program.Execute()
	got := program.Accumulator
	want := 5

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

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
	if !program.Commands[0].Executed {
		t.Errorf("The command %v should be marked as executed", program.Commands[0])
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
	if !program.Commands[0].Executed {
		t.Errorf("The command %v should be marked as executed", program.Commands[0])
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
	if !program.Commands[0].Executed {
		t.Errorf("The command %v should be marked as executed", program.Commands[0])
	}
}

func TestExecuteNextCommandErrorIfAlreadyExecuted(t *testing.T) {
	p := ParseInput(`acc +3
acc +1`)
	program := &p
	program.Commands[0].Executed = true

	// acc +3 is marked as executed - it should not execute again
	// and return error instead
	err := program.executeNextCommand()
	if err == nil {
		t.Errorf("Thereshould be an error, but:  %v", err)
	}

	got := program.Accumulator
	want := 0
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

	got = program.Cursor
	want = 0
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
	if !program.Commands[0].Executed {
		t.Errorf("The command %v should be marked as executed", program.Commands[0])
	}
}
