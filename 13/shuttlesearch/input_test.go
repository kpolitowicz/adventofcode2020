package shuttlesearch

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

func TestParseInput(t *testing.T) {
	got := ParseInput(`35
20
15
25`)
	want := JoltageAdapters{0, 15, 20, 25, 35, 38}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCountJoltageDiffs(t *testing.T) {
	joltages := ParseInput(`16
10
15
5
1
11
7
19
6
12
4`)
	got := joltages.CountJoltageDiffs()
	want := map[int]int{
		1: 7,
		3: 5,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCountChainsExample1(t *testing.T) {
	joltages := ParseInput(`16
10
15
5
1
11
7
19
6
12
4`)
	got := joltages.CountChains()
	want := uint64(8)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCountChainsExample2(t *testing.T) {
	joltages := ParseInput(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`)
	got := joltages.CountChains()
	want := uint64(19208)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
