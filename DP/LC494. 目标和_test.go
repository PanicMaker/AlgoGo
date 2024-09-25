package DP

import "testing"

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{nums: []int{1, 1, 1, 1, 1}, target: 3}, 5},
		{"2", args{nums: []int{1}, target: 1}, 1},
		{"3", args{nums: []int{1}, target: 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays4(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTargetSumWays1() = %v, want %v", got, tt.want)
			}
		})
	}
}
