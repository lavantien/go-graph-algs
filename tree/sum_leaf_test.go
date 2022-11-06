package tree

import (
	"go-graph-als/common"
	"testing"
)

func TestLeafSum(t *testing.T) {
	type args struct {
		node *common.Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "OK",
			args: args{
				node: &common.Node{
					ID: 5,
					ChildNodes: []*common.Node{
						{
							ID: 4,
							ChildNodes: []*common.Node{
								{
									ID: 1,
									ChildNodes: []*common.Node{
										{
											ID:         2,
											ChildNodes: []*common.Node{},
										},
										{
											ID:         9,
											ChildNodes: []*common.Node{},
										},
									},
								},
								{
									ID:         -6,
									ChildNodes: []*common.Node{},
								},
							},
						},
						{
							ID: 3,
							ChildNodes: []*common.Node{
								{
									ID:         0,
									ChildNodes: []*common.Node{},
								},
								{
									ID: 7,
									ChildNodes: []*common.Node{
										{
											ID:         8,
											ChildNodes: []*common.Node{},
										},
									},
								},
								{
									ID:         -4,
									ChildNodes: []*common.Node{},
								},
							},
						},
					},
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeafSum(tt.args.node); got != tt.want {
				t.Errorf("LeafSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeHeight(t *testing.T) {
	type args struct {
		node *common.Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "OK",
			args: args{
				node: &common.Node{
					ID: 5,
					ChildNodes: []*common.Node{
						{
							ID: 4,
							ChildNodes: []*common.Node{
								{
									ID: 1,
									ChildNodes: []*common.Node{
										{
											ID:         2,
											ChildNodes: []*common.Node{},
										},
										{
											ID:         9,
											ChildNodes: []*common.Node{},
										},
									},
								},
								{
									ID:         -6,
									ChildNodes: []*common.Node{},
								},
							},
						},
						{
							ID: 3,
							ChildNodes: []*common.Node{
								{
									ID:         0,
									ChildNodes: []*common.Node{},
								},
								{
									ID: 7,
									ChildNodes: []*common.Node{
										{
											ID:         8,
											ChildNodes: []*common.Node{},
										},
									},
								},
								{
									ID:         -4,
									ChildNodes: []*common.Node{},
								},
							},
						},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TreeHeight(tt.args.node); got != tt.want {
				t.Errorf("TreeHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
