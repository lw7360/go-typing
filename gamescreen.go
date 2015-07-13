package main

import (
	"github.com/nsf/termbox-go"
	"strconv"
)

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
	i := curGame.curInd / width * width
	for y := 0; y < height-2; y = y + 2 {
		for x := 0; x < width; x++ {
			fg, bg := default_fg, default_bg
			if i == curGame.curInd {
				termbox.SetCursor(x, y)
			}
			if curGame.indexHasErr(i) {
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
