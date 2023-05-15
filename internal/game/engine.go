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
	MaxDepth         = 6 // change this to adjust the search depth
)

type GameState struct {
	board         []int
	boardSize     int
	currentPlayer int
	player        int
	difficulty    int
}

func (gs *GameState) Play(move int) bool {
	if move < 0 || move >= gs.boardSize*gs.boardSize {
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
	size := gs.boardSize
	// Check rows
	for i := 0; i < size; i++ {
		win := true
		for j := 1; j < size; j++ {
			if gs.board[i*size] == 0 || gs.board[i*size] != gs.board[i*size+j] {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	// Check columns
	for i := 0; i < size; i++ {
		win := true
		for j := 1; j < size; j++ {
			if gs.board[i] == 0 || gs.board[i] != gs.board[i+j*size] {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	// Check diagonals
	win := true
	// Check main diagonal (top-left to bottom-right)
	for i := 1; i < size; i++ {
		if gs.board[0] == 0 || gs.board[0] != gs.board[i*size+i] {
			win = false
			break
		}
	}
	if win {
		return true
	}

	win = true
	// Check secondary diagonal (top-right to bottom-left)
	for i := 1; i < size; i++ {
		if gs.board[size-1] == 0 || gs.board[size-1] != gs.board[i*size+(size-i-1)] {
			win = false
			break
		}
	}

	return win
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
	size := gs.boardSize * gs.boardSize
	for i := 0; i < size; i++ {
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
	for i := 0; i < gs.boardSize*gs.boardSize; i++ {
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
	for i := 0; i < gs.boardSize*gs.boardSize; i++ {
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

	for i := 0; i < gs.boardSize*gs.boardSize; i++ {
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
	if gs.boardSize > 3 && depth == MaxDepth {
		return gs.heuristic()
	}
	winner := gs.checkWinner()
	if winner != 0 {
		return gs.score(winner, depth)
	}

	if isMaximizing {
		maxEval := math.Inf(-1)
		for i := 0; i < gs.boardSize*gs.boardSize; i++ {
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

		for i := 0; i < gs.boardSize*gs.boardSize; i++ {
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
	size := gs.boardSize

	// Check rows
	for i := 0; i < size; i++ {
		win := true
		for j := 1; j < size; j++ {
			if gs.board[i*size] == 0 || gs.board[i*size] != gs.board[i*size+j] {
				win = false
				break
			}
		}
		if win {
			return gs.board[i*size]
		}
	}

	// Check columns
	for i := 0; i < size; i++ {
		win := true
		for j := 1; j < size; j++ {
			if gs.board[i] == 0 || gs.board[i] != gs.board[i+j*size] {
				win = false
				break
			}
		}
		if win {
			return gs.board[i]
		}
	}

	// Check diagonals
	win := true
	// Check main diagonal (top-left to bottom-right)
	for i := 1; i < size; i++ {
		if gs.board[0] == 0 || gs.board[0] != gs.board[i*size+i] {
			win = false
			break
		}
	}
	if win {
		return gs.board[0]
	}

	win = true
	// Check secondary diagonal (top-right to bottom-left)
	for i := 1; i < size; i++ {
		if gs.board[size-1] == 0 || gs.board[size-1] != gs.board[i*size+(size-i-1)] {
			win = false
			break
		}
	}
	if win {
		return gs.board[size-1]
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
		return 100 - float64(depth)
	case GetOponent(gs.player):
		return float64(depth) - 100
	case Draw:
		return 0
	default:
		return 0
	}
}

func (gs *GameState) heuristic() float64 {
	aiPotentialWins := gs.countPotentialWins(gs.player)
	opponentPotentialWins := gs.countPotentialWins(GetOponent(gs.player))

	return float64(aiPotentialWins - opponentPotentialWins)
}

func (gs *GameState) countPotentialWins(player int) int {
	count := 0
	size := gs.boardSize

	// Check rows
	for i := 0; i < size; i++ {
		potentialWin := true
		emptyCount := 0
		for j := 0; j < size; j++ {
			if gs.board[i*size+j] == GetOponent(player) {
				potentialWin = false
				break
			}
			if gs.board[i*size+j] == 0 {
				emptyCount++
			}
		}
		if potentialWin && emptyCount == 1 {
			count++
		}
	}

	// Check columns
	for i := 0; i < size; i++ {
		potentialWin := true
		emptyCount := 0
		for j := 0; j < size; j++ {
			if gs.board[i+j*size] == GetOponent(player) {
				potentialWin = false
				break
			}
			if gs.board[i+j*size] == 0 {
				emptyCount++
			}
		}
		if potentialWin && emptyCount == 1 {
			count++
		}
	}

	// Check main diagonal
	potentialWin := true
	emptyCount := 0
	for i := 0; i < size; i++ {
		if gs.board[i*size+i] == GetOponent(player) {
			potentialWin = false
			break
		}
		if gs.board[i*size+i] == 0 {
			emptyCount++
		}
	}
	if potentialWin && emptyCount == 1 {
		count++
	}

	// Check secondary diagonal
	potentialWin = true
	emptyCount = 0
	for i := 0; i < size; i++ {
		if gs.board[i*size+(size-i-1)] == GetOponent(player) {
			potentialWin = false
			break
		}
		if gs.board[i*size+(size-i-1)] == 0 {
			emptyCount++
		}
	}
	if potentialWin && emptyCount == 1 {
		count++
	}

	return count
}
