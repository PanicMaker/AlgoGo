package DP

import "testing"

func Test_nthUglyNumber(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{10, 12},
		{15, 24},
		{100, 1536},
	}
	for _, tt := range tests {
		got := nthUglyNumberII(tt.n)
		if got != tt.want {
			t.Errorf("nthUglyNumber(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
