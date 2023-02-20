package model

type Cell struct {
	isMine bool
	isRevealed bool
	isFlagged bool
	isFlaggedWrong bool
	isFlaggedRight bool
	qtdMinesAround int
	cellContent string
}

const (
	CELL_CONTENT_FILL = " "
	CELL_CONTENT_EMPTY = " "
	CELL_CONTENT_MINE = "resources/bomb_icon.png"
	CELL_CONTENT_FLAG = "resources/flag-fill-black.png"
	CELL_CONTENT_RED_FLAG = "resources/flag-fill-red.png"
	CELL_CONTENT_GRAY_FLAG = "resources/flag-fill-gray.png"
)

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