package game

import (
	"testing"
)

func TestFindBestMove(t *testing.T) {
	gs := GameState{
		board:      []int{2, 2, 0, 0, 1, 0, 0, 1, 1},
		boardSize:  3,
		player:     OPlayer,
		difficulty: DifficultyHard,
	}

	expectedMove := 2
	actualMove := gs.findBestMove()

	if actualMove != expectedMove {
		t.Errorf("Expected move at index %d, but got %d", expectedMove, actualMove)
	}
}

func TestFindBestMove4By4Board(t *testing.T) {
	gs := GameState{
		board:      []int{2, 2, 2, 0, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0},
		boardSize:  4,
		player:     OPlayer,
		difficulty: DifficultyHard,
	}

	expectedMove := 3
	actualMove := gs.findBestMove()

	if actualMove != expectedMove {
		t.Errorf("Expected move at index %d, but got %d", expectedMove, actualMove)
	}
}

func TestFindBestMove5By5Board(t *testing.T) {
	gs := GameState{
		board:      []int{2, 2, 2, 2, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		boardSize:  5,
		player:     OPlayer,
		difficulty: DifficultyHard,
	}

	expectedMove := 4
	actualMove := gs.findBestMove()

	if actualMove != expectedMove {
		t.Errorf("Expected move at index %d, but got %d", expectedMove, actualMove)
	}
}

func TestFindBestMove6By6Board(t *testing.T) {
	gs := GameState{
		board:      []int{2, 2, 2, 2, 2, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		boardSize:  6,
		player:     OPlayer,
		difficulty: DifficultyHard,
	}

	expectedMove := 5
	actualMove := gs.findBestMove()

	if actualMove != expectedMove {
		t.Errorf("Expected move at index %d, but got %d", expectedMove, actualMove)
	}
}
