package server

import (
	"encoding/json"
	"log"
)

func CustomLog(endpoint string, payloadType string, i interface{}) {
	// s, _ := json.MarshalIndent(i, "", "\t")
	s, _ := json.Marshal(i)
	log.Printf("-- %s -- %s = %s", endpoint, payloadType, string(s))
}
