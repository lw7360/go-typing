package main

import (
	"github.com/nsf/termbox-go"
	"os"
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
const colErr = termbox.ColorRed

func drawCentered(default_fg termbox.Attribute, default_bg termbox.Attribute, template []string) {
	termbox.Clear(colDef, colDef)
	termbox.HideCursor()
	width, height := termbox.Size()
	start_x := (width) / 2
	start_y := (height - len(template)) / 2
	for index_y, line := range template {
		lineLength := len(line)
		for index_x, runeValue := range line {
			displayRune := ' '
			if runeValue != ' ' {
				displayRune = runeValue
			}
			termbox.SetCell(start_x+index_x-lineLength/2, start_y+index_y, displayRune, default_fg, default_bg)
		}
	}
	termbox.Flush()
}

func drawMainScreen(default_fg termbox.Attribute, default_bg termbox.Attribute) {
	template := []string{
		"GoTyping",
		"",
		"[1] Practice",
		"[2] Stats   ",
		"[3] About   ",
		"",
		"[Esc] to quit",
		"[h] for help",
	}
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
	wordsFile := "words.txt"
	statsFile := "stats.json"
	if len(args) > 0 {
		wordsFile = args[0]
	}

	var curGame *Game

	fgColor := termbox.ColorWhite
	bgColor := termbox.ColorDefault

mainloop:
	for {
		switch curScreen {
		case MainScreen:
			drawMainScreen(fgColor, bgColor)
		case GameScreen:
			curGame = NewGame(wordsFile, statsFile)

		gameloop:
			for {
				drawGameScreen(fgColor, bgColor, curGame)

				ev := termbox.PollEvent()
				curGame.initTime()

				switch ev.Key {
				case termbox.KeyEsc:
					curGame.saveStats(statsFile)
					curScreen = MainScreen
					drawMainScreen(fgColor, bgColor)
					break gameloop
				case termbox.KeyBackspace, termbox.KeyBackspace2:
					curGame.curInd--
					if curGame.curInd < 0 {
						curGame.curInd = 0
					}
					continue gameloop
				case termbox.KeySpace:
					ev.Ch = ' '
				}

				curInd := curGame.curInd
				curChar := curGame.getRune(curInd)

				if ev.Ch != curChar {
					curGame.setErr(curInd, true)
					curGame.curStats.Errors++
				} else {
					curGame.setErr(curInd, false)
					if curChar == ' ' && curGame.noErr(curInd) {
						curGame.curStats.Words++
					}
				}
				curGame.curInd++
			}
		case StatsScreen:
			drawStatsScreen(fgColor, bgColor, curGame)
			ev := termbox.PollEvent()
			switch ev.Key {
			case termbox.KeyEsc:
				curScreen = MainScreen
				continue mainloop
			}
			switch ev.Ch {
			case 'r':
				curGame.curStats.reset()
				curGame.stats.reset()
				curGame.saveStats(statsFile)
				continue mainloop
			}
		case AboutScreen:
			drawAboutScreen(fgColor, bgColor)
		}

		ev := termbox.PollEvent()
		switch ev.Key {
		case termbox.KeyEsc:
			if curScreen == MainScreen {
				break mainloop
			}
			curScreen = MainScreen
			continue
		}
		switch curScreen {
		case MainScreen:
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
}
