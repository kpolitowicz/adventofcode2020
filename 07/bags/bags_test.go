package bags

import (
	// "fmt"
	"reflect"
	"testing"
)

func TestAddEdgeOneEdge(t *testing.T) {
	gotBm := NewBagMap()
	gotBm.AddEdge("container", "contained")

	wantBm := NewBagMap()
	wantBm.SetNode("container", &Node{Label: "container"})
	wantBm.SetNode("contained", &Node{Label: "contained", CanBeIn: [](*Node){wantBm.GetNode("container")}})

	assertBmEqual(t, gotBm, wantBm)
}

func TestAddEdgeTwoContainedSameContainer(t *testing.T) {
	gotBm := NewBagMap()
	gotBm.AddEdge("container", "contained1")
	gotBm.AddEdge("container", "contained2")

	wantBm := NewBagMap()
	wantBm.SetNode("container", &Node{Label: "container"})
	wantBm.SetNode("contained1", &Node{Label: "contained1", CanBeIn: [](*Node){wantBm.GetNode("container")}})
	wantBm.SetNode("contained2", &Node{Label: "contained2", CanBeIn: [](*Node){wantBm.GetNode("container")}})

	assertBmEqual(t, gotBm, wantBm)
}

func TestAddEdgeTwoContainersSameContained(t *testing.T) {
	gotBm := NewBagMap()
	gotBm.AddEdge("container1", "contained")
	gotBm.AddEdge("container2", "contained")

	wantBm := NewBagMap()
	wantBm.SetNode("container1", &Node{Label: "container1"})
	wantBm.SetNode("container2", &Node{Label: "container2"})
	wantBm.SetNode("contained", &Node{
		Label: "contained",
		CanBeIn: [](*Node){
			wantBm.GetNode("container1"),
			wantBm.GetNode("container2"),
		},
	})

	assertBmEqual(t, gotBm, wantBm)
}

func TestAddEdgeContainedInContainerInAnotherContainer(t *testing.T) {
	gotBm := NewBagMap()
	gotBm.AddEdge("container", "containedAndContainer")
	gotBm.AddEdge("containedAndContainer", "contained")

	wantBm := NewBagMap()
	wantBm.SetNode("container", &Node{Label: "container"})
	wantBm.SetNode("containedAndContainer", &Node{
		Label:   "containedAndContainer",
		CanBeIn: [](*Node){wantBm.GetNode("container")},
	})
	wantBm.SetNode("contained", &Node{
		Label:   "contained",
		CanBeIn: [](*Node){wantBm.GetNode("containedAndContainer")},
	})

	assertBmEqual(t, gotBm, wantBm)
}

func TestSearchMarksStart(t *testing.T) {
	bm := NewBagMap()
	bm.AddEdge("container", "contained")
	visited := make(VisitedMap)

	bm.Search(bm.GetNode("container"), &visited)

	assertEqual(t, visited["container"], true)
	assertEqual(t, visited["contained"], false)
}

func TestSearchVisitsConnectedNodes(t *testing.T) {
	bm := NewBagMap()
	bm.AddEdge("container", "contained")
	visited := make(VisitedMap)

	bm.Search(bm.GetNode("contained"), &visited)

	assertEqual(t, visited["container"], true)
	assertEqual(t, visited["contained"], true)
}

func TestSearchWorksOnExample(t *testing.T) {
	bm := ParseInput(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`)
	visited := make(VisitedMap)

	bm.Search(bm.GetNode("shiny gold"), &visited)

	assertEqual(t, len(visited), 5)
}

func assertBmEqual(t *testing.T, gotBm, wantBm *BagMap) {
	assertEqual(t, len(*gotBm), len(*wantBm))
	for label, _ := range *gotBm {
		assertEqual(t, gotBm.GetNode(label), wantBm.GetNode(label))
	}
}

func assertEqual(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
