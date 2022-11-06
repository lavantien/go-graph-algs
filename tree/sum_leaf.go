package tree

import (
	"go-graph-als/common"
)

func LeafSum(node *common.Node) int {
	if node == nil {
		return 0
	}
	if len(node.ChildNodes) == 0 {
		return node.ID
	}
	sum := 0
	for _, child := range node.ChildNodes {
		sum += LeafSum(child)
	}
	return sum
}

func TreeHeight(node *common.Node) int {
	if node == nil {
		return -1
	}
	max := -1
	for _, child := range node.ChildNodes {
		max = common.Max(max, TreeHeight(child))
	}
	return max + 1
}
