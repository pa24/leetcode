package besttime

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		wont   int
	}{
		{
			name:   "Regular case",
			prices: []int{7, 6, 3, 1, 6},
			wont:   5,
		},
		{
			name:   "One day",
			prices: []int{5},
			wont:   0,
		},
		{
			name:   "Increasing",
			prices: []int{1, 2, 3, 4, 5, 6},
			wont:   5,
		},
		{
			name:   "Decrease",
			prices: []int{6, 5, 4, 3, 2, 1},
			wont:   0,
		},
		{
			name:   "Profit in middle",
			prices: []int{3, 8, 1, 10, 2},
			wont:   9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxProfit(tt.prices)
			if result != tt.wont {
				t.Errorf("maxProfit(%v)= %d, exepted %d", tt.prices, result, tt.wont)
			}
		})
	}
}
