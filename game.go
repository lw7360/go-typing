package main

import (
	"io/ioutil"
	"strings"
)

type Game struct {
	wordList []string
	stats    Stats
}

func (g *Game) loadWords(filename string) bool {
	words, err := ioutil.ReadFile(filename)
	if err != nil {
		return false
	}
	g.wordList = strings.Split(string(words), "\n")
	return true
}

func (g *Game) loadStats(filename string) bool {
	return g.stats.loadStats(filename)
}

func NewGame(wordsFile string, statsFile string) *Game {
	g := new(Game)
	if !g.loadWords(wordsFile) {
		panic("Failed to load wordsFile: \"" + wordsFile + "\"")
	}
	if !g.loadStats(statsFile) {
		panic("Failed to load statsFile: \"" + statsFile + "\"")
	}

	return g
}
