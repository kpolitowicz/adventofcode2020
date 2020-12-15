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

func TestExecuteMaskCmd(t *testing.T) {
	computer := NewComputer(1)
	maskCmd := MaskCmd{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"}
	maskCmd.ExecuteOn(computer)

	got := computer.GetMask()
	want := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestExecuteMemWrite(t *testing.T) {
	computer := NewComputer(1)
	computer.SetMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")

	memWriteCmd := MemWriteCmd{8, 11}
	memWriteCmd.ExecuteOn(computer)

	got := computer.GetMemory()[8]
	want := uint64(73)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestExecuteMemWriteVer2(t *testing.T) {
	computer := NewComputer(2)
	computer.SetMask("000000000000000000000000000000X1001X")

	memWriteCmd := MemWriteCmd{42, 100}
	memWriteCmd.ExecuteOn(computer)

	got := computer.GetMemory()
	want := MemMap{
		26: 100,
		27: 100,
		58: 100,
		59: 100,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
