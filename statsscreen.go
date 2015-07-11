package main

import (
	"github.com/nsf/termbox-go"
	"strconv"
)

func drawStatsScreen(default_fg termbox.Attribute, default_bg termbox.Attribute, curGame *Game) {
	totalStats := new(Stats)
	totalStats.Words = curGame.stats.Words + curGame.curStats.Words
	totalStats.Seconds = curGame.stats.Seconds + curGame.curStats.Seconds
	totalStats.Errors = curGame.stats.Errors + curGame.curStats.Errors

	template := []string{
		"Total Stats",
		"",
		"Words: " + strconv.Itoa(totalStats.Words),
		"Errors: " + strconv.Itoa(totalStats.Errors),
		"WPM: " + strconv.Itoa(totalStats.wpm()),
		"",
		"[Esc] to go back",
		"[r] to reset all stats",
	}

	drawCentered(default_fg, default_bg, template)
}
