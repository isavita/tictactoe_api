package main

import (
	"io/ioutil"
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

	// Handle ai-plugin.json request for OpenAI Plugins.
	http.HandleFunc("/.well-known/ai-plugin.json", openAIPluginHandler)

	// Handle openapi.yaml request.
	http.HandleFunc("/openapi.yaml", openapiHandler)

	// Handle logo.png request.
	http.HandleFunc("/logo.png", logoHandler)

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

func openAIPluginHandler(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("./.well-known/ai-plugin.json") // replace with the path to your ai-plugin.json file
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func logoHandler(w http.ResponseWriter, r *http.Request) {
	logoFile, err := os.Open("./.well-known/logo.png")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer logoFile.Close()

	logoData, err := ioutil.ReadAll(logoFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(logoData)
}

func openapiHandler(w http.ResponseWriter, r *http.Request) {
	openapiFile, err := os.Open("./.well-known/openapi.yaml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer openapiFile.Close()

	openapiData, err := ioutil.ReadAll(openapiFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/yaml")
	w.Write(openapiData)
}
