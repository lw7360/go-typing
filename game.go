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

type Word struct {
	word   string
	errMap map[int]struct{}
}

type Game struct {
	wordList  []Word    // Slice of all Words
	curInd    int       // Index of current char
	curStats  Stats     // Stats for current session
	stats     Stats     // Stats for current session + all past sessions
	startTime time.Time // Start time of current session
}

func (g *Game) loadWords(filename string) bool {
	rand.Seed(time.Now().UnixNano())

	words, err := ioutil.ReadFile(filename)
	if err != nil {
		return false
	}

	wordSlice := strings.Split(string(words), "\n")
	shuffle(wordSlice)

	for _, word := range wordSlice {
		g.wordList = append(g.wordList, Word{word, make(map[int]struct{})})
	}

	return true
}

func (g *Game) loadStats(filename string) bool {
	return g.stats.loadStats(filename)
}

func (g *Game) saveStats(filename string) bool {
	totalStats := new(Stats)
	totalStats.Words = g.stats.Words + g.curStats.Words
	totalStats.Seconds = g.stats.Seconds + g.curStats.Seconds
	totalStats.Errors = g.stats.Errors + g.curStats.Errors

	return totalStats.saveStats(filename)
}

func (g *Game) initTime() {
	if g.startTime.IsZero() {
		g.startTime = time.Now()
	}
}

func (g *Game) gameTime() float64 {
	return time.Since(g.startTime).Seconds()
}

func (g *Game) getRune(index int) rune {
	curIndex := 0
	for _, word := range g.wordList {
		for _, char := range word.word {
			if curIndex == index {
				return char
			}
			curIndex++
		}
		if curIndex == index { // Adds a space inbetween words
			return ' '
		}
		curIndex++
	}

	return g.getRune(index - curIndex) // Recurse if we run out of words
}

func (g *Game) getWordAndIndex(index int) (word Word, curWordInd int) {
	curIndex := 0

	for _, word := range g.wordList {
		for i, _ := range word.word {
			curWordInd = i
			if curIndex == index {
				return word, curWordInd
			}
			curIndex++
		}
		curWordInd++
		if curIndex == index {
			return word, curWordInd
		}
		curIndex++
	}

	return g.getWordAndIndex(index - curIndex)
}

func (g *Game) noErr(index int) bool { // Checks if word at index has any errors.
	curWord, _ := g.getWordAndIndex(index)
	return len(curWord.errMap) == 0
}

func (g *Game) indexHasErr(index int) bool {
	curWord, curWordInd := g.getWordAndIndex(index)
	_, ok := curWord.errMap[curWordInd]
	return ok
}

func (g *Game) setErr(index int, val bool) {
	curWord, curInd := g.getWordAndIndex(index)
	if val {
		curWord.errMap[curInd] = struct{}{}
	} else {
		delete(curWord.errMap, curInd)
	}
}

func NewGame(wordsFile string, statsFile string) *Game {
	g := new(Game)
	if !g.loadWords(wordsFile) {
		panic("Failed to load wordsFile: \"" + wordsFile + "\"")
	}
	g.loadStats(statsFile)

	return g
}
