package nowcoder

import (
	"reflect"
	"testing"
)

func TestSolveBM41(t *testing.T) {
	type args struct {
		preOrder []int
		inOrder  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				preOrder: []int{1, 2, 5, 3, 4},
				inOrder:  []int{2, 5, 1, 3, 4},
			},
			want: []int{1, 3, 4},
		},
		{
			name: "empty tree",
			args: args{
				preOrder: []int{},
				inOrder:  []int{},
			},
			want: []int{},
		},
		{
			name: "only left children",
			args: args{
				preOrder: []int{1, 2, 3, 4},
				inOrder:  []int{4, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "only right children",
			args: args{
				preOrder: []int{1, 2, 3, 4},
				inOrder:  []int{1, 2, 3, 4},
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveBM41(tt.args.preOrder, tt.args.inOrder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveBM41() = %v, want %v", got, tt.want)
			}
		})
	}
}
