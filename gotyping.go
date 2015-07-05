package main

import (
	"github.com/nsf/termbox-go"
	"io/ioutil"
	"os"
	"strings"
)

type DisplayScreen int

const (
	MainScreen = iota
	GameScreen
	StatsScreen
	AboutScreen
)

var curScreen = MainScreen

const colDef = termbox.ColorDefault

var wordList = []string{}

func loadWords(filename string) bool {
	words, err := ioutil.ReadFile(filename)
	if err != nil {
		return false
	}
	wordList = strings.Split(string(words), "\n")
	return true
}

func redrawAll() {
	termbox.Clear(colDef, colDef)
	// width, height := termbox.Size()
}

func drawCentered(default_fg termbox.Attribute, default_bg termbox.Attribute, template []string) {
	termbox.Clear(colDef, colDef)
	width, height := termbox.Size()
	start_x := (width) / 2
	start_y := (height - len(template)) / 2
	for index_y, line := range template {
		lineLength := len(line)
		for index_x, runeValue := range line {
			displayRune := ' '
			if runeValue != ' ' {
				if runeValue != '#' {
					displayRune = runeValue
				}
			}
			termbox.SetCell(start_x+index_x-lineLength/2, start_y+index_y, displayRune, default_fg, default_bg)
		}
	}
	termbox.Flush()
}

func drawMainScreen(default_fg termbox.Attribute, default_bg termbox.Attribute) {
	template := []string{
		"GoTyping.",
		"",
		"1: Practice",
		"2: Stats",
		"3: About",
	}
	drawCentered(default_fg, default_bg, template)
}

func drawGameScreen(default_fg termbox.Attribute, default_bg termbox.Attribute) {
	template := []string{"INGAME"}
	drawCentered(default_fg, default_bg, template)
}

func drawStatsScreen(default_fg termbox.Attribute, default_bg termbox.Attribute) {
	template := []string{"Stats"}
	drawCentered(default_fg, default_bg, template)
}

func drawAboutScreen(default_fg termbox.Attribute, default_bg termbox.Attribute) {
	template := []string{"About GoTyping"}
	drawCentered(default_fg, default_bg, template)
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	// load wordList
	args := os.Args[1:]
	filename := "words.txt"
	if len(args) > 0 {
		filename = args[0]
	}
	if !loadWords(filename) {
		panic("Failed to load: " + filename)
	}

mainloop:
	for {
		switch curScreen {
		case MainScreen:
			drawMainScreen(termbox.ColorWhite, termbox.ColorDefault)
		case GameScreen:
			drawGameScreen(termbox.ColorWhite, termbox.ColorDefault)
		case StatsScreen:
			drawStatsScreen(termbox.ColorWhite, termbox.ColorDefault)
		case AboutScreen:
			drawAboutScreen(termbox.ColorWhite, termbox.ColorDefault)
		}
		ev := termbox.PollEvent()
		switch ev.Key {
		case termbox.KeyEsc:
			if curScreen == MainScreen {
				break mainloop
			}
			curScreen = MainScreen
		}
		switch ev.Ch {
		case '1':
			curScreen = GameScreen
		case '2':
			curScreen = StatsScreen
		case '3':
			curScreen = AboutScreen
		}
	}
}
