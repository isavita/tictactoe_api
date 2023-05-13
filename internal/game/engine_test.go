package game

import (
	"testing"
)

func TestFindBestMove(t *testing.T) {
	gs := GameState{
		board:      [9]int{2, 2, 0, 0, 1, 0, 0, 1, 1},
		player:     OPlayer,
		difficulty: DifficultyHard,
	}

	expectedMove := 2
	actualMove := gs.findBestMove()

	if actualMove != expectedMove {
		t.Errorf("Expected move at index %d, but got %d", expectedMove, actualMove)
	}
}
