package main

import (
	"encoding/json"
	"io/ioutil"
)

type Stats struct {
	Words    int
	Seconds  int
	Errors   int
	Wordlist string
}

func (s *Stats) loadStats(filename string) bool {
	// TODO: Implement Later
	s.Words = 10
	s.Seconds = 11
	s.Errors = 12
	s.Wordlist = "words.txt"

	return true
}

func (s *Stats) saveStats(filename string) bool {
	// TODO: Implement Later
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
