package common

type Node struct {
	ID         int
	Parent     *Node
	ChildNodes []*Node
}

func NewNode(id int, parent *Node, childNodes []*Node) *Node {
	return &Node{
		ID:         id,
		Parent:     parent,
		ChildNodes: childNodes,
	}
}
