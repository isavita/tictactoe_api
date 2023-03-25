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

	if moveRequest.Player != 1 && moveRequest.Player != 2 {
		http.Error(w, "Invalid player value", http.StatusBadRequest)
		return
	}

	currentPlayer, err := getCurrentPlayer(moveRequest.Board)
	if err != nil {
		http.Error(w, "Invalid board value", http.StatusBadRequest)
		return
	}

	if currentPlayer != moveRequest.Player {
		http.Error(w, "It is not your turn", http.StatusBadRequest)
		return
	}

	if moveRequest.Difficulty == 0 {
		moveRequest.Difficulty = game.DifficultyHard
	} else if moveRequest.Difficulty != game.DifficultyEasy && moveRequest.Difficulty != game.DifficultyHard {
		http.Error(w, "Invalid difficulty value", http.StatusBadRequest)
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
