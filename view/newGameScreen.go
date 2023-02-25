package view

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

var flag string

func NewGameScreen(window fyne.Window) {
	// Create a title label with a large font size
	title := NewText("New Game", 42, fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: false})
	
	// Create a subtitle label with a smaller font size
    subtitle := widget.NewLabelWithStyle("Choose game difficulty:", fyne.TextAlignCenter, fyne.TextStyle{Bold: false, Italic: false})

	// Create a select to choose the game difficulty
	comboGameMode := widget.NewSelect([]string{"Easy", "Normal", "Hard", "Expert"}, func(value string) {
		log.Println("Select the difficulty", value)
		flag = value
	})

	// Create a button to start the game
	startButton := widget.NewButtonWithIcon("Start", theme.MediaPlayIcon(), func() {
		// NewGameScreen(window)
	})
	startButton.Importance = widget.HighImportance


	backButton := widget.NewButtonWithIcon("Back", theme.NavigateBackIcon(), func() {
		// I still don't know how to go back to the previous screen
	})
	backButton.Importance = widget.DangerImportance

	// Use Fyne's layout system to position the widgets in the header
	header := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
        title,
        subtitle,
        comboGameMode,
		startButton,
		backButton,
    )

	// Create a label to display the game credits in the bottom of the screen
	credits := widget.NewLabelWithStyle("© 2021 MinesGo by Yabamiah", fyne.TextAlignCenter, fyne.TextStyle{Bold: false, Italic: true})
	footer := fyne.NewContainerWithLayout(layout.NewMaxLayout(),
		credits,
	)

	// Combine the content and footer containers
	mainContent := NewMainContent(header, footer)

    // Set the content of the window and show it
    window.SetContent(mainContent)
    window.Show()

}