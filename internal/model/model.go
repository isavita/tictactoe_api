package model

type MoveRequest struct {
	Board      []int `json:"board,omitempty"`
	BoardSize  int   `json:"boardSize,omitempty"`
	Difficulty int   `json:"difficulty,omitempty"`
}

type MoveResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	Board        []int  `json:"board"`
	BoardDisplay string `json:"boardDisplay"`
	GameStatus   string `json:"gameStatus"`
	NextPlayer   int    `json:"nextPlayer"`
}

const (
	GameStatusOngoing     = "ongoing"
	GameStatusDraw        = "draw"
	GameStatusPlayer1Wins = "player1_wins"
	GameStatusPlayer2Wins = "player2_wins"
)
