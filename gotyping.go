package main

import (
	"github.com/nsf/termbox-go"
)

type DisplayScreen int

const (
	MainScreen = iota
	GameScreen
	AboutScreen
)

var curScreen = MainScreen

const colDef = termbox.ColorDefault

func redrawAll() {
	termbox.Clear(colDef, colDef)
	// width, height := termbox.Size()
}

func drawMainScreen(default_fg termbox.Attribute, default_bg termbox.Attribute) {
	termbox.Clear(colDef, colDef)
}

func drawAboutScreen(default_fg termbox.Attribute, default_bg termbox.Attribute) {
	termbox.Clear(colDef, colDef)
	width, height := termbox.Size()
	template := [...]string{"GoTyping"}

	first_line := template[0]
	start_x := (width - len(first_line)) / 2
	start_y := (height - len(template)) / 2
	for index_y, line := range template {
		for index_x, runeValue := range line {
			bg := default_bg
			displayRune := ' '
			if runeValue != ' ' {
				bg = termbox.Attribute(125)
				if runeValue != '#' {
					displayRune = runeValue
				}
			}
			termbox.SetCell(start_x+index_x, start_y+index_y, displayRune, default_fg, bg)
		}
	}
	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

mainloop:
	for {
		drawMainScreen(termbox.ColorWhite, termbox.ColorBlack)
		ev := termbox.PollEvent()
		if ev.Key == termbox.KeyEsc {
			break mainloop
		}
		drawAboutScreen(termbox.ColorWhite, termbox.ColorBlack)
	}
}
