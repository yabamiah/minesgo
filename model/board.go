package model

import (
	"fmt"
	"log"
	"math/rand"
)

// Board represents the board of the game
type Board struct {
	columns int // number of columns
	rows int // number of rows
	qtdMines int // number of mines
	cells map[string]*Cell // cells of the board
}

// GetCellPos returns the position of the cell
func (b *Board) GetCellPos() string {
	return fmt.Sprintf("%d,%d", b.columns, b.rows)
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

// uncoverCell uncovers a cell
func (b *Board) uncoverCell(pos string) {
	if (b.cells[pos].isMine) {
		b.cells[pos].cellContent = CELL_CONTENT_MINE
		b.cells[pos].isRevealed = true	
	} else {
		b.cells[pos].cellContent = CELL_CONTENT_EMPTY
		b.cells[pos].isRevealed = true
	}
}

// calculateSurroundingMines calculates the number of mines surrounding a cell
func (b *Board) CalculateSurroundingMines(pos string) int {
    surroundingCells := []string{
		// Convert the position to a string
        fmt.Sprintf("%d,%d", b.getCellX(pos)-1, b.getCellY(pos)-1),
        fmt.Sprintf("%d,%d", b.getCellX(pos)-1, b.getCellY(pos)),
        fmt.Sprintf("%d,%d", b.getCellX(pos)-1, b.getCellY(pos)+1),
        fmt.Sprintf("%d,%d", b.getCellX(pos), b.getCellY(pos)-1),
        fmt.Sprintf("%d,%d", b.getCellX(pos), b.getCellY(pos)+1),
        fmt.Sprintf("%d,%d", b.getCellX(pos)+1, b.getCellY(pos)-1),
        fmt.Sprintf("%d,%d", b.getCellX(pos)+1, b.getCellY(pos)),
        fmt.Sprintf("%d,%d", b.getCellX(pos)+1, b.getCellY(pos)+1),
    }
    
    var count int
	// Check if the cell exist and if the cell is a mine
    for _, cellPos := range surroundingCells {
        if cell := b.GetCell(cellPos); cell != nil && cell.isMine {
            count++
        }
    }
    
    return count
}

// getCell returns a cell through its position
func (b *Board) GetCell(pos string) *Cell {
    cell, exists := b.cells[pos]
    if !exists {
        return nil
    }
    return cell
}

// getCellX returns the x coordinate of a cell
func (b *Board) getCellX(pos string) int {
	var x, y int
    _, err := fmt.Sscanf(pos, "%d,%d", &x, &y)

	if err != nil {
		log.Fatal(err)
	}
    return x
}

// getCellY returns the y coordinate of a cell
func (b *Board) getCellY(pos string) int {
	var x, y int
	_, err := fmt.Sscanf(pos, "%d,%d", &x, &y)

	if err != nil {
		log.Fatal(err)
	}
    return y
}