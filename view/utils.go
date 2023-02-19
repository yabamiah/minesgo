package view

import (
    "image/color"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
)

// NewText returns a new Text implementation
func NewText(text string, size float32, alignment fyne.TextAlign, textStyle fyne.TextStyle) *canvas.Text {
    t := canvas.NewText(text, color.Black)
    t.TextSize = size
    t.Alignment = alignment
    t.TextStyle = textStyle
    return t
}
