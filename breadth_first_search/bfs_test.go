package breadth_first_search

import (
	"reflect"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	type args struct {
		graph [][][2]int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "OK",
			args: args{
				graph: [][][2]int{
					{{7, 0}, {9, 0}, {11, 0}},
					{},
					{},
					{{2, 0}, {4, 0}},
					{},
					{},
					{{5, 0}},
					{{3, 0}, {6, 0}, {11, 0}},
					{{1, 0}, {12, 0}},
					{{8, 0}, {10, 0}},
					{{1, 0}},
					{},
					{{2, 0}},
				},
				start: 0,
				end:   4,
			},
			want: []int{0, 7, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BreadthFirstSearch(tt.args.graph, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BreadthFirstSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrixShortestPath(t *testing.T) {
	type args struct {
		matrix [][]int
		start  [2]int
		end    [2]int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{
			name: "OK",
			args: args{
				matrix: [][]int{
					{1, 1, 1, 0, 1, 1, 1},
					{1, 0, 1, 1, 1, 0, 1},
					{1, 0, 1, 1, 1, 1, 1},
					{1, 1, 0, 0, 1, 1, 1},
					{0, 1, 0, 1, 1, 0, 1},
				},
				start: [2]int{0, 0},
				end:   [2]int{4, 3},
			},
			want:  10,
			want1: []int{0, 1, 2, 9, 10, 11, 18, 25, 32, 31},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MatrixShortestPath(tt.args.matrix, tt.args.start, tt.args.end)
			if got != tt.want {
				t.Errorf("MatrixShortestPath() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("MatrixShortestPath() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
