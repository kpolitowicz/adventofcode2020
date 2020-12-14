package shuttlesearch

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

func TestFindEarliestTimestampExample1(t *testing.T) {
	schedule := ParseInput(`939
7,13,x,x,59,x,31,19`)
	got := schedule.FindEarliestTimestamp()
	want := uint64(1068781)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestFindEarliestTimestampExample2(t *testing.T) {
	schedule := ParseInput(`939
17,x,13,19`)
	got := schedule.FindEarliestTimestamp()
	want := uint64(3417)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestFindEarliestTimestampExample3(t *testing.T) {
	schedule := ParseInput(`939
67,7,59,61`)
	got := schedule.FindEarliestTimestamp()
	want := uint64(754018)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestFindEarliestTimestampExample4(t *testing.T) {
	schedule := ParseInput(`939
67,x,7,59,61`)
	got := schedule.FindEarliestTimestamp()
	want := uint64(779210)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestFindEarliestTimestampExample5(t *testing.T) {
	schedule := ParseInput(`939
67,7,x,59,61`)
	got := schedule.FindEarliestTimestamp()
	want := uint64(1261476)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestFindEarliestTimestampExample6(t *testing.T) {
	schedule := ParseInput(`939
1789,37,47,1889`)
	got := schedule.FindEarliestTimestamp()
	want := uint64(1202161486)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCheckAtTimestampFromExample(t *testing.T) {
	schedule := ParseInput(`939
7,13,x,x,59,x,31,19`)
	got := schedule.checkAtTimestamp(1068781)
	want := true

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCheckAtWrongTimestamp(t *testing.T) {
	schedule := ParseInput(`939
7,13,x,x,59,x,31,19`)
	got := schedule.checkAtTimestamp(939)
	want := false

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
