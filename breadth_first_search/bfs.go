package breadth_first_search

import "fmt"

func BreadthFirstSearch(graph [][][2]int, start int, end int) []int {
	queue := []int{}
	visited := make([]bool, len(graph))
	prev := make([]int, len(graph))
	visited[start] = true
	prev[start] = -1
	queue = append(queue, start)
	for len(queue) != 0 {
		for _, adjacentVertex := range graph[queue[0]] {
			if !visited[adjacentVertex[0]] {
				visited[adjacentVertex[0]] = true
				prev[adjacentVertex[0]] = queue[0]
				queue = append(queue, adjacentVertex[0])
			}
		}
		queue = queue[1:]
	}
	path := []int{}
	for current := end; current >= start; current = prev[current] {
		path = append(path, current)
	}
	if path[0] != end {
		return []int{}
	}
	reverse(path)
	return path
}

func reverse(list []int) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}

/*
Suppose we have a matrix of size MxN = 4x7
col	    00 01 02 03 04 05 06
row 00: 00 01 02 03 04 05 06
row 01: 07 08 09 10 11 12 13
row 02: 14 15 16 17 18 19 20
row 03: 21 22 23 24 25 26 27
Then we can derive the cordinate (x, y) via a single cell number K like this: x = K/N and y = K-x*N

Suppose we have this matrix of size 5x7 that include obstacles
1 1 1 0 1 1 1
1 0 1 1 1 0 1
1 0 1 1 1 1 1
1 1 0 0 1 1 1
0 1 0 1 1 0 1
With starting vertex (0, 0) and ending vertex (4, 3), the shortest path is [0, 1, 2, 9, 10, 11, 18, 25, 32, 31]
Our adjacency list is composed of all K vertexes, each vertex contains a list of its adjacent vertexes, and each element of that list is
an array consists of 2 elements: Kth and W, K is its vertex number and W is its weight, in this case W is either 1 or 0 (obstacle)

To reconstruct K from any (x, y): K = x*N + y
*/
// Start and End here is a vector of position {0, 1}
func MatrixShortestPath(matrix [][]int, start [2]int, end [2]int) (int, []int) {
	n := len(matrix[0])
	return matrixShortestPath(matrixToAdjacencyList(matrix), start[0]*n+start[1], end[0]*n+end[1])
}

// Start and End here is the vertexes' order number
func matrixShortestPath(graph [][][2]int, start int, end int) (int, []int) {
	printAdjacencyList(graph)
	queue := []int{}
	visited := make([]bool, len(graph))
	prev := make([]int, len(graph))
	visited[start] = true
	prev[start] = -1
	queue = append(queue, start)
	for len(queue) != 0 {
		for _, adjacentVertex := range graph[queue[0]] {
			if !visited[adjacentVertex[0]] {
				visited[adjacentVertex[0]] = true
				fmt.Println("v:", adjacentVertex[0], ", q:", queue)
				prev[adjacentVertex[0]] = queue[0]
				queue = append(queue, adjacentVertex[0])
			}
		}
		queue = queue[1:]
	}
	path := []int{}
	distance := 0
	for current := end; current >= start; current = prev[current] {
		path = append(path, current)
		distance += 1
	}
	if path[0] != end {
		return 0, []int{}
	}
	reverse(path)
	return distance, path
}

/*
(x-1, y-1) | (x-1, y+0) | (x-1, y+1)
(x+0, y-1) |   (x, y)   | (x+0, y+1)
(x+1, y-1) | (x+1, y+0) | (x+1, y+1)
For only 4 directions:
kX: [-1, 0, 1, 0]
kY: [0, 1, 0, -1]
currX = x + kX
currY = y + kY
*/
func matrixToAdjacencyList(matrix [][]int) [][][2]int {
	ret := [][][2]int{}
	m := len(matrix)
	n := len(matrix[0])
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			currentList := [][2]int{}
			for k := 0; k < 4; k++ {
				x := i + dx[k]
				y := j + dy[k]
				if x < 0 || x >= m || y < 0 || y >= n {
					continue
				}
				if matrix[x][y] != 0 {
					currentList = append(currentList, [2]int{x*n + y, 1})
				}
			}
			ret = append(ret, currentList)
		}
	}
	return ret
}

func printAdjacencyList(graph [][][2]int) {
	for i, list := range graph {
		fmt.Print(i, ": ")
		for _, vertex := range list {
			fmt.Print(vertex, " ")
		}
		fmt.Println()
	}
}
