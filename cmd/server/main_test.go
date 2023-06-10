package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/isavita/tictactoe_api/internal/api"
	"github.com/isavita/tictactoe_api/internal/game"
	"github.com/isavita/tictactoe_api/internal/model"
)

func TestMain(m *testing.M) {
	// path setup to the root directory
	// we need this because the main file is in `/cmd/server`
	prevPath, _ := os.Getwd()
	os.Chdir("../..")
	defer os.Chdir(prevPath)

	m.Run()
}

func TestTicTacToeHandler(t *testing.T) {
	t.Parallel()

	t.Run("first move with empty payload", func(t *testing.T) {
		// setup the tic-tac-toe api
		ticTacToeGame := game.NewTicTacToeGame()
		ticTacToeAPI := api.NewTicTacToeAPI(ticTacToeGame)

		s := httptest.NewServer(http.HandlerFunc(ticTacToeAPI.TicTacToeHandler))
		url := s.URL + "/v1/tictactoe"
		payload := strings.NewReader(`{}`)
		req, err := http.NewRequest(http.MethodPost, url, payload)
		assertNoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		assertNoError(t, err)

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assertNoError(t, err)
		assertStatusCode(t, resp, http.StatusOK)

		got := model.MoveResponse{}
		err = json.Unmarshal(body, &got)
		assertNoError(t, err)

		want := model.MoveResponse{
			Success:      true,
			Message:      "Player 1 has placed 'X' in position 1. Please note: If the user is playing with 'X', disregard this move. Instead, ask the user where they would like to place their first move, then present the game board reflecting their choice.",
			Board:        []int{1, 0, 0, 0, 0, 0, 0, 0, 0},
			BoardSize:    3,
			BoardDisplay: " X | 2 | 3 \n --------- \n 4 | 5 | 6 \n --------- \n 7 | 8 | 9 ",
			GameStatus:   "ongoing",
			NextPlayer:   game.OPlayer,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestOpenAIPluginHandler(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/.well-known/ai-plugin.json", nil)
	recorder := httptest.NewRecorder()
	openAIPluginHandler(recorder, req)

	res := recorder.Result()
	defer res.Body.Close()

	assertStatusCode(t, res, http.StatusOK)
}

func TestOpenapiHandler(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/openapi.yaml", nil)
	recorder := httptest.NewRecorder()
	openapiHandler(recorder, req)

	res := recorder.Result()
	defer res.Body.Close()

	assertStatusCode(t, res, http.StatusOK)
}

func TestLogoHandler(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/logo.png", nil)
	recorder := httptest.NewRecorder()
	openapiHandler(recorder, req)
	res := recorder.Result()
	defer res.Body.Close()

	assertStatusCode(t, res, http.StatusOK)
}

func BenchmarkTicTacToeHandler(b *testing.B) {
	ticTacToeGame := game.NewTicTacToeGame()
	ticTacToeAPI := api.NewTicTacToeAPI(ticTacToeGame)
	s := httptest.NewServer(http.HandlerFunc(ticTacToeAPI.TicTacToeHandler))
	defer s.Close()
	url := s.URL + "/openapi.yaml"

	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		http.DefaultClient.Do(req)
	}
}

func assertStatusCode(t testing.TB, res *http.Response, want int) {
	if res.StatusCode != want {
		t.Errorf("got %v want %v", res.StatusCode, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got error %v when no error was expected", err)
	}
}
