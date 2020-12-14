package dockingdata

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

func TestExecute(t *testing.T) {
	computer := NewComputer()
	program := Program{
		MaskCmd{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"},
		MemWriteCmd{8, 11},
		MemWriteCmd{7, 101},
		MemWriteCmd{8, 0},
	}
	computer.Execute(&program)

	got := computer.GetMemory()
	want := MemMap{
		7: 101,
		8: 64,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestOrMask(t *testing.T) {
	computer := NewComputer()
	computer.SetMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")

	got := computer.OrMask()
	want := uint64(64)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestAndMask(t *testing.T) {
	computer := NewComputer()
	computer.SetMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")

	got := computer.AndMask()
	want := uint64(68719476733)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
