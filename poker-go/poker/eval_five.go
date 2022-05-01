package poker

import "fmt"

var FLUSH_FIVE_RANK = make([]int, MAX_FLUSH_FIVE_KEY+1)
var FACE_FIVE_RANK = make([]int, MAX_FACE_FIVE_KEY+1)
var HAND_FACES [][5]int
var HAND_TYPE []string
var NB_HAND_FIVE_RANK int

func BuildEvalFiveTables(verbose bool) {

	if verbose {
		fmt.Println(" ")
		defer track(runningtime("BuildEvalFiveTables"))
	}

	var c1, c2, c3, c4, c5 int
	var handFaceKey uint32
	rank := 0

	// High Card
	for c1 = 4; c1 < NB_FACE; c1++ {
		for c2 = 0; c2 < c1; c2++ {
			for c3 = 0; c3 < c2; c3++ {
				for c4 = 0; c4 < c3; c4++ {
					for c5 = 0; c5 < c4; c5++ {
						// No straights, including A2345
						if !((c1-c5 == 4) || (c1 == 12 && c2 == 3)) {
							handFaceKey = FACE_FIVE_KEY[c1] + FACE_FIVE_KEY[c2] + FACE_FIVE_KEY[c3] + FACE_FIVE_KEY[c4] + FACE_FIVE_KEY[c5]
							FACE_FIVE_RANK[handFaceKey] = rank
							HAND_FACES = append(HAND_FACES, [5]int{c1, c2, c3, c4, c5})
							HAND_TYPE = append(HAND_TYPE, "high-card")
							rank += 1
						}
					}
				}
			}
		}
	}

	// One Pair
	for c1 = 0; c1 < NB_FACE; c1++ {
		for c2 = 0; c2 < NB_FACE; c2++ {
			for c3 = 0; c3 < c2; c3++ {
				for c4 = 0; c4 < c3; c4++ {
					// No Three of a Kind
					if !((c1 == c2) || (c1 == c3) || (c1 == c4)) {
						handFaceKey = 2*FACE_FIVE_KEY[c1] + FACE_FIVE_KEY[c2] + FACE_FIVE_KEY[c3] + FACE_FIVE_KEY[c4]
						FACE_FIVE_RANK[handFaceKey] = rank
						HAND_FACES = append(HAND_FACES, [5]int{c1, c1, c2, c3, c4})
						HAND_TYPE = append(HAND_TYPE, "one-pair")
						rank += 1
					}
				}
			}
		}
	}

	// Two Pair
	for c1 = 0; c1 < NB_FACE; c1++ {
		for c2 = 0; c2 < c1; c2++ {
			for c3 = 0; c3 < NB_FACE; c3++ {
				// No Three of a Kind
				if !((c1 == c3) || (c2 == c3)) {
					handFaceKey = 2*FACE_FIVE_KEY[c1] + 2*FACE_FIVE_KEY[c2] + FACE_FIVE_KEY[c3]
					FACE_FIVE_RANK[handFaceKey] = rank
					HAND_FACES = append(HAND_FACES, [5]int{c1, c1, c2, c2, c3})
					HAND_TYPE = append(HAND_TYPE, "two-pairs")
					rank += 1
				}
			}
		}
	}

	// Three of a kind
	for c1 = 0; c1 < NB_FACE; c1++ {
		for c2 = 0; c2 < NB_FACE; c2++ {
			for c3 = 0; c3 < c2; c3++ {
				// No Four of a Kind
				if !((c1 == c2) || (c1 == c3)) {
					handFaceKey = 3*FACE_FIVE_KEY[c1] + FACE_FIVE_KEY[c2] + FACE_FIVE_KEY[c3]
					FACE_FIVE_RANK[handFaceKey] = rank
					HAND_FACES = append(HAND_FACES, [5]int{c1, c1, c1, c2, c3})
					HAND_TYPE = append(HAND_TYPE, "three-of-a-kind")

					rank += 1
				}
			}
		}
	}

	// Low Straight
	c1 = 3
	c5 = 12
	handFaceKey = FACE_FIVE_KEY[c1] + FACE_FIVE_KEY[c1-1] + FACE_FIVE_KEY[c1-2] + FACE_FIVE_KEY[c1-3] + FACE_FIVE_KEY[c5]
	FACE_FIVE_RANK[handFaceKey] = rank
	HAND_FACES = append(HAND_FACES, [5]int{c1, c1 - 1, c1 - 2, c1 - 3, c5})
	HAND_TYPE = append(HAND_TYPE, "straight")
	rank += 1

	// Other Straight
	for c1 = 4; c1 < NB_FACE; c1++ {
		handFaceKey = FACE_FIVE_KEY[c1] + FACE_FIVE_KEY[c1-1] + FACE_FIVE_KEY[c1-2] + FACE_FIVE_KEY[c1-3] + FACE_FIVE_KEY[c1-4]
		FACE_FIVE_RANK[handFaceKey] = rank
		HAND_FACES = append(HAND_FACES, [5]int{c1, c1 - 1, c1 - 2, c1 - 3, c1 - 4})
		HAND_TYPE = append(HAND_TYPE, "straight")
		rank += 1
	}

	// Flush
	for c1 = 4; c1 < NB_FACE; c1++ {
		for c2 = 0; c2 < c1; c2++ {
			for c3 = 0; c3 < c2; c3++ {
				for c4 = 0; c4 < c3; c4++ {
					for c5 = 0; c5 < c4; c5++ {
						// No straights, including A2345
						if !((c1-c5 == 4) || (c1 == 12 && c2 == 3)) {
							handFaceKey = FLUSH_FIVE_KEY[c1] + FLUSH_FIVE_KEY[c2] + FLUSH_FIVE_KEY[c3] + FLUSH_FIVE_KEY[c4] + FLUSH_FIVE_KEY[c5]
							FLUSH_FIVE_RANK[handFaceKey] = rank
							HAND_FACES = append(HAND_FACES, [5]int{c1, c2, c3, c4, c5})
							HAND_TYPE = append(HAND_TYPE, "flush")
							rank += 1
						}

					}
				}
			}
		}
	}

	// Full House
	for c1 = 0; c1 < NB_FACE; c1++ {
		for c2 = 0; c2 < NB_FACE; c2++ {
			// No Four of a Kind
			if !(c1 == c2) {
				handFaceKey = 3*FACE_FIVE_KEY[c1] + 2*FACE_FIVE_KEY[c2]
				FACE_FIVE_RANK[handFaceKey] = rank
				HAND_FACES = append(HAND_FACES, [5]int{c1, c1, c1, c2, c2})
				HAND_TYPE = append(HAND_TYPE, "full-house")
				rank += 1
			}
		}
	}

	// Four of a Kind
	for c1 = 0; c1 < NB_FACE; c1++ {
		for c2 = 0; c2 < NB_FACE; c2++ {
			// No 'Five of a Kind'
			if !(c1 == c2) {
				handFaceKey = 4*FACE_FIVE_KEY[c1] + 1*FACE_FIVE_KEY[c2]
				FACE_FIVE_RANK[handFaceKey] = rank
				HAND_FACES = append(HAND_FACES, [5]int{c1, c1, c1, c1, c2})
				HAND_TYPE = append(HAND_TYPE, "four-of-a-kind")
				rank += 1
			}
		}
	}

	// Low Straight Flush
	c1 = 3
	c5 = 12
	handFaceKey = FLUSH_FIVE_KEY[c1] + FLUSH_FIVE_KEY[c1-1] + FLUSH_FIVE_KEY[c1-2] + FLUSH_FIVE_KEY[c1-3] + FLUSH_FIVE_KEY[c5]
	FLUSH_FIVE_RANK[handFaceKey] = rank
	HAND_FACES = append(HAND_FACES, [5]int{c1, c1 - 1, c1 - 2, c1 - 3, c5})
	HAND_TYPE = append(HAND_TYPE, "straight-flush")
	rank += 1

	// Other Straight Flush
	for c1 = 4; c1 < NB_FACE; c1++ {
		handFaceKey = FLUSH_FIVE_KEY[c1] + FLUSH_FIVE_KEY[c1-1] + FLUSH_FIVE_KEY[c1-2] + FLUSH_FIVE_KEY[c1-3] + FLUSH_FIVE_KEY[c1-4]
		FLUSH_FIVE_RANK[handFaceKey] = rank
		HAND_FACES = append(HAND_FACES, [5]int{c1, c1 - 1, c1 - 2, c1 - 3, c1 - 4})
		HAND_TYPE = append(HAND_TYPE, "straight-flush")
		rank += 1
	}

	NB_HAND_FIVE_RANK = rank
	if verbose {
		fmt.Printf("NB_HAND_FIVE_RANK = %d\n", NB_HAND_FIVE_RANK)
	}

}

func GetRankFive(c [5]int) int {
	// input = array of 5 cards all distinct integers from 0 to NB_FACE*NB_SUIT
	// in order defined by CARD_NO

	var handFaceKey uint32
	var handRank int

	if CARD_SUIT[c[0]] == CARD_SUIT[c[1]] &&
		CARD_SUIT[c[0]] == CARD_SUIT[c[2]] &&
		CARD_SUIT[c[0]] == CARD_SUIT[c[3]] &&
		CARD_SUIT[c[0]] == CARD_SUIT[c[4]] {
		handFaceKey = FLUSH_FIVE_KEY[CARD_FACE[c[0]]] +
			FLUSH_FIVE_KEY[CARD_FACE[c[1]]] +
			FLUSH_FIVE_KEY[CARD_FACE[c[2]]] +
			FLUSH_FIVE_KEY[CARD_FACE[c[3]]] +
			FLUSH_FIVE_KEY[CARD_FACE[c[4]]]
		handRank = FLUSH_FIVE_RANK[handFaceKey]
	} else {
		handFaceKey = FACE_FIVE_KEY[CARD_FACE[c[0]]] +
			FACE_FIVE_KEY[CARD_FACE[c[1]]] +
			FACE_FIVE_KEY[CARD_FACE[c[2]]] +
			FACE_FIVE_KEY[CARD_FACE[c[3]]] +
			FACE_FIVE_KEY[CARD_FACE[c[4]]]
		handRank = FACE_FIVE_RANK[handFaceKey]
	}
	return handRank
}
