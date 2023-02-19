package model

import (
	"math/rand"
)

type Board struct {
	columns int  // number of columns
	rows int     // number of rows
	qtdMines int // number of mines
	cells []Cell // cells of the board
}

// placeMines randomly places mines on the board
func (b *Board) placeMines() {
	if (b.columns == 9 && b.rows == 9) {
		b.qtdMines = 10
		count := 0
		for i := 0; i < rand.Intn(9); i++ {
			for j := 0; j < rand.Intn(9); j++ {
				if (count != 10){
					b.cells[j].isMine = true	
				}
			}
		}
	} else if (b.columns == 18 && b.rows == 18) {
		b.qtdMines = 40
		count := 0
		for i := 0; i < rand.Intn(18); i++ {
			for j := 0; j < rand.Intn(18); j++ {
				if (count != 40){
					b.cells[j].isMine = true	
				}
			}
		}
	} else {
		b.qtdMines = 99
		count := 0
		for i := 0; i < rand.Intn(27); i++ {
			for j := 0; j < rand.Intn(27); j++ {
				if (count != 99){
					b.cells[j].isMine = true	
				}
			}
		}
	}
}

// calculateSurroundingMines calculates the number of mines surrounding a cell
func (b *Board) calculateSurroundingMines() int {
	
	return 0;
}