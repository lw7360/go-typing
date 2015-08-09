package main

import (
	"encoding/json"
	"io/ioutil"
)

type Stats struct {
	Words   int
	Seconds float64
	Errors  int
	Wpm     int
}

func (s *Stats) loadStats(filename string) bool {
	statsFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return false
	}

	err = json.Unmarshal(statsFile, s)
	if err != nil {
		return false
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
	if int(s.Seconds) == 0 {
		return 0
	}
	return int(float64(s.Words) / s.Seconds * 60.0)
}

func (s *Stats) reset() {
	s.Words = 0
	s.Seconds = 0
	s.Errors = 0
	s.Wpm = 0
}
