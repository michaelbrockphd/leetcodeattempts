package leet_code_attempts

import "testing"

func Test_297_StartingWithText(t *testing.T) {
	testCases := []struct {
		data string
	}{
		{"[1,2,3,null,null,4,5]"},
		{"[]"},
		{"[1]"},
		{"[1,2]"},
	}

	for _, testCase := range testCases {
		ser := Constructor()

		deser := Constructor()

		tmp := deser.deserialize(testCase.data)

		ans := ser.serialize(tmp)

		if ans != testCase.data {
			t.Errorf(
				"Search failed: expected %v but got %v",
				testCase.data,
				ans)
		}
	}
}
