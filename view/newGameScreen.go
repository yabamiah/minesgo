package view

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var flag string

type NewGame struct {
	app fyne.App
	window fyne.Window

	title *canvas.Text
	subtitle *canvas.Text
	credits *canvas.Text

	comboGameMode *widget.Select
	startButton *widget.Button
	backButton *widget.Button
}

func (ng *NewGame) initTitle() {
	title := NewText("New Game", 42, fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: false})
	ng.title = title
}

func (ng *NewGame) initSubtitle() {
	subtitle := NewText("Choose game difficulty:", 14, fyne.TextAlignCenter, fyne.TextStyle{Bold: false, Italic: false})
	ng.subtitle = subtitle
}

func (ng *NewGame) initCredits() {
	credits := NewText("© 2021 MinesGo by Yabamiah", 14, fyne.TextAlignCenter, fyne.TextStyle{Bold: false, Italic: true})
	ng.credits = credits
}

func (ng *NewGame) initComboGameMode() {
	comboGameMode := widget.NewSelect([]string{"Easy", "Normal", "Hard", "Expert"}, func(value string) {
		log.Println("Select the difficulty", value)
		flag = value
	})
	ng.comboGameMode = comboGameMode
}

func (ng *NewGame) initStartButton() {
	startButton := widget.NewButtonWithIcon("Start", theme.MediaPlayIcon(), func() {
		InitialGameScreen(ng.window)
	})
	startButton.Importance = widget.HighImportance
	ng.startButton = startButton
}

func (ng *NewGame) initBackButton() {
	backButton := widget.NewButtonWithIcon("Back", theme.NavigateBackIcon(), func() {
		InitialBack(ng.window)
	})
	backButton.Importance = widget.DangerImportance
	ng.backButton = backButton
}

func (ng *NewGame) initWindow() {
	header := container.New(layout.NewVBoxLayout(),
		ng.title,
		ng.subtitle,
		ng.comboGameMode,
		ng.startButton,
		ng.backButton,
	)

	footer := container.New(layout.NewMaxLayout(),
		ng.credits,
	)

	mainContent := NewMainContent(header, footer)

	ng.window.CenterOnScreen()
	ng.window.SetContent(mainContent)
}

func (ng *NewGame) initNewGame() {
	ng.initTitle()
	ng.initSubtitle()
	ng.initCredits()
	ng.initComboGameMode()
	ng.initStartButton()
	ng.initBackButton()
	ng.initWindow()
}

func InitialNewGameScreen(window fyne.Window) {
	app := fyne.CurrentApp()

	newGame := NewGame{app: app, window: window}
	newGame.initNewGame()
	
	newGame.window.Show()
}