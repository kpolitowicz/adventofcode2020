package main

import (
	"reflect"
	"testing"
)

func TestCalcManhattanDistanceExample1(t *testing.T) {
	wire1_str := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	wire2_str := "U62,R66,U55,R34,D71,R55,D58,R83"

	got := CalcManhattanDistance(wire1_str, wire2_str)
	want := 159

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCalcManhattanDistanceExample2(t *testing.T) {
	wire1_str := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	wire2_str := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"

	got := CalcManhattanDistance(wire1_str, wire2_str)
	want := 135

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCalcSignalDistanceExample1(t *testing.T) {
	wire1_str := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	wire2_str := "U62,R66,U55,R34,D71,R55,D58,R83"

	got := CalcSignalDistance(wire1_str, wire2_str)
	want := 610

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCalcSignalDistanceExample2(t *testing.T) {
	wire1_str := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	wire2_str := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"

	got := CalcSignalDistance(wire1_str, wire2_str)
	want := 410

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestParseInputString(t *testing.T) {
	got := ParseInputString("R8,U5,L5,D3")
	want := []string{"R8", "U5", "L5", "D3"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestConvertToWireDef(t *testing.T) {
	got := ConvertToWireDef([]string{"R8", "U555"})
	want := []wireDef{wireDef{'R', 8}, wireDef{'U', 555}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetWirePoints(t *testing.T) {
	got := GetWirePoints([]wireDef{wireDef{'R', 3}, wireDef{'U', 2}, wireDef{'L', 2}, wireDef{'D', 1}})
	want := []point{
		point{1, 0},
		point{2, 0},
		point{3, 0},
		point{3, 1},
		point{3, 2},
		point{2, 2},
		point{1, 2},
		point{1, 1},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestFindCommonPoints(t *testing.T) {
	wire1 := GetWirePoints([]wireDef{
		wireDef{'R', 8},
		wireDef{'U', 5},
		wireDef{'L', 5},
		wireDef{'D', 3},
	})
	wire2 := GetWirePoints([]wireDef{
		wireDef{'U', 7},
		wireDef{'R', 6},
		wireDef{'D', 4},
		wireDef{'L', 4},
	})
	got := FindCommonPoints(wire1, wire2)
	want := []pointWithSignalDist{
		pointWithSignalDist{6, 5, 30},
		pointWithSignalDist{3, 3, 40},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestFindClosestPointByManhattanDist(t *testing.T) {
	points := []pointWithSignalDist{
		pointWithSignalDist{6, 5, 0},
		pointWithSignalDist{3, 3, 0},
	}
	got := FindClosestPointByManhattanDist(points)
	want := pointWithSignalDist{3, 3, 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestFindClosestPointByiSignalDist(t *testing.T) {
	points := []pointWithSignalDist{
		pointWithSignalDist{6, 5, 30},
		pointWithSignalDist{3, 3, 40},
	}
	got := FindClosestPointBySignalDist(points)
	want := pointWithSignalDist{6, 5, 30}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestManhattanDistance(t *testing.T) {
	got := ManhattanDistance(pointWithSignalDist{6, 5, 0})
	want := 11
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestManhattanDistanceNegativeCoords(t *testing.T) {
	got := ManhattanDistance(pointWithSignalDist{-6, -5, 0})
	want := 11
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
