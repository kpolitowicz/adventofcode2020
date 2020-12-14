package dockingdata

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

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
