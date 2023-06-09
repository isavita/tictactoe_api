package game

import (
	"fmt"
	"math"
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
			board:         make([]int, 9),
			boardSize:     3,
			currentPlayer: XPlayer,
		},
	}
}

func (g *TicTacToeGame) MakeMove(currentPlayer int, moveRequest model.MoveRequest) model.MoveResponse {
	g.gameState.board = moveRequest.Board
	g.gameState.boardSize = moveRequest.BoardSize
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
			message = "Player " + strconv.Itoa(currentPlayer) + " has placed "
			if currentPlayer == XPlayer {
				message += "'X'"
			} else {
				message += "'O'"
			}
			message += " in position " + strconv.Itoa(aiMove+1) + "."

			if isFirstMove(g.gameState.board) {
				message += " Please note: If the user is playing with 'X', disregard this move. Instead, ask the user where they would like to place their first move, then present the game board reflecting their choice."
			}
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
		BoardSize:    g.gameState.boardSize,
		BoardDisplay: boardToDisplay(g.gameState.board),
		GameStatus:   gameStatus,
		NextPlayer:   nextPlayer,
	}
}

func isFirstMove(board []int) bool {
	count := 0
	for _, val := range board {
		if val != 0 {
			count++
			if count > 1 {
				return false
			}
		}
	}
	return true
}

func boardToDisplay(board []int) string {
	size := int(math.Sqrt(float64(len(board))))
	if size > 3 {
		return boardToDisplayWhenBig(board, size)
	} else {
		return boardToDisplayWhenSmall(board, size)
	}
}

func boardToDisplayWhenSmall(board []int, size int) string {
	var display strings.Builder

	for i := 0; i < len(board); i++ {
		if i > 0 && i%size == 0 {
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

		if i%size != 2 {
			display.WriteString("|")
		}
	}

	return display.String()
}

func boardToDisplayWhenBig(board []int, size int) string {
	var display strings.Builder

	for i := 0; i < len(board); i++ {
		if i > 0 && i%size == 0 {
			display.WriteString("\n")
			for j := 0; j < size; j++ {
				display.WriteString("-----")
			}
			display.WriteString("\n")
		}

		switch board[i] {
		case XPlayer:
			display.WriteString("  X ")
		case OPlayer:
			display.WriteString("  O ")
		default:
			display.WriteString(fmt.Sprintf(" %2d ", i+1))
		}

		if i%size != size-1 {
			display.WriteString("|")
		}
	}

	return display.String()
}
