package leet_code_attempts

import "testing"

// Constraints
// 1) 0 <= digits.length <= 4
// 2) digits[i] is a digit in the range ['2', '9'].

func Test_017_letterCombinations(t *testing.T) {
	testCases := []struct {
		digits   string
		expected []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", []string{}},
		{"2", []string{"a", "b", "c"}},
	}

	for _, testCase := range testCases {
		result := letterCombinations(testCase.digits)

		lenR := len(result)
		lenE := len(testCase.expected)

		if lenR != lenE {
			t.Errorf("Fail: expected %v elements, but got %v", lenR, lenE)

		} else {
			for i, v := range result {
				if v != testCase.expected[i] {
					t.Errorf(
						"Unexpected value: expected %v but got %v",
						v,
						testCase.expected[i])
				}
			}
		}
	}
}
