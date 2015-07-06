package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func shuffle(a []string) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

type Game struct {
	wordList  string           // String containing all words
	curInd    int              // Index of current char
	errMap    map[int]struct{} // Map of indexes with errors. Empty struct as value saves space!
	curStats  Stats            // Stats for current session
	stats     Stats            // Stats for current session + all past sessions
	startTime time.Time        // Start time of current session
}

func (g *Game) loadWords(filename string) bool {
	rand.Seed(time.Now().UnixNano())

	words, err := ioutil.ReadFile(filename)
	if err != nil {
		return false
	}

	wordSlice := strings.Split(string(words), "\n")
	shuffle(wordSlice)

	g.wordList = strings.Join(wordSlice, " ")

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

func (g *Game) getRune(index int) rune {
	return rune(g.wordList[index])
}

func NewGame(wordsFile string, statsFile string) *Game {
	g := new(Game)
	if !g.loadWords(wordsFile) {
		panic("Failed to load wordsFile: \"" + wordsFile + "\"")
	}
	if !g.loadStats(statsFile) {
		panic("Failed to load statsFile: \"" + statsFile + "\"")
	}
	g.errMap = make(map[int]struct{})
	return g
}
