package bags

type BagMap map[string]*Node
type VisitedMap map[string]bool

type Node struct {
	Label   string
	CanBeIn [](*Node)
}

func NewBagMap() *BagMap {
	bm := make(BagMap)
	return &bm
}

func (bm *BagMap) AddEdge(container, contained string) {
	// if container node already exists, no need to change it
	// otherwise we create new Node
	var containerNode *Node
	if containerNode = bm.GetNode(container); containerNode == nil {
		containerNode = &Node{Label: container}
		bm.SetNode(container, containerNode)
	}

	// if contained node already exists, we only need to add to CanBeIn
	// otherwise we create new Node
	if containedNode := bm.GetNode(contained); containedNode != nil {
		containedNode.CanBeIn = append(containedNode.CanBeIn, containerNode)
	} else {
		containedNode = &Node{Label: contained, CanBeIn: [](*Node){containerNode}}
		bm.SetNode(contained, containedNode)
	}
}

func (bm *BagMap) GetNode(label string) *Node {
	return (*bm)[label]
}

func (bm *BagMap) SetNode(label string, node *Node) {
	(*bm)[label] = node
}

func (bm *BagMap) Search(startNode *Node, visited *VisitedMap) {
	for _, connectedNode := range startNode.CanBeIn {
		if !(*visited)[connectedNode.Label] {
			bm.Search(connectedNode, visited)
		}
	}
	(*visited)[startNode.Label] = true
}
