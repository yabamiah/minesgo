package model

type Cell struct {
	isMine bool
	isRevealed bool
	qtdMinesAround int
	cellContent string
}

const (
	CELL_CONTENT_FILL = "▇"
	CELL_CONTENT_EMPTY = " "
	CELL_CONTENT_MINE = "✱"
	CELL_CONTENT_FLAG = "⍝"
)