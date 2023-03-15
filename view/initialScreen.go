package view

import (
    "image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/driver/desktop"
)

type InitialScreen struct {
    app fyne.App
    window fyne.Window

    title *canvas.Text
    subtitle *widget.Label

    newGameButton *widget.Button
    exitButton *widget.Button

    credits *widget.Label
}

func Initial() {
    myApp := app.New()
    window := myApp.NewWindow("MinesGo")
    window.Resize(fyne.NewSize(400, 400))

    is := InitialScreen{
        app: myApp,
        window: window,
    }

    // Create a title label with a large font size
    title := NewText("MinesGo", 42, fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: false})
    is.title = title

    myCanvas := canvas.NewText("Click me!", color.Black)
    meuclick := widget.Entry

	meuclick = func(e *desktop.MouseEvent) {
        if e.Button == desktop.LeftMouseButton {
            myCanvas.Text = "Clicked!"
            myCanvas.Refresh()
        } else if e.Button == desktop.RightMouseButton {
            myCanvas.Text = "Right clicked!"
            myCanvas.Refresh()
        }
    }

    // Create a subtitle label with a smaller font size
    subtitle := widget.NewLabelWithStyle("Get ready to play!", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: false})
    is.subtitle = subtitle

    // Create a button to start the game
    newGameButton := widget.NewButtonWithIcon("New game", theme.MediaPlayIcon(), func() {
        NewGameScreen(is.window)
    })
    newGameButton.Importance = widget.HighImportance
    is.newGameButton = newGameButton

    exitButton := widget.NewButtonWithIcon("Exit", theme.LogoutIcon(), func() {
        is.window.Close()
    })
    exitButton.Importance = widget.DangerImportance
    is.exitButton = exitButton

    // Use Fyne's layout system to position the widgets
    header := container.New(layout.NewVBoxLayout(),
        is.title,
        is.subtitle,
        is.newGameButton,
        is.exitButton,
    )

    // Create a label to display the game credits in the bottom of the screen
    credits := widget.NewLabelWithStyle("© 2021 MinesGo by Yabamiah", fyne.TextAlignCenter, fyne.TextStyle{Bold: false, Italic: true})
    is.credits = credits

    // Create a footer container
    footer := container.New(layout.NewMaxLayout(),
        is.credits,
    )

    // Combine the content and footer containers
    mainContent := NewMainContent(header, footer)

    // Set the content of the window and show it
    is.window.SetContent(mainContent)
    is.window.ShowAndRun()
}