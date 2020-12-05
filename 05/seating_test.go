package main

import (
	// "fmt"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	got := ParseInput(`FBFBBFFRLR
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`)

	want := []BoardingPass{
		{AirlineSpec: "FBFBBFFRLR"},
		{AirlineSpec: "BFFFBBFRRR"},
		{AirlineSpec: "FFFBBBFRRR"},
		{AirlineSpec: "BBFFBBFRLL"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestFindMaxSeatId(t *testing.T) {
	got := FindMaxSeatId([]BoardingPass{
		{AirlineSpec: "FBFBBFFRLR"},
		{AirlineSpec: "BFFFBBFRRR"},
		{AirlineSpec: "FFFBBBFRRR"},
		{AirlineSpec: "BBFFBBFRLL"},
	})
	want := 820

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetRow(t *testing.T) {
	pass := BoardingPass{"FBFBBFFRLR"}
	got := pass.GetRow()
	want := 44

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetCol(t *testing.T) {
	pass := BoardingPass{"FBFBBFFRLR"}
	got := pass.GetCol()
	want := 5

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetSeatId(t *testing.T) {
	pass := BoardingPass{"FBFBBFFRLR"}
	got := pass.GetSeatId()
	want := 357

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
