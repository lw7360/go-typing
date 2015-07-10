package main

import (
	"encoding/json"
	"io/ioutil"
)

type Stats struct {
	Words   int
	Seconds int
	Errors  int
}

func (s *Stats) loadStats(filename string) bool {
	statsFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(statsFile, s)
	if err != nil {
		panic(err)
	}

	return true
}

func (s *Stats) saveStats(filename string) bool {
	statsJson, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(filename, statsJson, 0644)
	return true
}

func (s *Stats) wpm() int {
	if s.Seconds == 0 {
		return 0
	}
	return s.Words / s.Seconds / 60
}
