package poker

import "testing"

func TestGetRankFive(t *testing.T) {

	Setup(false)

	cases := []struct {
		cards [5]int
		rank  int
	}{
		{[5]int{21, 33, 24, 22, 39}, 2459},
		{[5]int{51, 38, 14, 36, 17}, 3431},
		{[5]int{45, 8, 48, 34, 5}, 1171},
		{[5]int{13, 37, 33, 20, 35}, 3106},
		{[5]int{31, 26, 50, 16, 49}, 3971},
		{[5]int{28, 24, 25, 29, 2}, 4434},
		{[5]int{41, 13, 28, 25, 16}, 310},
		{[5]int{20, 36, 7, 42, 43}, 3572},
		{[5]int{38, 42, 8, 22, 44}, 761},
		{[5]int{32, 3, 18, 5, 42}, 320},
	}

	for _, c := range cases {
		got := GetRankFive(c.cards)
		if got != c.rank {
			t.Errorf("GetRankFive(%v) = %d, want %d", c.cards, got, c.rank)
		} else {
			// t.Logf("GetRankFive(%v) = %d", c.cards, got)
		}
	}
}
