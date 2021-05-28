package closestpair

import "testing"

func TestClosestPairSummingTo(t *testing.T) {
	cases := []struct {
		numbers                                 []int32
		target                                  int32
		expectedFirstValue, expectedSecondValue int32
	}{
		{[]int32{1, 60, -10, 70, -80, 85}, 0, -80, 85},
		{[]int32{1, -21, -8, 29, -1, 85}, 9, -21, 29},
	}

	for _, c := range cases {
		gotFirstValue, gotSecondValue := ClosestPairSummingTo(c.numbers, c.target)
		if gotFirstValue != c.expectedFirstValue || gotSecondValue != c.expectedSecondValue {
			t.Errorf("Closest pair in %v summing to %v, want %v and %v, got %v and %v", c.numbers, c.expectedFirstValue, c.target, c.expectedSecondValue, gotFirstValue, gotSecondValue)
		}
	}
}
