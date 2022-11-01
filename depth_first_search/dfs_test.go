package depth_first_search

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDepthFirstSearch(t *testing.T) {
	type args struct {
		graph         [][][2]int
		visited       []bool
		currentVertex int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "OK",
			args: args{
				graph: [][][2]int{
					{{1, 0}, {2, 0}, {3, 0}},
					{{0, 0}, {2, 0}},
					{{0, 0}, {1, 0}, {4, 0}},
					{{0, 0}},
					{{2, 0}},
				},
				visited:       []bool{false, false, false, false, false},
				currentVertex: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DepthFirstSearch(tt.args.graph, tt.args.visited, tt.args.currentVertex)
			fmt.Println()
		})
	}
}

func TestConnectedComponents(t *testing.T) {
	type args struct {
		graph         [][][2]int
		visited       []bool
		currentVertex int
		currentCount  int
		components    []int
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
				graph: [][][2]int{
					{{8, 0}},
					{{5, 0}},
					{{9, 0}},
					{{9, 0}},
					{{0, 0}},
					{{16, 0}, {17, 0}},
					{{11, 0}},
					{{6, 0}},
					{{4, 0}, {14, 0}},
					{{15, 0}},
					{},
					{{7, 0}},
					{},
					{{0, 0}},
					{{0, 0}, {13, 0}},
					{{2, 0}, {10, 0}},
					{},
					{},
				},
				visited:       make([]bool, 18),
				currentVertex: 0,
				currentCount:  0,
				components:    make([]int, 18),
			},
			want:  6,
			want1: []int{1, 2, 3, 4, 1, 2, 5, 5, 1, 3, 3, 5, 6, 1, 1, 3, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ConnectedComponents(tt.args.graph, tt.args.visited, tt.args.currentVertex, tt.args.currentCount, tt.args.components)
			if got != tt.want {
				t.Errorf("ConnectedComponents() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ConnectedComponents() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
