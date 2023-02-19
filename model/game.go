package model

type Flag string
const (
	small Flag = "tiny"
	medium Flag = "medium"
	large Flag = "large"
)

type Game struct {
	board Board
	gameOver bool
	gameWon bool
	modeGame Flag
}

func (g *Game) uncoverCell() {
	if (g.board.cells[0].isMine) {
		g.board.cells[0].cellContent = CELL_CONTENT_MINE
		g.board.cells[0].isRevealed = true	
	} else {
		g.board.cells[0].cellContent = CELL_CONTENT_EMPTY
		g.board.cells[0].isRevealed = true
	}
}

func (g *Game) checkGameWon() {

}

func (g *Game) checkGameOver() {
	
}

func (g *Game) flagBoard(flag Flag) {
	if(flag == small) {
		g.board.columns = 9
		g.board.rows = 9
	} else if (flag == large) {
		g.board.columns = 27
		g.board.rows = 27
	} else {
		g.board.columns = 18
		g.board.rows = 18
	}
}