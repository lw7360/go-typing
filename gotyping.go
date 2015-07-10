package main

import (
	"github.com/nsf/termbox-go"
	"os"
	"strconv"
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

func statsString(curGame *Game) string {
	curStats := &curGame.curStats
	curStats.Seconds = curGame.gameTime()

	words := strconv.Itoa(curStats.Words)
	errors := strconv.Itoa(curStats.Errors)
	wpm := strconv.Itoa(int(curStats.wpm()))

	statsString := "Words: " + words + " | Errors: " + errors + " | WPM: " + wpm + " | [Esc] to quit"

	return statsString
}

func drawGameScreen(default_fg termbox.Attribute, default_bg termbox.Attribute, curGame *Game) {
	termbox.Clear(colDef, colDef)

	width, height := termbox.Size()
	i := 0
	for y := 0; y < height-2; y = y + 2 {
		for x := 0; x < width; x++ {
			fg, bg := default_fg, default_bg
			if i == curGame.curInd {
				termbox.SetCursor(x, y)
			}
			_, err := curGame.errMap[i]
			if err {
				bg = colErr
			}

			displayRune := curGame.getRune(i)
			i++
			termbox.SetCell(x, y, displayRune, fg, bg)
		}
	}

	for x := 0; x < width; x++ {
		termbox.SetCell(x, height-2, '_', default_fg, default_bg)
	}
	statsString := statsString(curGame)
	for x := 0; x < len(statsString); x++ {
		termbox.SetCell(x, height-1, rune(statsString[x]), termbox.ColorGreen, default_bg)
	}
	termbox.Flush()
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
	wordsFile := "words.txt"
	statsFile := "stats.json"
	if len(args) > 0 {
		wordsFile = args[0]
	}

	fgColor := termbox.ColorWhite
	bgColor := termbox.ColorDefault

mainloop:
	for {
		switch curScreen {
		case MainScreen:
			drawMainScreen(fgColor, bgColor)
		case GameScreen:
			curGame := NewGame(wordsFile, statsFile)
			curGame.initTime()

			furthestInd := 0
		gameloop:
			for {
				drawGameScreen(fgColor, bgColor, curGame)

				ev := termbox.PollEvent()

				switch ev.Key {
				case termbox.KeyEsc:
					curGame.saveStats(statsFile)
					curScreen = MainScreen
					continue mainloop
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
					curGame.errMap[curInd] = struct{}{}
					curGame.curStats.Errors++
				} else {
					delete(curGame.errMap, curInd)
					if curChar == ' ' && curInd > furthestInd && curGame.noErr(curInd) {
						furthestInd = curInd
						curGame.curStats.Words++
					}
				}
				curGame.curInd++
			}
		case StatsScreen:
			drawStatsScreen(fgColor, bgColor)
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
