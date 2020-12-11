package bags

import (
	// "fmt"
	"reflect"
	"testing"
)

func TestAddEdgeOneEdge(t *testing.T) {
	gotBm := NewBagMap()
	gotBm.AddEdge("container", "contained", 1)

	wantBm := NewBagMap()
	containerNode := &Node{Label: "container"}
	containedNode := &Node{Label: "contained", CanBeIn: [](*Node){containerNode}}
	wantBm.SetNode("container", containerNode)
	wantBm.SetNode("contained", containedNode)
	containerNode.Contains = append(containerNode.Contains, Contains{
		Node:  containedNode,
		Count: 1,
	})

	assertBmEqual(t, gotBm, wantBm)
}

func TestAddEdgeTwoContainedSameContainer(t *testing.T) {
	gotBm := NewBagMap()
	gotBm.AddEdge("container", "contained1", 2)
	gotBm.AddEdge("container", "contained2", 3)

	wantBm := NewBagMap()
	containerNode := &Node{Label: "container"}
	containedNode1 := &Node{Label: "contained1", CanBeIn: [](*Node){containerNode}}
	containedNode2 := &Node{Label: "contained2", CanBeIn: [](*Node){containerNode}}
	wantBm.SetNode("container", containerNode)
	wantBm.SetNode("contained1", containedNode1)
	wantBm.SetNode("contained2", containedNode2)
	containerNode.Contains = append(containerNode.Contains, Contains{
		Node:  containedNode1,
		Count: 2,
	})
	containerNode.Contains = append(containerNode.Contains, Contains{
		Node:  containedNode2,
		Count: 3,
	})

	assertBmEqual(t, gotBm, wantBm)
}

func TestAddEdgeTwoContainersSameContained(t *testing.T) {
	gotBm := NewBagMap()
	gotBm.AddEdge("container1", "contained", 4)
	gotBm.AddEdge("container2", "contained", 5)

	wantBm := NewBagMap()
	containerNode1 := &Node{Label: "container1"}
	containerNode2 := &Node{Label: "container2"}
	containedNode := &Node{
		Label: "contained",
		CanBeIn: [](*Node){
			containerNode1,
			containerNode2,
		},
	}
	wantBm.SetNode("container1", containerNode1)
	wantBm.SetNode("container2", containerNode2)
	wantBm.SetNode("contained", containedNode)
	containerNode1.Contains = append(containerNode1.Contains, Contains{
		Node:  containedNode,
		Count: 4,
	})
	containerNode2.Contains = append(containerNode2.Contains, Contains{
		Node:  containedNode,
		Count: 5,
	})

	assertBmEqual(t, gotBm, wantBm)
}

func TestAddEdgeContainedInContainerInAnotherContainer(t *testing.T) {
	gotBm := NewBagMap()
	gotBm.AddEdge("container", "containedAndContainer", 6)
	gotBm.AddEdge("containedAndContainer", "contained", 7)

	wantBm := NewBagMap()
	containerNode := &Node{Label: "container"}
	containedAndContainerNode := &Node{
		Label:   "containedAndContainer",
		CanBeIn: [](*Node){containerNode},
	}
	containedNode := &Node{
		Label:   "contained",
		CanBeIn: [](*Node){containedAndContainerNode},
	}
	wantBm.SetNode("container", containerNode)
	wantBm.SetNode("containedAndContainer", containedAndContainerNode)
	wantBm.SetNode("contained", containedNode)
	containerNode.Contains = append(containerNode.Contains, Contains{
		Node:  containedAndContainerNode,
		Count: 6,
	})
	containedAndContainerNode.Contains = append(containedAndContainerNode.Contains, Contains{
		Node:  containedNode,
		Count: 7,
	})

	assertBmEqual(t, gotBm, wantBm)
}

func TestSearchContainersMarksStart(t *testing.T) {
	bm := NewBagMap()
	bm.AddEdge("container", "contained", 1)
	visited := make(VisitedMap)

	bm.SearchContainers(bm.GetNode("container"), &visited)

	assertEqual(t, visited["container"], true)
	assertEqual(t, visited["contained"], false)
}

func TestSearchContainersVisitsConnectedNodes(t *testing.T) {
	bm := NewBagMap()
	bm.AddEdge("container", "contained", 1)
	visited := make(VisitedMap)

	bm.SearchContainers(bm.GetNode("contained"), &visited)

	assertEqual(t, visited["container"], true)
	assertEqual(t, visited["contained"], true)
}

func TestSearchContainersWorksOnExample(t *testing.T) {
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

	bm.SearchContainers(bm.GetNode("shiny gold"), &visited)

	assertEqual(t, len(visited), 5)
}

func TestCountBagsVisitsConnectedNodes(t *testing.T) {
	bm := NewBagMap()
	bm.AddEdge("container", "contained", 10)
	visited := make(VisitedMap)
	count := 0

	bm.CountBags(bm.GetNode("container"), &visited, &count)

	assertEqual(t, count, 10)
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
