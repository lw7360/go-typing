package main

import ()

type Stats struct {
	words    int
	seconds  int
	errors   int
	wordlist string
}

func (s *Stats) loadStats(filename string) bool {
	// TODO: Implement Later
	s.words = 10
	s.seconds = 11
	s.errors = 12
	s.wordlist = "words.txt"

	return true
}

func (s *Stats) saveStats(filename string) bool {
	// TODO: Implement Later

	return true
}

func (s *Stats) wpm() float64 {
	return float64(s.words / s.seconds / 60)
}
