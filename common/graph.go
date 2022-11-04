package common

import "fmt"

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
func MatrixToAdjacencyList(matrix [][]int) [][][2]int {
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

func PrintAdjacencyList(graph [][][2]int) {
	for i, list := range graph {
		fmt.Print(i, ": ")
		for _, vertex := range list {
			fmt.Print(vertex, " ")
		}
		fmt.Println()
	}
}

func PrintMatrix[T any](g [][]T) {
	for i := range g {
		for j := range g[i] {
			fmt.Print(g[i][j], "  ")
		}
		fmt.Println()
	}
}
