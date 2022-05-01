package poker

import (
	"encoding/json"
	"fmt"
)

type HandTypeStatsStruct struct {
	NbHand  int
	MinRank int
	MaxRank int
	NbOccur int
}
type handStatsStruct map[string]HandTypeStatsStruct

var FiveHandTypeStatsTarget = map[string]HandTypeStatsStruct{
	"high-card":       {NbHand: 1277, MinRank: 0, MaxRank: 1276, NbOccur: 1302540},
	"one-pair":        {NbHand: 2860, MinRank: 1277, MaxRank: 4136, NbOccur: 1098240},
	"two-pairs":       {NbHand: 858, MinRank: 4137, MaxRank: 4994, NbOccur: 123552},
	"three-of-a-kind": {NbHand: 858, MinRank: 4995, MaxRank: 5852, NbOccur: 54912},
	"straight":        {NbHand: 10, MinRank: 5853, MaxRank: 5862, NbOccur: 10200},
	"flush":           {NbHand: 1277, MinRank: 5863, MaxRank: 7139, NbOccur: 5108},
	"full-house":      {NbHand: 156, MinRank: 7140, MaxRank: 7295, NbOccur: 3744},
	"four-of-a-kind":  {NbHand: 156, MinRank: 7296, MaxRank: 7451, NbOccur: 624},
	"straight-flush":  {NbHand: 10, MinRank: 7452, MaxRank: 7461, NbOccur: 40},
}

var SevenHandTypeStatsTarget = map[string]HandTypeStatsStruct{
	"high-card":       {NbHand: 407, MinRank: 48, MaxRank: 1276, NbOccur: 23294460},
	"one-pair":        {NbHand: 1470, MinRank: 1295, MaxRank: 4136, NbOccur: 58627800},
	"two-pairs":       {NbHand: 763, MinRank: 4140, MaxRank: 4994, NbOccur: 31433400},
	"three-of-a-kind": {NbHand: 575, MinRank: 5003, MaxRank: 5852, NbOccur: 6461620},
	"straight":        {NbHand: 10, MinRank: 5853, MaxRank: 5862, NbOccur: 6180020},
	"flush":           {NbHand: 1277, MinRank: 5863, MaxRank: 7139, NbOccur: 4047644},
	"full-house":      {NbHand: 156, MinRank: 7140, MaxRank: 7295, NbOccur: 3473184},
	"four-of-a-kind":  {NbHand: 156, MinRank: 7296, MaxRank: 7451, NbOccur: 224848},
	"straight-flush":  {NbHand: 10, MinRank: 7452, MaxRank: 7461, NbOccur: 41584},
}

func BuildFiveHandStats(verbose bool) map[string]HandTypeStatsStruct {

	if verbose {
		fmt.Println(" ")
		defer track(runningtime("BuildFiveHandStats"))
	}

	var FiveHandTypeStats = make(map[string]HandTypeStatsStruct)

	stats := make(map[string]*HandTypeStatsStruct)
	var rankCount = make(map[int]int)

	var c1, c2, c3, c4, c5 int
	var cards [5]int
	var rank int

	for c1 = 0; c1 < DECK_SIZE; c1++ {
		for c2 = 0; c2 < c1; c2++ {
			for c3 = 0; c3 < c2; c3++ {
				for c4 = 0; c4 < c3; c4++ {
					for c5 = 0; c5 < c4; c5++ {
						cards = [5]int{c1, c2, c3, c4, c5}
						rank = GetRankFive(cards)
						rankCount[rank] += 1
					}
				}
			}
		}
	}

	for rank, nbOccur := range rankCount {
		handType := HAND_TYPE[rank]
		_, present := stats[handType]
		if !present {
			stats[handType] = &HandTypeStatsStruct{NbHand: 0, MinRank: rank, MaxRank: rank, NbOccur: 0}
		}
		obj := stats[handType]
		obj.NbHand += 1
		obj.NbOccur += nbOccur
		obj.MinRank = min(obj.MinRank, rank)
		obj.MaxRank = max(obj.MaxRank, rank)
	}

	for k, v := range stats {
		FiveHandTypeStats[k] = *v
	}

	if verbose {
		fmt.Printf("stats five cards\n")
		for k, v := range stats {
			jsonObj, _ := json.Marshal(*v)
			fmt.Printf("\thand-type=%16s\tstats=%s\n", k, jsonObj)
		}
	}

	return FiveHandTypeStats
}

func BuildSevenHandStats(verbose bool) map[string]HandTypeStatsStruct {

	if verbose {
		fmt.Println(" ")
		defer track(runningtime("BuildSevenHandStats"))
	}

	var SevenHandTypeStats = make(map[string]HandTypeStatsStruct)

	stats := make(map[string]*HandTypeStatsStruct)

	var rankCount = make(map[int]int)

	var c1, c2, c3, c4, c5, c6, c7 int
	var cards [7]int
	var rank int

	for c1 = 0; c1 < DECK_SIZE; c1++ {
		for c2 = 0; c2 < c1; c2++ {
			for c3 = 0; c3 < c2; c3++ {
				for c4 = 0; c4 < c3; c4++ {
					for c5 = 0; c5 < c4; c5++ {
						for c6 = 0; c6 < c5; c6++ {
							for c7 = 0; c7 < c6; c7++ {
								cards = [7]int{c1, c2, c3, c4, c5, c6, c7}
								rank = GetRank(cards)
								rankCount[rank] += 1
							}
						}
					}
				}
			}
		}
	}

	for rank, NbOccur := range rankCount {
		handType := HAND_TYPE[rank]
		_, present := stats[handType]
		if !present {
			stats[handType] = &HandTypeStatsStruct{NbHand: 0, MinRank: rank, MaxRank: rank, NbOccur: 0}
		}
		obj := stats[handType]
		obj.NbHand += 1
		obj.NbOccur += NbOccur
		obj.MinRank = min(obj.MinRank, rank)
		obj.MaxRank = max(obj.MaxRank, rank)
	}

	for k, v := range stats {
		SevenHandTypeStats[k] = *v
	}

	if verbose {
		fmt.Printf("stats seven cards\n")
		for k, v := range stats {
			jsonObj, _ := json.Marshal(*v)
			fmt.Printf("\thand-type=%16s\tstats=%s\n", k, jsonObj)
		}
	}

	return SevenHandTypeStats

}
