package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Initial() {
    myApp := app.New()
    window := myApp.NewWindow("MinesGo")
    window.Resize(fyne.NewSize(400, 400))

    // Create a title label with a large font size
    title := NewText("MinesGo", 42, fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: false})
    title.Move(fyne.NewPos(190, 20))

    // Create a subtitle label with a smaller font size
    subtitle := widget.NewLabelWithStyle("Get ready to play!", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: false})
    subtitle.Move(fyne.NewPos(190, 100))

    // Create a button to start the game
    newGameButton := widget.NewButtonWithIcon("New game", theme.MediaPlayIcon(), func() {
        NewGameScreen(window)
    })
    newGameButton.Importance = widget.HighImportance
    newGameButton.Resize(fyne.NewSize(125, 40))
    newGameButton.Move(fyne.NewPos(130, 150))

    exitButton := widget.NewButtonWithIcon("Exit", theme.LogoutIcon(), func() {
        window.Close()
    })
    exitButton.Importance = widget.DangerImportance
    exitButton.Resize(fyne.NewSize(125, 40))
    exitButton.Move(fyne.NewPos(130, 210))

    // Use Fyne's layout system to position the widgets
    header := fyne.NewContainerWithoutLayout(
        title,
        subtitle,
        newGameButton,
        exitButton,
    )

    // Create a label to display the game credits in the bottom of the screen
    credits := widget.NewLabelWithStyle("© 2021 MinesGo by Yabamiah", fyne.TextAlignCenter, fyne.TextStyle{Bold: false, Italic: true})

    // Create a footer container
    footer := fyne.NewContainerWithLayout(layout.NewMaxLayout(),
        credits,
    )

    mainContent := NewMainContent(header, footer)

    // Set the content of the window and show it
    window.SetContent(mainContent)
    window.ShowAndRun()
}