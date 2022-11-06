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
					Value: 5,
					ChildNodes: []*common.Node{
						{
							Value: 4,
							ChildNodes: []*common.Node{
								{
									Value: 1,
									ChildNodes: []*common.Node{
										{
											Value:      2,
											ChildNodes: []*common.Node{},
										},
										{
											Value:      9,
											ChildNodes: []*common.Node{},
										},
									},
								},
								{
									Value:      -6,
									ChildNodes: []*common.Node{},
								},
							},
						},
						{
							Value: 3,
							ChildNodes: []*common.Node{
								{
									Value:      0,
									ChildNodes: []*common.Node{},
								},
								{
									Value: 7,
									ChildNodes: []*common.Node{
										{
											Value:      8,
											ChildNodes: []*common.Node{},
										},
									},
								},
								{
									Value:      -4,
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
					Value: 5,
					ChildNodes: []*common.Node{
						{
							Value: 4,
							ChildNodes: []*common.Node{
								{
									Value: 1,
									ChildNodes: []*common.Node{
										{
											Value:      2,
											ChildNodes: []*common.Node{},
										},
										{
											Value:      9,
											ChildNodes: []*common.Node{},
										},
									},
								},
								{
									Value:      -6,
									ChildNodes: []*common.Node{},
								},
							},
						},
						{
							Value: 3,
							ChildNodes: []*common.Node{
								{
									Value:      0,
									ChildNodes: []*common.Node{},
								},
								{
									Value: 7,
									ChildNodes: []*common.Node{
										{
											Value:      8,
											ChildNodes: []*common.Node{},
										},
									},
								},
								{
									Value:      -4,
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
