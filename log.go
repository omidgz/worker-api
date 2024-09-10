package main

import (
	"encoding/json"
	"time"
)

type log struct {
	At   time.Time `json:"at"`
	Text string    `json:"text"`
}

var theLog []log

func Log(text string) {
	theLog = append(theLog, log{At: time.Now(), Text: text})
}

func JsonLog() string {
	if j, err := json.Marshal(theLog); err == nil {
		return string(j)
	} else {
		return err.Error()
	}
}
