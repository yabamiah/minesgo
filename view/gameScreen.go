package view

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Game struct {
	app     fyne.App
	window  fyne.Window

	title   *canvas.Text
	mines   *canvas.Text
	subtitle *widget.Label
	clock   *widget.Label

	board   *fyne.Container
	cells   []fyne.CanvasObject

	gameOn  bool
	
	buttonFlag *widget.Button
   	buttonMine *widget.Button
   	buttonBack *widget.Button
}

func (g *Game) initCells(tam int) {
	g.cells = make([]fyne.CanvasObject, tam*tam)

	for i := 0; i < tam*tam; i++ {
		g.cells[i] = widget.NewButton("", func() {
			fmt.Println("tapped", i)
		})
		g.cells[i].(*widget.Button).Importance = widget.LowImportance
		g.cells[i].(*widget.Button).Icon = theme.QuestionIcon()
	}

	for i := 0; i < tam*tam; i++ {
		cell := g.cells[i].(*widget.Button)
		cell.Icon = theme.QuestionIcon()
		x, y := i%tam, i/tam
		cell.OnTapped = func() {
			fmt.Println("tapped", x, y)
		}
	}
}

func (g *Game) initBoard() {
	board := container.New(layout.NewGridLayout(8), g.cells...)
	g.board = board
}

func (g *Game) initTitle() {
	title := NewText("MinesGo", 42, fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: false})
	mines := NewText("8 mines", 14, fyne.TextAlignCenter, fyne.TextStyle{Bold: false, Italic: true})
	g.title = title
	g.mines = mines
}

func (g *Game) initClock() {
	clock := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: false, Italic: true})
	g.clock = clock
}

func (g *Game) initButtons() {
	g.buttonFlag = widget.NewButtonWithIcon("Flag", theme.ColorChromaticIcon() , func() {
		fmt.Println("Flag")
	})
	g.buttonFlag.Importance = widget.HighImportance

	g.buttonMine = widget.NewButtonWithIcon("Mine", theme.ComputerIcon(), func() {
		fmt.Println("Mine")
	})
	g.buttonMine.Importance = widget.HighImportance
	
	g.buttonBack = widget.NewButtonWithIcon("Back", theme.NavigateBackIcon(), func() {
		InitialNewGameScreen(g.window)
	})
	g.buttonBack.Importance = widget.DangerImportance
}	

func (g *Game) initWindow() {
	header := container.New(layout.NewVBoxLayout(),
		g.title,
		g.mines,
		g.clock,
		g.buttonBack,
		g.board,
		g.buttonFlag,
		g.buttonMine,
	)

	g.window.CenterOnScreen()
	g.window.SetContent(header)
}

func (g *Game) updateClock() {
	start := time.Now()
	for g.gameOn {
		elapsed := time.Since(start)
		var formatted string
		if minutes := int(elapsed.Minutes()); minutes >= 1 {
			seconds := int(elapsed.Seconds()) - (minutes * 60)
			formatted = fmt.Sprintf("%dM %dS\n", minutes, seconds)
		} else {
			seconds := int(elapsed.Seconds())
			formatted = fmt.Sprintf("%dS\n", seconds)
		}
		g.clock.SetText(formatted)
		time.Sleep(time.Second)
	}
}

func (g *Game) initGame(tam int) {
	g.initCells(tam)
	g.initBoard()
	g.initTitle()
	g.initClock()
	g.initButtons()
	g.initWindow()
}

func (g *Game) startGame() {
	g.gameOn = true
	go g.updateClock()
}

func (g *Game) stopGame() {
	g.gameOn = false
}

func InitialGameScreen(window fyne.Window) {
	app := fyne.CurrentApp()

	gameScreen := Game{ app: app, window: window }
	gameScreen.initGame(8)
	gameScreen.startGame()
	
	gameScreen.window.Show()
}
