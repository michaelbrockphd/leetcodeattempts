package leet_code_attempts

import "testing"

// Constraints
// 1) m == matrix.length
// 1) n == matrix[i].length
// 1) 1 <= n, m <= 300
// 1) -109 <= matix[i][j] <= 109
// 1) All the integers in each row are sorted in ascending order.
// 1) All the integers in each column are sorted in ascending order.
// 1) -109 <= target <= 109

func Test_240_searchMatrix2(t *testing.T) {
	testCases := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{[][]int{
			{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30}}, 5, true},

		{[][]int{
			{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30}}, 20, false},

		{[][]int{
			{-5}}, -5, true},

		{[][]int{
			{5},
			{6}}, 6, true},

		{[][]int{
			{1, 4},
			{2, 5}}, 5, true},
	}

	for i, testCase := range testCases {
		result := searchMatrix2(testCase.matrix, testCase.target)

		if result != testCase.expected {
			t.Errorf(
				"Test Case %v Failed: expected %v but got %v",
				i+1,
				testCase.expected,
				result)
		}
	}
}
