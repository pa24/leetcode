package SearchInsertPosition

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		wont   int
	}{{
		name:   "Target 5",
		nums:   []int{1, 3, 5, 6},
		target: 5,
		wont:   2,
	}, {
		name:   "Target 2",
		nums:   []int{1, 3, 5, 6},
		target: 2,
		wont:   1,
	}, {
		name:   "Target 7",
		nums:   []int{1, 3, 5, 6},
		target: 7,
		wont:   4,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchInsert(tt.nums, tt.target)
			if result != tt.wont {
				t.Errorf("maxProfit(%v)= %d, exepted %d", tt.nums, result, tt.wont)
			}
		})
	}
}
