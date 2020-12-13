package rainrisk

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

func TestExecuteNavigation(t *testing.T) {
	naviData := NavigationData{
		{'F', 10},
		{'N', 3},
		{'F', 7},
		{'R', 90},
		{'F', 11},
	}
	ferry := NewFerry()
	ferry.ExecuteNavigation(naviData, (*Ferry).ExecuteCommand)

	gotPos := ferry.Pos
	wantPos := Position{17, -8}
	if !reflect.DeepEqual(gotPos, wantPos) {
		t.Errorf("got %v want %v", gotPos, wantPos)
	}

	gotFace := ferry.Facing
	wantFace := 180
	if !reflect.DeepEqual(gotFace, wantFace) {
		t.Errorf("got %v want %v", gotFace, wantFace)
	}
}

func TestExecuteNavigationByWaypoint(t *testing.T) {
	naviData := NavigationData{
		{'F', 10},
		{'N', 3},
		{'F', 7},
		{'R', 90},
		{'F', 11},
	}
	ferry := NewFerry()
	ferry.ExecuteNavigation(naviData, (*Ferry).ExecuteWaypointCommand)

	gotPos := ferry.Pos
	wantPos := Position{214, -72}
	if !reflect.DeepEqual(gotPos, wantPos) {
		t.Errorf("got %v want %v", gotPos, wantPos)
	}
}

func TestExecuteCommandN(t *testing.T) {
	cmd := NavigationCmd{'N', 10}
	ferry := NewFerry()
	ferry.ExecuteCommand(cmd)

	gotPos := ferry.Pos
	wantPos := Position{0, 10}
	if !reflect.DeepEqual(gotPos, wantPos) {
		t.Errorf("got %v want %v", gotPos, wantPos)
	}

	gotFace := ferry.Facing
	wantFace := 90
	if !reflect.DeepEqual(gotFace, wantFace) {
		t.Errorf("got %v want %v", gotFace, wantFace)
	}
}

func TestExecuteCommandS(t *testing.T) {
	cmd := NavigationCmd{'S', 10}
	ferry := NewFerry()
	ferry.ExecuteCommand(cmd)

	gotPos := ferry.Pos
	wantPos := Position{0, -10}
	if !reflect.DeepEqual(gotPos, wantPos) {
		t.Errorf("got %v want %v", gotPos, wantPos)
	}

	gotFace := ferry.Facing
	wantFace := 90
	if !reflect.DeepEqual(gotFace, wantFace) {
		t.Errorf("got %v want %v", gotFace, wantFace)
	}
}

func TestExecuteCommandE(t *testing.T) {
	cmd := NavigationCmd{'E', 10}
	ferry := NewFerry()
	ferry.ExecuteCommand(cmd)

	gotPos := ferry.Pos
	wantPos := Position{10, 0}
	if !reflect.DeepEqual(gotPos, wantPos) {
		t.Errorf("got %v want %v", gotPos, wantPos)
	}

	gotFace := ferry.Facing
	wantFace := 90
	if !reflect.DeepEqual(gotFace, wantFace) {
		t.Errorf("got %v want %v", gotFace, wantFace)
	}
}

func TestExecuteCommandW(t *testing.T) {
	cmd := NavigationCmd{'W', 1}
	ferry := NewFerry()
	ferry.ExecuteCommand(cmd)

	gotPos := ferry.Pos
	wantPos := Position{-1, 0}
	if !reflect.DeepEqual(gotPos, wantPos) {
		t.Errorf("got %v want %v", gotPos, wantPos)
	}

	gotFace := ferry.Facing
	wantFace := 90
	if !reflect.DeepEqual(gotFace, wantFace) {
		t.Errorf("got %v want %v", gotFace, wantFace)
	}
}

func TestExecuteCommandR(t *testing.T) {
	cmd := NavigationCmd{'R', 450}
	ferry := NewFerry()
	ferry.ExecuteCommand(cmd)

	gotPos := ferry.Pos
	wantPos := Position{0, 0}
	if !reflect.DeepEqual(gotPos, wantPos) {
		t.Errorf("got %v want %v", gotPos, wantPos)
	}

	gotFace := ferry.Facing
	wantFace := 180
	if !reflect.DeepEqual(gotFace, wantFace) {
		t.Errorf("got %v want %v", gotFace, wantFace)
	}
}

func TestExecuteCommandL(t *testing.T) {
	cmd := NavigationCmd{'L', 180}
	ferry := NewFerry()
	ferry.ExecuteCommand(cmd)

	gotPos := ferry.Pos
	wantPos := Position{0, 0}
	if !reflect.DeepEqual(gotPos, wantPos) {
		t.Errorf("got %v want %v", gotPos, wantPos)
	}

	gotFace := ferry.Facing
	wantFace := 270
	if !reflect.DeepEqual(gotFace, wantFace) {
		t.Errorf("got %v want %v", gotFace, wantFace)
	}
}

func TestExecuteCommandF(t *testing.T) {
	cmd := NavigationCmd{'F', 5}
	ferry := NewFerry()
	ferry.ExecuteCommand(cmd)

	gotPos := ferry.Pos
	wantPos := Position{5, 0}
	if !reflect.DeepEqual(gotPos, wantPos) {
		t.Errorf("got %v want %v", gotPos, wantPos)
	}

	gotFace := ferry.Facing
	wantFace := 90
	if !reflect.DeepEqual(gotFace, wantFace) {
		t.Errorf("got %v want %v", gotFace, wantFace)
	}
}

func TestExecuteWaypointCommandN(t *testing.T) {
	cmd := NavigationCmd{'N', 10}
	ferry := NewFerry()
	ferry.ExecuteWaypointCommand(cmd)

	gotPos := ferry.Pos
	wantPos := Position{0, 0}
	if !reflect.DeepEqual(gotPos, wantPos) {
		t.Errorf("got %v want %v", gotPos, wantPos)
	}

	gotWpPos := ferry.WaypointPos
	wantWpPos := Position{10, 11}
	if !reflect.DeepEqual(gotWpPos, wantWpPos) {
		t.Errorf("got %v want %v", gotWpPos, wantWpPos)
	}
}

func TestExecuteWaypointCommandF(t *testing.T) {
	cmd := NavigationCmd{'F', 10}
	ferry := NewFerry()
	ferry.ExecuteWaypointCommand(cmd)

	gotPos := ferry.Pos
	wantPos := Position{100, 10}
	if !reflect.DeepEqual(gotPos, wantPos) {
		t.Errorf("got %v want %v", gotPos, wantPos)
	}

	gotWpPos := ferry.WaypointPos
	wantWpPos := Position{10, 1}
	if !reflect.DeepEqual(gotWpPos, wantWpPos) {
		t.Errorf("got %v want %v", gotWpPos, wantWpPos)
	}
}

func TestExecuteWaypointCommandR(t *testing.T) {
	cmd := NavigationCmd{'R', 90}
	ferry := NewFerry()
	ferry.ExecuteWaypointCommand(cmd)

	gotPos := ferry.Pos
	wantPos := Position{0, 0}
	if !reflect.DeepEqual(gotPos, wantPos) {
		t.Errorf("got %v want %v", gotPos, wantPos)
	}

	gotWpPos := ferry.WaypointPos
	wantWpPos := Position{1, -10}
	if !reflect.DeepEqual(gotWpPos, wantWpPos) {
		t.Errorf("got %v want %v", gotWpPos, wantWpPos)
	}
}

func TestExecuteWaypointCommandL(t *testing.T) {
	cmd := NavigationCmd{'L', 90}
	ferry := NewFerry()
	ferry.ExecuteWaypointCommand(cmd)

	gotPos := ferry.Pos
	wantPos := Position{0, 0}
	if !reflect.DeepEqual(gotPos, wantPos) {
		t.Errorf("got %v want %v", gotPos, wantPos)
	}

	gotWpPos := ferry.WaypointPos
	wantWpPos := Position{-1, 10}
	if !reflect.DeepEqual(gotWpPos, wantWpPos) {
		t.Errorf("got %v want %v", gotWpPos, wantWpPos)
	}
}
