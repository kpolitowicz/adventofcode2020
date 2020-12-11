package bags

type BagMap map[string]*Node
type VisitedMap map[string]bool
type Contains struct {
	Count int
	Node  *Node
}

type Node struct {
	Label    string
	CanBeIn  [](*Node)
	Contains []Contains
}

func NewBagMap() *BagMap {
	bm := make(BagMap)
	return &bm
}

func (bm *BagMap) AddEdge(container, contained string, count int) {
	var containerNode, containedNode *Node

	// if container node already exists, no need to change it
	// otherwise we create new Node
	if containerNode = bm.GetNode(container); containerNode == nil {
		containerNode = &Node{Label: container}
		bm.SetNode(container, containerNode)
	}

	// if contained node already exists, we only need to add to CanBeIn
	// otherwise we create new Node

	if containedNode = bm.GetNode(contained); containedNode != nil {
		containedNode.CanBeIn = append(containedNode.CanBeIn, containerNode)
	} else {
		containedNode = &Node{Label: contained, CanBeIn: [](*Node){containerNode}}
		bm.SetNode(contained, containedNode)
	}

	// and now we add to containers.Contains
	containerNode.Contains = append(containerNode.Contains, Contains{
		Node:  containedNode,
		Count: count,
	})
}

func (bm *BagMap) GetNode(label string) *Node {
	return (*bm)[label]
}

func (bm *BagMap) SetNode(label string, node *Node) {
	(*bm)[label] = node
}

func (bm *BagMap) SearchContainers(startNode *Node, visited *VisitedMap) {
	for _, connectedNode := range startNode.CanBeIn {
		if !(*visited)[connectedNode.Label] {
			bm.SearchContainers(connectedNode, visited)
		}
	}
	(*visited)[startNode.Label] = true
}

func (bm *BagMap) CountBags(startNode *Node, visited *VisitedMap, count *int) {
	for _, edge := range startNode.Contains {
		if !(*visited)[edge.Node.Label] {
			(*count) += edge.Count
			bm.CountBags(edge.Node, visited, count)
		}
	}
	(*visited)[startNode.Label] = true
}
