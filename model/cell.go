package model

import (
	"fmt"
)

// Cell is a struct that represents a cell in the board
type Cell struct {
	isMine bool // isMine is true if the cell has a mine
	isRevealed bool // isRevealed is true if the cell is revealed
	isFlagged bool // isFlagged is true if the cell is flagged
	isFlaggedWrong bool // isFlaggedWrong is true if the cell is flagged wrong
	isFlaggedRight bool // isFlaggedRight is true if the cell is flagged right
	qtdMinesAround int // qtdMinesAround is the quantity of mines around the cell
	cellContent string // cellContent is the content of the cell
}

// Constants that represents the content of the cell
const (
	CELL_CONTENT_FILL = " "
	CELL_CONTENT_EMPTY = " "
	CELL_CONTENT_MINE = "resources/bomb_icon.png"
	CELL_CONTENT_FLAG = "resources/flag-fill-black.png"
	CELL_CONTENT_RED_FLAG = "resources/flag-fill-red.png"
	CELL_CONTENT_GRAY_FLAG = "resources/flag-fill-gray.png"
)

func (c *Cell) GetIsMine() bool {
	return c.isMine
}

func (c *Cell) SetIsMine(isMine bool) {
	c.isMine = isMine
}

func (c *Cell) GetIsRevealed() bool {
	return c.isRevealed
}

func (c *Cell) SetIsRevealed(isRevealed bool) {
	c.isRevealed = isRevealed
}

func (c *Cell) GetIsFlagged() bool {
	return c.isFlagged
}

func (c *Cell) SetIsFlagged(isFlagged bool) {
	c.isFlagged = isFlagged
}

func (c *Cell) GetIsFlaggedWrong() bool {
	return c.isFlaggedWrong
}

func (c *Cell) SetIsFlaggedWrong(isFlaggedWrong bool) {
	c.isFlaggedWrong = isFlaggedWrong
}

func (c *Cell) GetIsFlaggedRight() bool {
	return c.isFlaggedRight
}

func (c *Cell) SetIsFlaggedRight(isFlaggedRight bool) {
	c.isFlaggedRight = isFlaggedRight
}

func (c *Cell) GetQtdMinesAround() int {
	return c.qtdMinesAround
}

func (c *Cell) SetQtdMinesAround(qtdMinesAround int) {
	c.qtdMinesAround = qtdMinesAround
}

func (c *Cell) GetCellContent() string {
	return c.cellContent
}

func (c *Cell) SetCellContent(cellContent string) {
	if c.isFlagged {
		c.cellContent = CELL_CONTENT_FLAG
	} else if c.isFlaggedWrong {
		c.cellContent = CELL_CONTENT_RED_FLAG
	} else if c.isFlaggedRight {
		c.cellContent = CELL_CONTENT_GRAY_FLAG
	} else if c.isMine {
		c.cellContent = CELL_CONTENT_MINE
	} else {
		// TODO: To place a icon with the number of mines around
		switch cellContent {
		case "0":
			c.cellContent = CELL_CONTENT_EMPTY
		case "1":
			c.cellContent = "resources/1.png"
		case "2":
			c.cellContent = "resources/2.png"
		case "3":
			c.cellContent = "resources/3.png"
		case "4":
			c.cellContent = "resources/4.png"
		case "5":
			c.cellContent = "resources/5.png"
		case "6":
			c.cellContent = "resources/6.png"
		case "7":
			c.cellContent = "resources/7.png"
		case "8":
			c.cellContent = "resources/8.png"
		default:
			fmt.Println("PANIC: Invalid number of mines around")
		}
	}
}

/*
func (c *Cell) DefineIcons() {
	if c.isMine {
		c.cellContent = CELL_CONTENT_MINE
	} else if c.isFlagged {
		c.cellContent = CELL_CONTENT_FLAG
	} else if c.isFlaggedWrong {
		c.cellContent = CELL_CONTENT_GRAY_FLAG
	} else if c.isFlaggedRight {
		c.cellContent = CELL_CONTENT_RED_FLAG
	}
}
*/