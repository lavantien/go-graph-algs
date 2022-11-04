package breadth_first_search

import (
	"fmt"
	"go-graph-als/common"
)

// This is the place to experiment with matrix related BFS algorithms

// g is the graph represents directly in matrix form, s is the starting position and e is the ending position
// 0 value cell means obstacles
func BFSMatrix(g [][]int, s [2]int, e [2]int) (int, [][2]int) {
	common.PrintMatrix(g)
	m := len(g)
	n := len(g[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[s[0]][s[1]] = true
	prev := make([][][2]int, m)
	for i := range prev {
		prev[i] = make([][2]int, n)
	}
	for i := range prev {
		for j := range prev[i] {
			prev[i][j] = [2]int{-1, -1}
		}
	}
	q := [][2]int{}
	q = append(q, s)
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	for len(q) != 0 {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if len(q) != 0 && q[0][0] == i && q[0][1] == j {
					topQ := q[0]
					q = q[1:]
					for k := 0; k < 4; k++ {
						x := i + dx[k]
						y := j + dy[k]
						if x < 0 || x >= m || y < 0 || y >= n {
							continue
						}
						if g[x][y] != 1 || visited[x][y] {
							continue
						}
						fmt.Println(topQ, [2]int{x, y})
						visited[x][y] = true
						prev[x][y] = topQ
						q = append(q, [2]int{x, y})
					}
				}
			}
		}
	}
	common.PrintMatrix(visited)
	common.PrintMatrix(prev)
	path := [][2]int{}
	dist := 0
	if prev[e[0]][e[1]][0] == -1 {
		return 0, [][2]int{}
	}
	for curr := e; curr != s; curr = prev[curr[0]][curr[1]] {
		path = append(path, curr)
		dist++
	}
	path = append(path, s)
	dist++
	if path[0] != e {
		return 0, [][2]int{}
	}
	common.Reverse(path)
	return dist, path
}
