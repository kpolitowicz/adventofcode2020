package bags

import (
	// "fmt"
	"reflect"
	"testing"
)

func TestParseInputOneBag(t *testing.T) {
	gotBm := ParseInput(`muted yellow bags contain no other bags.`)

	wantBm := NewBagMap()
	wantBm.SetNode("muted yellow", &Node{Label: "muted yellow"})

	assertBmEqual(t, gotBm, wantBm)
}

func TestParseInputOneBagInAnother(t *testing.T) {
	gotBm := ParseInput(`muted yellow bags contain 2 shiny gold bags.
shiny gold bags contain no other bags.`)

	wantBm := NewBagMap()
	wantBm.SetNode("muted yellow", &Node{Label: "muted yellow"})
	wantBm.SetNode("shiny gold", &Node{
		Label:   "shiny gold",
		CanBeIn: [](*Node){wantBm.GetNode("muted yellow")},
	})

	assertBmEqual(t, gotBm, wantBm)
}

func TestParseInputFromExercise(t *testing.T) {
	gotBm := ParseInput(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`)

	assertEqual(t, len(*gotBm), 9)

	directContainers := gotBm.GetNode("shiny gold").CanBeIn
	wantContainers := [](*Node){
		gotBm.GetNode("bright white"),
		gotBm.GetNode("muted yellow"),
	}
	assertEqual(t, directContainers, wantContainers)
}

func TestParseInputLineNoBags(t *testing.T) {
	container, contained := parseInputLine("muted yellow bags contain no other bags.")

	wantContainer := "muted yellow"
	wantContained := []ContainsSpec{}

	if !reflect.DeepEqual(container, wantContainer) {
		t.Errorf("got %v want %v", container, wantContainer)
	}

	if !reflect.DeepEqual(contained, wantContained) {
		t.Errorf("got %v want %v", contained, wantContained)
	}
}

func TestParseInputLineOneBag(t *testing.T) {
	container, contained := parseInputLine("bright white bags contain 1 shiny gold bag.")

	wantContainer := "bright white"
	wantContained := []ContainsSpec{
		{"shiny gold", 1},
	}

	if !reflect.DeepEqual(container, wantContainer) {
		t.Errorf("got %v want %v", container, wantContainer)
	}

	if !reflect.DeepEqual(contained, wantContained) {
		t.Errorf("got %v want %v", contained, wantContained)
	}
}

func TestParseInputLineThreeBags(t *testing.T) {
	container, contained := parseInputLine("mirrored brown bags contain 1 pale teal bag, 3 muted gray bags, 4 dark bronze bags.")

	wantContainer := "mirrored brown"
	wantContained := []ContainsSpec{
		{"pale teal", 1},
		{"muted gray", 3},
		{"dark bronze", 4},
	}

	if !reflect.DeepEqual(container, wantContainer) {
		t.Errorf("got %v want %v", container, wantContainer)
	}

	if !reflect.DeepEqual(contained, wantContained) {
		t.Errorf("got %v want %v", contained, wantContained)
	}
}
