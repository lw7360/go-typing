package main

import (
	"io/ioutil"
	"strings"
	"time"
)

type Game struct {
	wordList  []string  // Slice of words to practice
	curStats  Stats     // Stats for current session
	stats     Stats     // Stats for current session + all past sessions
	curWord   int       // Index of current word in wordList
	curChar   int       // Index of current char of current word
	startTime time.Time // Start time of current session
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

func (g *Game) initTime() {
	g.startTime = time.Now()
}

func (g *Game) gameTime() float64 {
	return time.Since(g.startTime).Seconds()
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
