package monstermessages

import (
	"reflect"
	"testing"
)

func TestResolveRuleSimple(t *testing.T) {
	rules := Rules{
		"0": "4 1 5",
		"1": "2 3 | 3 2",
		"2": "4 4 | 5 5",
		"3": "4 5 | 5 4",
		"4": "\"a\"",
		"5": "\"b\"",
	}
	got := rules.ResolveRule("5")
	want := "b"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestResolveRuleOr(t *testing.T) {
	rules := Rules{
		"0": "4 1 5",
		"1": "2 3 | 3 2",
		"2": "4 4 | 5 5",
		"3": "4 5 | 5 4",
		"4": "\"a\"",
		"5": "\"b\"",
	}
	got := rules.ResolveRule("3")
	want := "(ab|ba)"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestResolveRuleZero(t *testing.T) {
	rules := Rules{
		"0": "4 1 5",
		"1": "2 3 | 3 2",
		"2": "4 4 | 5 5",
		"3": "4 5 | 5 4",
		"4": "\"a\"",
		"5": "\"b\"",
	}
	got := rules.ResolveRule("0")
	rule2 := "(aa|bb)"
	rule3 := "(ab|ba)"
	rule1 := "(" + rule2 + rule3 + "|" + rule3 + rule2 + ")"
	want := "a" + rule1 + "b"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestParseInput(t *testing.T) {
	gotRules, gotMsg := ParseInput(`0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`)
	wantRules := Rules{
		"0": "4 1 5",
		"1": "2 3 | 3 2",
		"2": "4 4 | 5 5",
		"3": "4 5 | 5 4",
		"4": "\"a\"",
		"5": "\"b\"",
	}
	wantMsg := Messages{
		"ababbb",
		"bababa",
		"abbbab",
		"aaabbb",
		"aaaabbb",
	}

	if !reflect.DeepEqual(gotRules, wantRules) {
		t.Errorf("got %v want %v", gotRules, wantRules)
	}
	if !reflect.DeepEqual(gotMsg, wantMsg) {
		t.Errorf("got %v want %v", gotMsg, wantMsg)
	}
}
