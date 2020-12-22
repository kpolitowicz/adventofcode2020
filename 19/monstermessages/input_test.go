package monstermessages

import (
	"reflect"
	"regexp"
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
	wantMsg := []Message{
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

func TestIsValidExampleValid(t *testing.T) {
	strRule := "^a((aa|bb)(ab|ba)|(ab|ba)(aa|bb))b$"
	validRule := regexp.MustCompile(strRule)
	message := Message("ababbb")

	if !message.IsValid(validRule) {
		t.Errorf("the message should be valid: %v", message)
	}
}

func TestIsValidExampleInvalid(t *testing.T) {
	strRule := "^a((aa|bb)(ab|ba)|(ab|ba)(aa|bb))b$"
	validRule := regexp.MustCompile(strRule)
	message := Message("bababa")

	if message.IsValid(validRule) {
		t.Errorf("the message should be invalid: %v", message)
	}
}

func TestIsValidExampleInvalid2(t *testing.T) {
	strRule := "^a((aa|bb)(ab|ba)|(ab|ba)(aa|bb))b$"
	validRule := regexp.MustCompile(strRule)
	message := Message("aaaabbb")

	if message.IsValid(validRule) {
		t.Errorf("the message should be invalid: %v", message)
	}
}

func TestIsValidCountNoLoop(t *testing.T) {
	rules, messages := ParseInput(`42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`)
	ruleStr := rules.ResolveRule("0")
	validMsg := regexp.MustCompile("^" + ruleStr + "$")
	got := ValidCount(messages, validMsg)
	want := 3

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
