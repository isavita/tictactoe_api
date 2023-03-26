package main

import (
	"log"
	"net/http"
	"os"

	"github.com/isavita/tictactoe_api/internal/api"
	"github.com/isavita/tictactoe_api/internal/game"
)

func main() {
	ticTacToeGame := game.NewTicTacToeGame()
	ticTacToeAPI := api.NewTicTacToeAPI(ticTacToeGame)

	http.HandleFunc("/v1/tictactoe", ticTacToeAPI.TicTacToeHandler)

	// Handle all other requests with a 404 Not Found.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("404 - %s", r.URL.Path)
		http.NotFound(w, r)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("default to port %s", port)
	}

	http.ListenAndServe(":"+port, nil)
}
