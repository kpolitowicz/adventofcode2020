package customs

import (
	// "fmt"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	got := ParseInput(`abc

a
b
c

ab
ac

a
a
a
a

b`)

	want := []CustomGroup{
		{RawInputs: []string{"abc"}},
		{RawInputs: []string{"a", "b", "c"}},
		{RawInputs: []string{"ab", "ac"}},
		{RawInputs: []string{"a", "a", "a", "a"}},
		{RawInputs: []string{"b"}},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllQuestions(t *testing.T) {
	got := SumAllQuestions([]CustomGroup{
		{RawInputs: []string{"abc"}},
		{RawInputs: []string{"a", "b", "c"}},
		{RawInputs: []string{"ab", "ac"}},
		{RawInputs: []string{"a", "a", "a", "a"}},
		{RawInputs: []string{"b"}},
	}, CustomGroup.CountAnyoneYeses)
	want := 11

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCountAnyoneYesesForOnePerson(t *testing.T) {
	group := CustomGroup{[]string{"abc"}}
	got := group.CountAnyoneYeses()
	want := 3

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCountEveryoneYesesForOnePerson(t *testing.T) {
	group := CustomGroup{[]string{"abc"}}
	got := group.CountEveryoneYeses()
	want := 3

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCountEveryoneYesesFor3AllDifferent(t *testing.T) {
	group := CustomGroup{[]string{"a", "b", "c"}}
	got := group.CountEveryoneYeses()
	want := 0

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCountEveryoneYesesFor4AllSame(t *testing.T) {
	group := CustomGroup{[]string{"a", "a", "a", "a"}}
	got := group.CountEveryoneYeses()
	want := 1

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
