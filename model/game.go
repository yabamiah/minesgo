package model

type Flag string
const (
	easy Flag = "Easy"
	medium Flag = "Medium"
	hard Flag = "Hard"
	expert Flag = "Expert"
)

type Game struct {
	board Board
	gameOver bool
	gameWon bool
	modeGame Flag
}

func (g *Game) checkGameWon() {


}

func (g *Game) checkGameOver() {
	
}

func (g *Game) flagBoard(flag Flag) {
	if(flag == easy) {
		g.board.columns = 9
		g.board.rows = 9
	} else if (flag == medium) {
		g.board.columns = 18
		g.board.rows = 18
	} else if (flag == hard){
		g.board.columns = 27
		g.board.rows = 27
	} else {
		g.board.columns = 36
		g.board.rows = 36
	}
}