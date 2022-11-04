package breadth_first_search

import (
	"reflect"
	"testing"
)

func TestBFSMatrix(t *testing.T) {
	type args struct {
		g [][]int
		s [2]int
		e [2]int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 [][2]int
	}{
		{
			name: "OK",
			args: args{
				g: [][]int{
					{1, 1, 1, 0, 1, 1, 1},
					{1, 0, 1, 1, 1, 0, 1},
					{1, 0, 1, 1, 1, 1, 1},
					{1, 1, 0, 0, 1, 1, 1},
					{0, 1, 0, 1, 1, 0, 1},
				},
				s: [2]int{0, 0},
				e: [2]int{4, 3},
			},
			want:  10,
			want1: [][2]int{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {1, 3}, {1, 4}, {2, 4}, {3, 4}, {4, 4}, {4, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := BFSMatrix(tt.args.g, tt.args.s, tt.args.e)
			if got != tt.want {
				t.Errorf("BFSMatrix() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BFSMatrix() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
