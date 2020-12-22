package monstermessages

import (
	"strings"
)

type Rules map[string]string

func (r Rules) ResolveRule(rule string) (res string) {
	ruleStr := r[rule]

	switch {
	case strings.ContainsRune(ruleStr, '"'):
		res = string(ruleStr[1])
	case strings.ContainsRune(ruleStr, '|'):
		var resParts []string
		for _, part := range strings.Split(ruleStr, " | ") {
			resParts = append(resParts, r.resolveSequence(part))
		}
		res = "(" + strings.Join(resParts, "|") + ")"
	default:
		res = r.resolveSequence(ruleStr)
	}
	return
}

func (r Rules) resolveSequence(str string) (res string) {
	for _, rule := range strings.Split(str, " ") {
		res += r.ResolveRule(rule)
	}
	return
}

type Message string

func ParseInput(input string) (Rules, []Message) {
	rules := make(Rules)
	messages := []Message{}
	readMessages := false
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			readMessages = true
			continue
		}
		if readMessages {
			messages = append(messages, Message(line))
		} else { // read rules
			strSplit := strings.Split(line, ": ")
			rules[strSplit[0]] = strSplit[1]
		}
	}

	return rules, messages
}
