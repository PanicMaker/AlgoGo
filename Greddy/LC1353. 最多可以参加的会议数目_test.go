package greddy

import (
	"testing"
)

func TestMaxEvents(t *testing.T) {
	tests := []struct {
		name   string
		events [][]int
		want   int
	}{
		{
			name:   "Example 1",
			events: [][]int{{1, 2}, {2, 3}, {3, 4}},
			want:   3,
		},
		{
			name:   "Example 2",
			events: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}},
			want:   4,
		},
		{
			name:   "Example 3",
			events: [][]int{{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1}},
			want:   4,
		},
		{
			name:   "No events",
			events: [][]int{},
			want:   0,
		},
		{
			name:   "Single event",
			events: [][]int{{1, 1}},
			want:   1,
		},
		{
			name:   "Overlapping events",
			events: [][]int{{1, 5}, {1, 5}, {1, 5}, {2, 3}, {2, 3}},
			want:   5,
		},
		{
			name:   "Complex case",
			events: [][]int{{1, 10}, {2, 3}, {4, 5}, {6, 7}, {8, 9}},
			want:   5,
		},
		{
			name:   "Events with same start day",
			events: [][]int{{1, 2}, {1, 2}, {1, 2}},
			want:   2,
		},
		{
			name:   "Events with same end day",
			events: [][]int{{1, 5}, {2, 5}, {3, 5}},
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEvents(tt.events); got != tt.want {
				t.Errorf("maxEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}
