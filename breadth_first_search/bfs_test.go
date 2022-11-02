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
