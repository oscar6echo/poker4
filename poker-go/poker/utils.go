package poker

import (
	"log"
	"time"
)

func runningtime(s string) (string, time.Time) {
	log.Println("Start:	", s)
	return s, time.Now()
}

func track(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println("End:	", s, "took", endTime.Sub(startTime))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func uniqueStr(arr []string) []string {
	occured := map[string]bool{}
	result := []string{}
	for _, e := range arr {
		if occured[e] != true {
			occured[e] = true
			result = append(result, e)
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
