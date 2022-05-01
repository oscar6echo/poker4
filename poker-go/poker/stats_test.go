package poker

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBuildFiveHandStats(t *testing.T) {

	Setup(false)
	stats := BuildFiveHandStats(false)

	got := stats
	want := FiveHandTypeStatsTarget
	same := reflect.DeepEqual(got, want)

	if !same {
		jsonWant, _ := json.MarshalIndent(want, "", "  ")
		jsonGot, _ := json.MarshalIndent(got, "", "  ")
		t.Errorf("BuildFiveHandStats ERROR\ngot = %s\nwant = %s", jsonGot, jsonWant)
	} else {
		// t.Logf("BuildFiveHandStats OK")

	}
}
func TestBuildSevenHandStats(t *testing.T) {

	Setup(false)
	stats := BuildSevenHandStats(false)

	got := stats
	want := SevenHandTypeStatsTarget
	same := reflect.DeepEqual(got, want)

	if !same {
		jsonWant, _ := json.MarshalIndent(want, "", "  ")
		jsonGot, _ := json.MarshalIndent(got, "", "  ")
		t.Errorf("BuildSevenHandStats ERROR\ngot = %s\nwant = %s", jsonGot, jsonWant)
	} else {
		// t.Logf("BuildSevenHandStats OK")

	}
}
