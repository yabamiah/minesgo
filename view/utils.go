package view

import (
	"github.com/yabamiah/minesgo/model"

	"image/color"
    "fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
)

// NewText returns a new Text implementation
func NewText(text string, size float32, alignment fyne.TextAlign, textStyle fyne.TextStyle) *canvas.Text {
    t := canvas.NewText(text, color.Black)
    t.TextSize = size
    t.Alignment = alignment
    t.TextStyle = textStyle
    return t
}

// NewMainContent is a function that will create the main content of the screen
func NewMainContent(header, footer *fyne.Container) *fyne.Container {
    return fyne.NewContainerWithLayout(layout.NewBorderLayout(header, footer, nil, nil),
        header,
        footer,
    )
}

// FormatPosition is a function that will format the position of the cell in the board
func FormatPosition(x, y int) string {
    return fmt.Sprintf("%d,%d", x, y)
}

// IndentifyClick is a function that will identify the type of click
func IndentifyClick(mouseBtn desktop.MouseButton, board *model.Board, pos string) {
    switch mouseBtn {
    case desktop.MouseButtonPrimary:
        LeftClick(board, pos)
    case desktop.MouseButtonSecondary:
        RightClick(board, pos)
    }
}

// RightClick is a function that handles the right click event
func RightClick(board *model.Board, pos string) {
    cell := board.GetCell(pos)

    if cell.GetIsRevealed() {
        return
    }

    if cell.GetIsFlagged() {
        cell.SetIsFlagged(false)
    } else {
        cell.SetIsFlagged(true)
    }
}   

// LeftClick is the function that will be called when the user clicks on a cell
func LeftClick(board *model.Board, pos string) {
    cell := board.GetCell(pos)

    if cell.GetIsRevealed() {
        return
    }

    if cell.GetIsFlagged() {
        return
    }

    if cell.GetIsMine() {
        cell.SetIsRevealed(true)
        return
    }

    cell.SetIsRevealed(true)
    cell.SetCellContent(fmt.Sprint((board.CalculateSurroundingMines(pos))))
}