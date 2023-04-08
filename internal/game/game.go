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

func (g *TicTacToeGame) MakeMove(moveRequest model.MoveRequest) model.MoveResponse {
	// If restart is true, reset the game state and update the difficulty
	if moveRequest.Restart {
		g.gameState.board = [9]int{}
		g.gameState.currentPlayer = XPlayer
		g.gameState.player = OPlayer
	} else {
		g.gameState.board = moveRequest.Board
		g.gameState.currentPlayer = moveRequest.Player
		g.gameState.player = GetOponent(moveRequest.Player)
	}

	g.gameState.difficulty = moveRequest.Difficulty
	fmt.Println(g.gameState, moveRequest.Restart)

	var message string = "Game Over."
	// Make a move and update the game state
	aiMove := g.gameState.MakeMove()
	success := false

	if aiMove != -1 {
		success = g.gameState.Play(aiMove)
		if success {
			message = "Player " + strconv.Itoa(moveRequest.Player) + " placed "
			if moveRequest.Player == XPlayer {
				message += "X"
			} else {
				message += "O"
			}
			message += " in position " + strconv.Itoa(aiMove) + "."
		} else {
			message = "Invalid move."
		}
	}

	// Check for winner and game status
	var gameStatus string
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
	}

	// Create a response
	return model.MoveResponse{
		Success:      success,
		Message:      message,
		Board:        g.gameState.board,
		BoardDisplay: boardToDisplay(g.gameState.board),
		GameStatus:   gameStatus,
		NextPlayer:   g.gameState.currentPlayer,
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
			display.WriteString("   ")
		}

		if i%3 != 2 {
			display.WriteString("|")
		}
	}

	return display.String()
}
