package main

import (
	"testing"
)

func TestChicken_many(t *testing.T) {
	testCases := []struct {
		roof     int
		chickens []int
		expected int
	}{
		{roof: 5, chickens: []int{2, 5, 10, 12, 15}, expected: 2},
		{roof: 4, chickens: []int{1, 3, 5, 7, 9, 999_999_998, 999_999_999, 1_000_000_000}, expected: 3},
		{roof: 10, chickens: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expected: 10},
		{roof: 4, chickens: []int{2, 4, 6, 8, 10, 12}, expected: 2},
		{roof: 1, chickens: []int{1, 1, 1, 1, 1}, expected: 5},
		{roof: 0, chickens: []int{1, 2, 3}, expected: 0},
		{roof: 4, chickens: []int{5, 1, 4, 3, 2}, expected: 4},
		{roof: 2, chickens: []int{2, 4, 6, 8, 10, 12, 14, 16}, expected: 1},
	}

	for _, tc := range testCases {

		if result := protectedChickens(tc.roof, tc.chickens); result != tc.expected {
			t.Errorf("FAIL: roof=%d, chickens=%v, expected=%d, got=%d\n", tc.roof, tc.chickens, tc.expected, result)
		}
	}
}
