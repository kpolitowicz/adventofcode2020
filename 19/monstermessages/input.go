package monstermessages

import (
	"strings"
)

type Rules map[string]string
type Messages []string

func ParseInput(input string) (Rules, Messages) {
	rules := make(Rules)
	messages := Messages{}
	readMessages := false
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			readMessages = true
			continue
		}
		if readMessages {
			messages = append(messages, line)
		} else { // read rules
			strSplit := strings.Split(line, ": ")
			rules[strSplit[0]] = strSplit[1]
		}
	}

	return rules, messages
}
