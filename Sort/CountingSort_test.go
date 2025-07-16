package Sort

import (
	"reflect"
	"testing"
)

func TestCountingSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{5}, []int{5}},
		{[]int{3, 1, 2, 2, 3}, []int{1, 2, 2, 3, 3}},
		{[]int{-2, -5, -3, -2, 0}, []int{-5, -3, -2, -2, 0}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{0, -1, 1, 0, -1, 1}, []int{-1, -1, 0, 0, 1, 1}},
	}

	for _, tt := range tests {
		result := CountingSort(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("CountingSort(%v) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}
