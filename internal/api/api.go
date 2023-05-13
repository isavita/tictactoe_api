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
	MISSING_PLAYER     = "Missing player value. Allowed values are 1 (X) or 2 (O)."
	MISSING_BOARD      = "Missing board value. Must have exactly 9 numbers (0, 1, or 2); 0 (empty), 1 (Player 1), 2 (Player 2)."
	INVALID_PLAYER     = "Invalid player value. Allowed values are 1 (X) or 2 (O)."
	INVALID_DIFFICULTY = "Invalid difficulty: Use 1 (Easy), 2 (Medium), or 3 (Hard). Default is 3 (Hard) if not provided."
	INVALID_TURN       = "It's not the submitted player's turn. Please submit the correct player's move."
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

	if moveRequest.Player == 0 {
		http.Error(w, MISSING_PLAYER, http.StatusBadRequest)
		return
	}

	if moveRequest.Board == [9]int{} {
		// If the board is empty, this is a new game and we do not swap the player
	} else if moveRequest.Player == game.XPlayer {
		moveRequest.Player = game.OPlayer
	} else if moveRequest.Player == game.OPlayer {
		moveRequest.Player = game.XPlayer
	} else {
		http.Error(w, INVALID_PLAYER, http.StatusBadRequest)
		return
	}

	if moveRequest.Board == [9]int{} && moveRequest.Player != game.XPlayer {
		http.Error(w, MISSING_BOARD, http.StatusBadRequest)
		return
	}
	if moveRequest.Player != game.XPlayer && moveRequest.Player != game.OPlayer {
		http.Error(w, INVALID_PLAYER, http.StatusBadRequest)
		return
	}

	currentPlayer, err := getCurrentPlayer(moveRequest.Board)
	if err != nil {
		http.Error(w, INVALID_BOARD, http.StatusBadRequest)
		return
	}

	if currentPlayer != moveRequest.Player {
		http.Error(w, INVALID_TURN, http.StatusBadRequest)
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

	moveResponse := api.game.MakeMove(moveRequest)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(moveResponse)
}

func getCurrentPlayer(board [9]int) (int, error) {
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
