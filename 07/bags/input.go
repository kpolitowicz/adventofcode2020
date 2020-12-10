package bags

import (
	"regexp"
	"strconv"
	"strings"
)

type ContainsSpec struct {
	Label string
	Count int
}

func ParseInput(input string) *BagMap {
	bm := NewBagMap()
	for _, line := range strings.Split(input, "\n") {
		container, containedBags := parseInputLine(line)
		if len(containedBags) == 0 {
			if bm.GetNode(container) == nil {
				bm.SetNode(container, &Node{Label: container})
			}
		} else {
			for _, containedBag := range containedBags {
				bm.AddEdge(container, containedBag.Label)
			}
		}
	}

	return bm
}

func parseInputLine(line string) (string, []ContainsSpec) {
	extractBagInfo := regexp.MustCompile(`^(.*) bags contain (.*).$`)
	matches := extractBagInfo.FindSubmatch([]byte(line))
	container := string(matches[1])
	contains := string(matches[2])

	return container, parseContains(contains)
}

func parseContains(input string) (res []ContainsSpec) {
	if input == "no other bags" {
		return []ContainsSpec{}
	}

	extractBagInfo := regexp.MustCompile(`^(\d+) (.*) bags?$`)
	for _, bagSpecStr := range strings.Split(input, ", ") {
		matches := extractBagInfo.FindSubmatch([]byte(bagSpecStr))
		count, _ := strconv.Atoi(string(matches[1]))
		bag := string(matches[2])

		res = append(res, ContainsSpec{bag, count})
	}

	return
}
