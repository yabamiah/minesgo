package model

import (
	"math/rand"
	"fmt"
)

type Board struct {
	columns int // number of columns
	rows int // number of rows
	qtdMines int // number of mines
	cells map[string]*Cell // cells of the board
}
	
	// PlaceMines randomly places mines on the board
	func (b *Board) PlaceMines() {
		if (b.columns == 9 && b.rows == 9) {
			b.qtdMines = 10
		} else if (b.columns == 18 && b.rows == 18) {
			b.qtdMines = 40
		} else {
			b.qtdMines = 99
		}

		for i := 0; i < b.qtdMines; i++ {
			for {
				x := rand.Intn(b.columns)
				y := rand.Intn(b.rows)
				key := fmt.Sprintf("%d,%d", x, y)
				if cell, exists := b.cells[key]; exists {
					if !cell.isMine {
						cell.isMine = true
						break
					}
				} else {
					b.cells[key] = &Cell{}
					b.cells[key].isMine = true
					break
				}
			}
		}
	}	

// calculateSurroundingMines calculates the number of mines surrounding a cell
func (b *Board) CalculateSurroundingMines(col, row int) int {
	for i := 0; i < b.columns; i++ {
		for j := 0; j < b.rows; j++ {
			return 0;
		}
	}
	return 0;
}