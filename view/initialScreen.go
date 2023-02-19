package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/theme"
)

func Initial() {
    myApp := app.New()
    myWindow := myApp.NewWindow("MinesGo")
    myWindow.Resize(fyne.NewSize(600, 600))

    // Create a title label with a large font size
    title := NewText("MinesGo", 42, fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: false})

    // Create a subtitle label with a smaller font size
    subtitle := widget.NewLabelWithStyle("Get ready to play!", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: false})


    // Create a button to start the game
    startButton := widget.NewButtonWithIcon("Start", theme.MediaPlayIcon(), func() {
        // Start the game here
    })

    exitButton := widget.NewButtonWithIcon("Exit", theme.LogoutIcon(), func() {
        myWindow.Close()
    })

    // Use Fyne's layout system to position the widgets
    content := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
        title,
        subtitle,
        startButton,
        exitButton,
    )

    // Set the content of the window and show it
    myWindow.SetContent(content)
    myWindow.ShowAndRun()
}
