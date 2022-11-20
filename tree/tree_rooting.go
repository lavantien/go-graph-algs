package tree

import "go-graph-als/common"

func RootTree(graph [][]int, rootID int) *common.Node {
	root := common.NewNode(rootID, nil, []*common.Node{})
	return buildTree(graph, root, nil)
}

func buildTree(graph [][]int, node *common.Node, parent *common.Node) *common.Node {
	for _, childID := range graph[node.ID] {
		if parent != nil && childID == parent.ID {
			continue
		}
		child := common.NewNode(childID, node, []*common.Node{})
		node.ChildNodes = append(node.ChildNodes, child)
		buildTree(graph, child, node)
	}
	return node
}
