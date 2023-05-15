package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/isavita/tictactoe_api/internal/game"
	"github.com/isavita/tictactoe_api/internal/model"
)

type TicTacToeAPI struct {
	game *game.TicTacToeGame
}

func NewTicTacToeAPI(game *game.TicTacToeGame) *TicTacToeAPI {
	return &TicTacToeAPI{
		game: game,
	}
}

const (
	MISSING_BOARD      = "Missing board value. Must have exactly 9 (3x3) or 16 (4x4) or 25 (5x5) or 36 (6x6) numbers (0, 1, or 2); 0 (empty), 1 (Player 1), 2 (Player 2)."
	INVALID_DIFFICULTY = "Invalid difficulty: Use 1 (Easy), 2 (Medium), or 3 (Hard). Default is 3 (Hard) if not provided."
	INVALID_BOARD_SIZE = "The supported boardSize values are 3, 4, 5 and 6."
	INVALID_BOARD      = "Invalid board: Must have exactly 9 numbers (0, 1, or 2); 0 (empty), 1 (Player 1), 2 (Player 2); Player 1 moves >= Player 2 moves; max difference: 1."
)

func (api *TicTacToeAPI) TicTacToeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var moveRequest model.MoveRequest
	err := json.NewDecoder(r.Body).Decode(&moveRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Sets default value to 3 for 3x3 board
	if moveRequest.BoardSize <= 0 {
		moveRequest.BoardSize = 3
	}

	// Check the board is not too big
	if moveRequest.BoardSize > 6 {
		http.Error(w, INVALID_BOARD_SIZE, http.StatusBadRequest)
		return
	}

	// if the board is not initialized
	if moveRequest.Board == nil {
		moveRequest.Board = make([]int, moveRequest.BoardSize*moveRequest.BoardSize)
	}

	currentPlayer, err := getCurrentPlayer(moveRequest.Board)
	if err != nil {
		http.Error(w, INVALID_BOARD, http.StatusBadRequest)
		return
	}

	if moveRequest.Difficulty == 0 || moveRequest.Difficulty == 3 {
		moveRequest.Difficulty = game.DifficultyHard
	} else if moveRequest.Difficulty == 1 {
		moveRequest.Difficulty = game.DifficultyEasy
	} else if moveRequest.Difficulty == 2 {
		moveRequest.Difficulty = game.DifficultyMedium
	} else {
		http.Error(w, INVALID_DIFFICULTY, http.StatusBadRequest)
		return
	}

	moveResponse := api.game.MakeMove(currentPlayer, moveRequest)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(moveResponse)
}

func getCurrentPlayer(board []int) (int, error) {
	xCount := 0
	oCount := 0

	for _, val := range board {
		if val == game.XPlayer {
			xCount++
		} else if val == game.OPlayer {
			oCount++
		}
	}

	if xCount < oCount || xCount > oCount+1 {
		return 0, errors.New("invalid board")
	}

	if xCount == oCount {
		return game.XPlayer, nil
	}

	return game.OPlayer, nil
}
