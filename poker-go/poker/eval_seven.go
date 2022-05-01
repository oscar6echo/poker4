package poker

import "fmt"

var FACE_RANK = make([]int, MAX_FACE_SEVEN_KEY+1)
var FLUSH_RANK = make([]int, MAX_FLUSH_SEVEN_KEY+1)
var FLUSH_SUIT = make([]int, MAX_SUIT_KEY+1)
var NB_HAND_SEVEN_RANK int

func Setup(verbose bool) {
	InitKeys(verbose)
	BuildEvalFiveTables(verbose)
	BuildEvalSevenTables(verbose)
}

func BuildEvalSevenTables(verbose bool) {

	if verbose {
		fmt.Println(" ")
		defer track(runningtime("BuildEvalSevenTables"))
	}

	var c1, c2, c3, c4, c5, c6, c7 int
	var handFaceKey uint32
	var handSuitKey uint32
	var suitCount int

	// FACE_RANK
	for c1 = 0; c1 < NB_FACE; c1++ {
		for c2 = 0; c2 < c1+1; c2++ {
			for c3 = 0; c3 < c2+1; c3++ {
				for c4 = 0; c4 < c3+1; c4++ {
					for c5 = 0; c5 < c4+1; c5++ {
						for c6 = 0; c6 < c5+1; c6++ {
							for c7 = 0; c7 < c6+1; c7++ {
								// No 5, or more, same faces
								if (c1-c5 > 0) && (c2-c6 > 0) && (c3-c7 > 0) {
									handFaceKey = FACE_SEVEN_KEY[c1] + FACE_SEVEN_KEY[c2] + FACE_SEVEN_KEY[c3] + FACE_SEVEN_KEY[c4] + FACE_SEVEN_KEY[c5] + FACE_SEVEN_KEY[c6] + FACE_SEVEN_KEY[c7]
									// arbitrary valid suits (4*0+,3*1)
									FACE_RANK[handFaceKey] = GetRankSeven([7]int{4 * c1, 4 * c2, 4 * c3, 4 * c4, 4*c5 + 1, 4*c6 + 1, 4*c7 + 1})
								}
							}
						}
					}
				}
			}
		}
	}

	// FLUSH_RANK 7 cards
	for c1 = 6; c1 < NB_FACE; c1++ {
		for c2 = 0; c2 < c1; c2++ {
			for c3 = 0; c3 < c2; c3++ {
				for c4 = 0; c4 < c3; c4++ {
					for c5 = 0; c5 < c4; c5++ {
						for c6 = 0; c6 < c5; c6++ {
							for c7 = 0; c7 < c6; c7++ {
								handFaceKey = FLUSH_SEVEN_KEY[c1] + FLUSH_SEVEN_KEY[c2] + FLUSH_SEVEN_KEY[c3] + FLUSH_SEVEN_KEY[c4] + FLUSH_SEVEN_KEY[c5] + FLUSH_SEVEN_KEY[c6] + FLUSH_SEVEN_KEY[c7]
								// arbitrary suit (7*0)
								FLUSH_RANK[handFaceKey] = GetRankSeven([7]int{4 * c1, 4 * c2, 4 * c3, 4 * c4, 4 * c5, 4 * c6, 4 * c7})
							}
						}
					}
				}
			}
		}
	}
	// FLUSH_RANK 6 cards
	for c1 = 5; c1 < NB_FACE; c1++ {
		for c2 = 0; c2 < c1; c2++ {
			for c3 = 0; c3 < c2; c3++ {
				for c4 = 0; c4 < c3; c4++ {
					for c5 = 0; c5 < c4; c5++ {
						for c6 = 0; c6 < c5; c6++ {
							handFaceKey = FLUSH_SEVEN_KEY[c1] + FLUSH_SEVEN_KEY[c2] + FLUSH_SEVEN_KEY[c3] + FLUSH_SEVEN_KEY[c4] + FLUSH_SEVEN_KEY[c5] + FLUSH_SEVEN_KEY[c6]
							// arbitrary suit (6*0, 1*1)
							FLUSH_RANK[handFaceKey] = GetRankSeven([7]int{4 * c1, 4 * c2, 4 * c3, 4 * c4, 4 * c5, 4 * c6, 1})
						}
					}
				}
			}
		}
	}
	// FLUSH_RANK 5 cards
	for c1 = 4; c1 < NB_FACE; c1++ {
		for c2 = 0; c2 < c1; c2++ {
			for c3 = 0; c3 < c2; c3++ {
				for c4 = 0; c4 < c3; c4++ {
					for c5 = 0; c5 < c4; c5++ {
						handFaceKey = FLUSH_SEVEN_KEY[c1] + FLUSH_SEVEN_KEY[c2] + FLUSH_SEVEN_KEY[c3] + FLUSH_SEVEN_KEY[c4] + FLUSH_SEVEN_KEY[c5]
						// arbitrary suit (5*0)
						FLUSH_RANK[handFaceKey] = GetRankFive([5]int{4 * c1, 4 * c2, 4 * c3, 4 * c4, 4 * c5})
					}
				}
			}
		}
	}
	// FLUSH_SUIT
	for c1 = 0; c1 < NB_SUIT; c1++ {
		for c2 = 0; c2 < c1+1; c2++ {
			for c3 = 0; c3 < c2+1; c3++ {
				for c4 = 0; c4 < c3+1; c4++ {
					for c5 = 0; c5 < c4+1; c5++ {
						for c6 = 0; c6 < c5+1; c6++ {
							for c7 = 0; c7 < c6+1; c7++ {
								handSuitKey = SUIT_KEY[c1] + SUIT_KEY[c2] + SUIT_KEY[c3] + SUIT_KEY[c4] + SUIT_KEY[c5] + SUIT_KEY[c6] + SUIT_KEY[c7]
								FLUSH_SUIT[handSuitKey] = -1
								for suit := 0; suit < NB_SUIT; suit++ {
									suitCount = 0
									if c1 == suit {
										suitCount++
									}
									if c2 == suit {
										suitCount++
									}
									if c3 == suit {
										suitCount++
									}
									if c4 == suit {
										suitCount++
									}
									if c5 == suit {
										suitCount++
									}
									if c6 == suit {
										suitCount++
									}
									if c7 == suit {
										suitCount++
									}
									if suitCount >= 5 {
										FLUSH_SUIT[handSuitKey] = suit
									}
								}
							}
						}
					}
				}
			}
		}
	}

}

func GetRankSeven(c [7]int) int {
	// input = array of 5 cards all distinct integers from 0 to NB_FACE*NB_SUIT
	// in order defined by CARD_NO

	bestHandRank := -1
	var handRank int
	var arr [5]int

	for c1 := 0; c1 < 7; c1++ {
		for c2 := 0; c2 < c1; c2++ {
			k := 0
			for i := 0; i < 7; i++ {
				// exclude cards c1 and c2
				if !(i == c1) && !(i == c2) {
					arr[k] = c[i]
					k += 1
				}
			}
			handRank = GetRankFive(arr)
			if handRank > bestHandRank {
				bestHandRank = handRank
			}
		}
	}
	return bestHandRank
}

func GetRank(c [7]int) int {
	// input = array of 5 cards all distinct integers from 0 to NB_FACE*NB_SUIT
	// in order defined by CARD_NO

	var handKey uint32
	var handFaceKey uint32
	var handFlushKey uint32
	var handSuitKey uint32
	var handSuit int
	var handRank int

	handKey = CARD_FACE_KEY[c[0]] +
		CARD_FACE_KEY[c[1]] +
		CARD_FACE_KEY[c[2]] +
		CARD_FACE_KEY[c[3]] +
		CARD_FACE_KEY[c[4]] +
		CARD_FACE_KEY[c[5]] +
		CARD_FACE_KEY[c[6]]

	handSuitKey = handKey & SUIT_MASK
	handSuit = FLUSH_SUIT[handSuitKey]
	if handSuit == -1 {
		handFaceKey = handKey >> SUIT_BIT_SHIFT
		handRank = FACE_RANK[handFaceKey]
	} else {
		handFlushKey = 0
		for i := 0; i < 7; i++ {
			if handSuit == CARD_SUIT[c[i]] {
				handFlushKey += CARD_FLUSH_KEY[c[i]]
			}
		}
		handRank = FLUSH_RANK[handFlushKey]
	}

	return handRank
}
