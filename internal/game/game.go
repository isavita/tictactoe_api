package game

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/isavita/tictactoe_api/internal/model"
)

type TicTacToeGame struct {
	gameState *GameState
}

func NewTicTacToeGame() *TicTacToeGame {
	return &TicTacToeGame{
		gameState: &GameState{
			board:         [9]int{},
			currentPlayer: XPlayer,
		},
	}
}

func (g *TicTacToeGame) MakeMove(currentPlayer int, moveRequest model.MoveRequest) model.MoveResponse {
	g.gameState.board = moveRequest.Board
	g.gameState.currentPlayer = currentPlayer
	g.gameState.player = GetOponent(currentPlayer)
	g.gameState.difficulty = moveRequest.Difficulty

	var message string = "Game Over."
	// Make a move and update the game state
	aiMove := g.gameState.MakeMove()
	success := false

	if aiMove != -1 {
		success = g.gameState.Play(aiMove)
		if success {
			message = "Player " + strconv.Itoa(currentPlayer) + " placed "
			if currentPlayer == XPlayer {
				message += "X"
			} else {
				message += "O"
			}
			message += " in position " + strconv.Itoa(aiMove+1) + "."
		} else {
			message = "Invalid move."
		}
	}

	// Check for winner and game status
	var gameStatus string
	var nextPlayer int = -1
	if g.gameState.HasWinner() {
		gameStatus = model.GameStatusPlayer1Wins
		if g.gameState.currentPlayer == OPlayer {
			gameStatus = model.GameStatusPlayer2Wins
		}
	} else if g.gameState.IsDraw() {
		gameStatus = model.GameStatusDraw
	} else {
		gameStatus = model.GameStatusOngoing
		g.gameState.NextTurn()
		nextPlayer = g.gameState.currentPlayer
	}

	// Create a response
	return model.MoveResponse{
		Success:      success,
		Message:      message,
		Board:        g.gameState.board,
		BoardDisplay: boardToDisplay(g.gameState.board),
		GameStatus:   gameStatus,
		NextPlayer:   nextPlayer,
	}
}

func boardToDisplay(board [9]int) string {
	var display strings.Builder

	for i := 0; i < len(board); i++ {
		if i > 0 && i%3 == 0 {
			display.WriteString("\n --------- \n")
		}

		switch board[i] {
		case XPlayer:
			display.WriteString(" X ")
		case OPlayer:
			display.WriteString(" O ")
		default:
			display.WriteString(fmt.Sprintf(" %d ", i+1))
		}

		if i%3 != 2 {
			display.WriteString("|")
		}
	}

	return display.String()
}
