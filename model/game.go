package model

import (
	//"math/rand"
)

// Type Flag is a string to define the game mode
type Flag string
const (
	easy Flag = "Easy"
	medium Flag = "Medium"
	hard Flag = "Hard"
	expert Flag = "Expert"
)


// Game struct is the main struct of the game
type Game struct {
	board Board // Board is a struct
	gameOver bool // gameOver is a bool
	gameWon bool // gameWon is a bool
	modeGame Flag // modeGame is a Flag
}

func NewGame(flag string) *Game {
	return &Game{
		board: Board{},
		gameOver: false,
		gameWon: false,
		modeGame: Flag(flag),
	}
}

func (g *Game) StartGame() {
	g.flagBoard(g.modeGame)
	//g.firstClick()
	g.board.PlaceMines()
}

func (g *Game) firstClick(x, y int) {
	
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