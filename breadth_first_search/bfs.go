package breadth_first_search

import "fmt"

func BreadthFirstSearch(graph [][][2]int, start int, end int) []int {
	qu := []int{}
	visited := make([]bool, len(graph))
	prev := make([]int, len(graph))
	prev[start] = -1
	visited[start] = true
	qu = append(qu, start)
	for len(qu) != 0 {
		first := qu[0]
		qu = qu[1:]
		for _, adjacentVertex := range graph[first] {
			if !visited[adjacentVertex[0]] {
				qu = append(qu, adjacentVertex[0])
				visited[adjacentVertex[0]] = true
				prev[adjacentVertex[0]] = first
			}
		}
	}
	path := []int{}
	for i := end; i >= 0; i = prev[i] {
		path = append(path, i)
	}
	reverse(path)
	fmt.Println(path)
	if path[0] == start {
		return path
	}
	return []int{}
}

func reverse(list []int) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}
