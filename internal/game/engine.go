package game

import (
	"math"
	"math/rand"
)

const (
	XPlayer        = 1
	OPlayer        = 2
	Draw           = 3
	DifficultyEasy = 101
	DifficultyHard = 102
)

type GameState struct {
	board         [9]int
	currentPlayer int
	player        int
	difficulty    int
}

func (gs *GameState) Play(move int) bool {
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
	}
	return gs.findBestMove()
}

func (gs *GameState) findRandomMove() int {
	for i := 0; i < 100; i++ {
		n := rand.Intn(9)
		if gs.board[n] == 0 {
			return n
		}
	}

	return -1
}

func (gs *GameState) findBestMove() int {
	bestScore := math.Inf(-1)
	bestMove := -1

	for i := 0; i < 9; i++ {
		if gs.board[i] == 0 {
			gs.board[i] = gs.player
			score := gs.minimax(0, false)
			gs.board[i] = 0
			if score > bestScore {
				bestScore = score
				bestMove = i
			}
		}
	}

	return bestMove
}

func (gs *GameState) minimax(depth int, isMaximizing bool) float64 {
	winner := gs.checkWinner()
	if winner != 0 {
		return gs.score(winner, depth)
	}

	if isMaximizing {
		maxEval := math.Inf(-1)
		for i := 0; i < 9; i++ {
			if gs.board[i] == 0 {
				gs.board[i] = gs.player
				eval := gs.minimax(depth+1, false)
				gs.board[i] = 0
				maxEval = math.Max(maxEval, eval)
			}
		}
		return maxEval
	} else {
		minEval := math.Inf(1)
		for i := 0; i < 9; i++ {
			if gs.board[i] == 0 {
				gs.board[i] = GetOponent(gs.player)
				eval := gs.minimax(depth+1, true)
				gs.board[i] = 0
				minEval = math.Min(minEval, eval)
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
