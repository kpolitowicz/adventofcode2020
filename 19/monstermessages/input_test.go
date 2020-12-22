package monstermessages

import (
	"reflect"
	"testing"
)

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
