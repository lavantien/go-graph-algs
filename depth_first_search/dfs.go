package depth_first_search

import "fmt"

// graph is a adjacency list represented by a list of a list of 2 value: destination and distance from the source vertex
func DepthFirstSearch(graph [][][2]int, visited []bool, currentVertex int) {
	if visited[currentVertex] {
		return
	}
	fmt.Print(currentVertex, " ")
	visited[currentVertex] = true
	for _, adjacentVertex := range graph[currentVertex] {
		DepthFirstSearch(graph, visited, adjacentVertex[0])
	}
}

// component name at a certain vertex
func ConnectedComponents(graph [][][2]int, visited []bool, currentVertex int, currentCount int, components []int) (int, []int) {
	for i := 0; i < len(graph); i++ {
		if !visited[i] {
			currentCount++
			dfs(graph, visited, i, currentCount, components)
		}
	}
	return currentCount, components
}

func dfs(graph [][][2]int, visited []bool, currentVertex int, currentCount int, components []int) {
	visited[currentVertex] = true
	components[currentVertex] = currentCount
	for _, adjacentVertex := range graph[currentVertex] {
		if !visited[adjacentVertex[0]] {
			dfs(graph, visited, adjacentVertex[0], currentCount, components)
		}
	}
}
