package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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

func (is *InitialScreen) initTitle() {
    title := NewText("MinesGo", 42, fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: false})
    is.title = title
}

func (is *InitialScreen) initSubtitle() {
    subtitle := widget.NewLabelWithStyle("Get ready to play!", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: false})
    is.subtitle = subtitle
}

func (is *InitialScreen) initButtons() {
    newGameButton := widget.NewButtonWithIcon("New game", theme.MediaPlayIcon(), func() {
        InitialNewGameScreen(is.window)
    })
    newGameButton.Importance = widget.HighImportance
    is.newGameButton = newGameButton

    exitButton := widget.NewButtonWithIcon("Exit", theme.LogoutIcon(), func() {
        is.window.Close()
    })
    exitButton.Importance = widget.DangerImportance
    is.exitButton = exitButton
}

func (is *InitialScreen) initCredits() {
    credits := widget.NewLabelWithStyle("© 2021 MinesGo by Yabamiah", fyne.TextAlignCenter, fyne.TextStyle{Bold: false, Italic: true})
    is.credits = credits
}

func (is *InitialScreen) initWindow() {
    header := container.New(layout.NewVBoxLayout(),
        is.title,
        is.subtitle,
        is.newGameButton,
        is.exitButton,
    )

    footer := container.New(layout.NewMaxLayout(),
    is.credits,
    )

    mainContent := NewMainContent(header, footer)
    is.window.CenterOnScreen()
    is.window.SetContent(mainContent)
}

func (is *InitialScreen) initInitialScreen() {
    is.initTitle()
    is.initSubtitle()
    is.initButtons()
    is.initCredits()
    is.initWindow()
}

func Initial() {
    myTheme := BeutifulTheme{}

    myApp := app.New()
    myApp.Settings().SetTheme(myTheme)
    
    window := myApp.NewWindow("MinesGo")
    window.Resize(fyne.NewSize(400, 400))

    initialScreen := InitialScreen{ app: myApp, window: window }
    initialScreen.initInitialScreen()

    initialScreen.window.ShowAndRun()
}

func InitialBack(window fyne.Window) {
    initialScreen := InitialScreen{ window: window }

    initialScreen.initInitialScreen()
    initialScreen.window.Show()
}