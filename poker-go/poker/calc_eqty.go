package poker

import (
	"fmt"
	"log"
	"math/rand"
)

type PlayerCards [][2]int
type TableCards []int

type handEquity struct {
	Win float32
	Tie float32
}

func TryCalcEquity() {
	defer track(runningtime("TryCalcEquity"))

	// var playerCards = PlayerCards{{2, 18}, {5, 22}, {6, 34}}
	// var tableCards = TableCards{41, 8, 30}

	var playerCards = PlayerCards{{7, 29}, {4, 11}}
	var tableCards = TableCards{}

	var eqty = CalcEquity(playerCards, tableCards)
	fmt.Println(eqty)

}

func TryCalcEquityMonteCarlo() {
	defer track(runningtime("TryCalcEquityMonteCarlo"))

	var playerCards = [2]int{2, 18}
	var tableCards = []int{41, 8, 30}
	nbPlayer := 4
	nbGame := int(1e6)

	var eqty = CalcEquityMonteCarlo(playerCards, tableCards, nbPlayer, nbGame)
	fmt.Println(eqty)

}

func CalcEquity(playerCards [][2]int, tableCards []int) []handEquity {

	T := len(tableCards)
	if T != 0 && T != 3 && T != 4 && T != 5 {
		log.Fatal("len(tableCards) must be 0, 3, 4, 5")
	}

	var deckCards []int = buildDeckCards(playerCards, tableCards)
	D := len(deckCards)

	P := len(playerCards)
	var eqty = make([]*handEquity, P)
	var equity = make([]handEquity, P)
	for i := range playerCards {
		eqty[i] = &handEquity{Win: 0, Tie: 0}
	}

	var rank = make([]int, len(playerCards))
	var c1, c2, c3, c4, c5, p, nbGame int
	var cards [7]int

	nbGame = 0

	// zero table cards
	if T == 0 {
		for c1 = 0; c1 < D-2*P-T; c1++ {
			for c2 = 0; c2 < c1; c2++ {
				for c3 = 0; c3 < c2; c3++ {
					for c4 = 0; c4 < c3; c4++ {
						for c5 = 0; c5 < c4; c5++ {
							for p = 0; p < P; p++ {
								cards = [7]int{
									playerCards[p][0],
									playerCards[p][1],
									deckCards[c1],
									deckCards[c2],
									deckCards[c3],
									deckCards[c4],
									deckCards[c5],
								}
								rank[p] = GetRank(cards)
							}
							updateEquity(playerCards, tableCards, rank, eqty)
							nbGame++
						}
					}
				}
			}
		}
	}

	// 3 table cards
	if T == 3 {
		for c1 = 0; c1 < D-2*P-T; c1++ {
			for c2 = 0; c2 < c1; c2++ {
				for p = 0; p < P; p++ {
					cards = [7]int{
						playerCards[p][0],
						playerCards[p][1],
						tableCards[0],
						tableCards[1],
						tableCards[2],
						deckCards[c1],
						deckCards[c2],
					}
					rank[p] = GetRank(cards)
				}
				updateEquity(playerCards, tableCards, rank, eqty)
				nbGame++
			}
		}
	}

	// 4 table cards
	if T == 4 {
		for c1 = 0; c1 < D-2*P-T; c1++ {
			for p = 0; p < P; p++ {
				cards = [7]int{
					playerCards[p][0],
					playerCards[p][1],
					tableCards[0],
					tableCards[1],
					tableCards[2],
					tableCards[3],
					deckCards[c1],
				}
				rank[p] = GetRank(cards)
			}
			updateEquity(playerCards, tableCards, rank, eqty)
			nbGame++
		}
	}

	// 5 table cards
	if T == 5 {

		for p = 0; p < P; p++ {
			cards = [7]int{
				playerCards[p][0],
				playerCards[p][1],
				tableCards[0],
				tableCards[1],
				tableCards[2],
				tableCards[3],
				tableCards[4],
			}
			rank[p] = GetRank(cards)
		}
		updateEquity(playerCards, tableCards, rank, eqty)
		nbGame++
	}

	for k, v := range eqty {
		equity[k].Win = v.Win / float32(nbGame)
		equity[k].Tie = v.Tie / float32(nbGame)
	}
	return equity
}

func buildDeckCards(playerCards [][2]int, tableCards []int) []int {
	var usedCards []int
	var deckCards []int
	var isUsed bool

	for _, e := range playerCards {
		usedCards = append(usedCards, e[0])
		usedCards = append(usedCards, e[1])
	}

	for _, e := range tableCards {
		usedCards = append(usedCards, e)
	}

	for i := 0; i < DECK_SIZE; i++ {
		isUsed = false
		for _, u := range usedCards {
			if i == u {
				isUsed = true
				break
			}
		}
		if !isUsed {
			deckCards = append(deckCards, i)
		}
	}

	return deckCards
}

func updateEquity(playerCards PlayerCards, tableCards TableCards, rank []int, eqty []*handEquity) {
	var maxRank, nbMax, p int

	maxRank = rank[0]
	nbMax = 1
	P := len(playerCards)

	for p = 1; p < P; p++ {
		if rank[p] > maxRank {
			maxRank = rank[p]
			nbMax = 1
		} else if rank[p] == maxRank {
			nbMax += 1
		}
	}
	for p = 0; p < P; p++ {
		if rank[p] == maxRank {
			if nbMax == 1 {
				eqty[p].Win += 1
			} else {
				eqty[p].Tie += 1 / float32(nbMax)
			}
		}
	}
}

func CalcEquityMonteCarlo(playerCards [2]int, tableCards []int, nbPlayer int, nbGame int) handEquity {

	if nbPlayer < 1 || nbPlayer > 9 {
		log.Fatal("nbPlayer must be between 1 and 9")
	}

	T := len(tableCards)
	if T != 0 && T != 3 && T != 4 && T != 5 {
		log.Fatal("len(tableCards) must be 0, 3, 4, 5")
	}

	var eqty handEquity = handEquity{Win: 0, Tie: 0}

	var pCards = [1][2]int{{playerCards[0], playerCards[1]}}
	var deckCards []int = buildDeckCards(pCards[:], tableCards)
	var rndCards = make([]int, 2*(nbPlayer-1)+5-T)
	var rndTableCards = make([]int, 5-T)

	var p, r, t, g int
	var c1, c2, c3, c4, c5, c6, c7 int
	var cards [7]int
	var maxRank, nbMax int
	var rank = make([]int, nbPlayer)

	for g = 0; g < nbGame; g++ {
		drawRandomCards(rndCards, deckCards)

		for t = 0; t < 5-T; t++ {
			rndTableCards[t] = rndCards[t]
		}
		r = 5 - T

		for p = 0; p < nbPlayer; p++ {
			// fmt.Println(">", p, r)
			if p == 0 {
				c1 = playerCards[0]
				c2 = playerCards[1]
			} else {
				c1 = rndCards[r]
				r++
				c2 = rndCards[r]
				r++
			}
			if T > 0 {
				c3 = tableCards[0]
			} else {
				c3 = rndTableCards[4]
			}
			if T > 1 {
				c4 = tableCards[1]
			} else {
				c4 = rndTableCards[3]
			}
			if T > 2 {
				c5 = tableCards[2]
			} else {
				c5 = rndTableCards[2]
			}
			if T > 3 {
				c6 = tableCards[3]
			} else {
				c6 = rndTableCards[1]
			}
			if T > 4 {
				c7 = tableCards[4]
			} else {
				c7 = rndTableCards[0]
			}

			cards = [7]int{c1, c2, c3, c4, c5, c6, c7}
			rank[p] = GetRank(cards)
		}

		maxRank = rank[0]
		nbMax = 1
		for p = 1; p < nbPlayer; p++ {
			if rank[p] > maxRank {
				maxRank = rank[p]
				nbMax = 1
			} else if rank[p] == maxRank {
				nbMax += 1
			}
		}
		if rank[0] == maxRank {
			if nbMax == 1 {
				eqty.Win += 1
			} else {
				eqty.Tie += float32(1 / nbMax)
			}
		}

	}

	eqty.Win /= float32(nbGame)
	eqty.Tie /= float32(nbGame)

	return eqty
}

func drawRandomCards(rndCards []int, deckCards []int) {

	var r, i int
	var isUsed bool
	D := len(deckCards)
	R := len(rndCards)
	c := 0

	for c < R {
		r = rand.Intn(D)
		isUsed = false
		for i = 0; i < c; i++ {
			if rndCards[i] == deckCards[r] {
				isUsed = true
				break
			}
		}
		if !isUsed {
			rndCards[c] = deckCards[r]
			c++
		}
	}
}
