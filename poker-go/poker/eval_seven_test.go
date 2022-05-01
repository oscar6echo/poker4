package poker

import "testing"

type CASE struct {
	cards [7]int
	rank  int
}
type CASES []CASE

var resultBench int

func buildCases() CASES {

	cases := CASES{
		{[7]int{50, 6, 0, 5, 38, 7, 17}, 5124},
		{[7]int{23, 16, 34, 26, 0, 10, 8}, 1766},
		{[7]int{14, 4, 0, 7, 20, 8, 47}, 1625},
		{[7]int{10, 32, 43, 3, 25, 8, 49}, 1925},
		{[7]int{1, 16, 49, 24, 43, 42, 33}, 3676},
		{[7]int{49, 17, 1, 26, 11, 34, 20}, 887},
		{[7]int{5, 4, 18, 31, 34, 48, 22}, 1689},
		{[7]int{13, 47, 1, 25, 38, 26, 51}, 2815},
		{[7]int{44, 2, 28, 1, 3, 18, 22}, 5046},
		{[7]int{49, 27, 33, 51, 22, 1, 30}, 4000},
	}

	return cases
}

func TestGetRankSeven(t *testing.T) {

	Setup(false)
	cases := buildCases()

	for _, c := range cases {
		got := GetRankSeven(c.cards)
		if got != c.rank {
			t.Errorf("GetRankSeven(%v) = %d, want %d", c.cards, got, c.rank)
		} else {
			// t.Logf("GetRankSeven(%v) = %d (%s)", c.cards, got, HAND_TYPE[got])
		}
	}
}

func TestGetRank(t *testing.T) {

	Setup(false)

	cases := buildCases()

	for _, c := range cases {
		got := GetRank(c.cards)
		if got != c.rank {
			t.Errorf("GetRank(%v) = %d, want %d", c.cards, got, c.rank)
		} else {
			// t.Logf("GetRank(%v) = %d (%s)", c.cards, got, HAND_TYPE[got])
		}
	}
}

func BenchmarkGetRank(b *testing.B) {

	Setup(false)
	cases := buildCases()
	var r int

	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			r = GetRank(c.cards)
		}
	}
	resultBench = r
}
