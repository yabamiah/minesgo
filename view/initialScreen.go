package view

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
)

func Initial() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Welcome to My Game")

    // Create a title label with a large font size
    title := widget.NewLabelWithStyle("My Game", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: true})

    // Create a subtitle label with a smaller font size
    subtitle := widget.NewLabelWithStyle("Get ready to play!", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: false})

    // Create a button to start the game
    startButton := widget.NewButton("Start", func() {
        // Start the game here
    })

    // Use Fyne's layout system to position the widgets
    content := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
        title,
        subtitle,
        startButton,
    )

    // Set the content of the window and show it
    myWindow.SetContent(content)
    myWindow.ShowAndRun()
}
