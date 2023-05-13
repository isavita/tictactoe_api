package game

import (
	"math"
	"math/rand"
)

const (
	XPlayer          = 1
	OPlayer          = 2
	Draw             = 3
	DifficultyEasy   = 101
	DifficultyMedium = 102
	DifficultyHard   = 103
)

type GameState struct {
	board         [9]int
	currentPlayer int
	player        int
	difficulty    int
}

func (gs *GameState) Play(move int) bool {
	if move < 0 || move >= 9 {
		return false
	}

	if gs.board[move] == 0 {
		gs.board[move] = gs.currentPlayer
		return true
	}

	return false
}

func (gs *GameState) NextTurn() {
	gs.currentPlayer = GetOponent(gs.currentPlayer)
}

func (gs *GameState) HasWinner() bool {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if (gs.board[i*3] != 0 && gs.board[i*3] == gs.board[i*3+1] && gs.board[i*3] == gs.board[i*3+2]) ||
			(gs.board[i] != 0 && gs.board[i] == gs.board[i+3] && gs.board[i] == gs.board[i+6]) {
			return true
		}
	}

	// Check diagonals
	return (gs.board[0] != 0 && gs.board[0] == gs.board[4] && gs.board[0] == gs.board[8]) ||
		(gs.board[2] != 0 && gs.board[2] == gs.board[4] && gs.board[2] == gs.board[6])
}

func (gs *GameState) IsDraw() bool {
	for _, item := range gs.board {
		if item == 0 {
			return false
		}
	}

	return true
}

func GetOponent(player int) int {
	if player == XPlayer {
		return OPlayer
	}

	return XPlayer
}

func (gs *GameState) MakeMove() int {
	if gs.difficulty == DifficultyEasy {
		return gs.findRandomMove()
	} else if gs.difficulty == DifficultyMedium {
		return gs.findMediumMove()
	}
	return gs.findBestMove()
}

func (gs *GameState) findRandomMove() int {
	emptyCells := make([]int, 0)
	for i := 0; i < 9; i++ {
		if gs.board[i] == 0 {
			emptyCells = append(emptyCells, i)
		}
	}

	if len(emptyCells) == 0 {
		return -1
	}

	return emptyCells[rand.Intn(len(emptyCells))]
}

func (gs *GameState) findMediumMove() int {
	// Check if there's a move that wins the game
	for i := 0; i < 9; i++ {
		if gs.board[i] == 0 {
			gs.board[i] = gs.player
			if gs.HasWinner() {
				gs.board[i] = 0
				return i
			}
			gs.board[i] = 0
		}
	}

	// Check if there's a move that prevents the opponent from winning
	for i := 0; i < 9; i++ {
		if gs.board[i] == 0 {
			gs.board[i] = GetOponent(gs.player)
			if gs.HasWinner() {
				gs.board[i] = 0
				return i
			}
			gs.board[i] = 0
		}
	}

	// If there's no winning or blocking move, make a random move
	return gs.findRandomMove()
}

func (gs *GameState) findBestMove() int {
	bestScore := math.Inf(-1)
	bestMove := -1

	for i := 0; i < 9; i++ {
		if gs.board[i] == 0 {
			gs.board[i] = gs.player
			score := gs.minimax(0, true, math.Inf(-1), math.Inf(1))
			gs.board[i] = 0
			if score > bestScore {
				bestScore = score
				bestMove = i
			}
		}
	}

	return bestMove
}

func (gs *GameState) minimax(depth int, isMaximizing bool, alpha, beta float64) float64 {
	winner := gs.checkWinner()
	if winner != 0 {
		return gs.score(winner, depth)
	}

	if isMaximizing {
		maxEval := math.Inf(-1)
		for i := 0; i < 9; i++ {
			if gs.board[i] == 0 {
				gs.board[i] = gs.player
				eval := gs.minimax(depth+1, false, alpha, beta)
				gs.board[i] = 0
				maxEval = math.Max(maxEval, eval)
				alpha = math.Max(alpha, eval)
				if beta <= alpha {
					break
				}
			}
		}
		return maxEval
	} else {
		minEval := math.Inf(1)
		for i := 0; i < 9; i++ {
			if gs.board[i] == 0 {
				gs.board[i] = GetOponent(gs.player)
				eval := gs.minimax(depth+1, true, alpha, beta)
				gs.board[i] = 0
				minEval = math.Min(minEval, eval)
				beta = math.Min(beta, eval)
				if beta <= alpha {
					break
				}
			}
		}
		return minEval
	}
}

func (gs *GameState) checkWinner() int {
	winningPositions := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}

	for _, pos := range winningPositions {
		if gs.board[pos[0]] != 0 && gs.board[pos[0]] == gs.board[pos[1]] && gs.board[pos[1]] == gs.board[pos[2]] {
			return gs.board[pos[0]]
		}
	}

	isDraw := true
	for _, cell := range gs.board {
		if cell == 0 {
			isDraw = false
			break
		}
	}

	if isDraw {
		return Draw
	}

	return 0
}

func (gs *GameState) score(winner int, depth int) float64 {
	switch winner {
	case gs.player:
		return 10 - float64(depth)
	case GetOponent(gs.player):
		return float64(depth) - 10
	case Draw:
		return 0
	default:
		return 0
	}
}
