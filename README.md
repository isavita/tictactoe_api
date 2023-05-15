# API for making moves in a game of Tic Tac Toe
This repository contains an API for playing Tic Tac Toe with varying board sizes. The game board is represented as a single array, and the API allows users to submit their board with their move reflected in it and receive the AI's response move. The input and output are both JSON formatted.

## The API has only one endpoint:
`POST /v1/tictactoe`

Play a move in a Tic Tac Toe game. The game board is represented as a single array, and the API allows users to submit their moves and receive the AI's response move. The input and output are both JSON formatted.

Request Input:

To get a move, send a POST request with a JSON object containing the following properties:

- board: The current state of the board as a single array. The array size depends on the chosen board size: 9 for 3x3, 16 for 4x4, 25 for 5x5, or 36 for 6x6. Each element can be one of the following values:
  - 0: Empty cell
  - 1: X
  - 2: O
- boardSize: The size of one side of the board. This value can be 3, 4, 5, or 6.

Example of a valid request:
```json
{
    "board": [0, 0, 0, 1, 0, 0, 0, 0, 0],
    "boardSize": 3
}
```
Response:

The API will return a JSON object with the following properties:

- success: A boolean value indicating if the move was successful.
- message: A text description of the move made by the AI.
- board: The updated game board as an array.
- boardDisplay: The visual representation of the board as a string.
- gameStatus: The current game status. Possible values include:
  - ongoing
  - player1_wins
  - player2_wins
  - draw
- nextPlayer: The next player to make a move (1 for X or 2 for O).

Example of a valid response:
```json
{
    "success": true,
    "message": "Player 2 placed O in position 0.",
    "board": [
        2,
        0,
        0,
        1,
        0,
        0,
        0,
        0,
        0
    ],
    "boardDisplay": " O |   |   \n --------- \n X |   |   \n --------- \n   |   |   ",
    "gameStatus": "ongoing",
    "nextPlayer": 1
}
```
To play the game, send requests with your board and which player turn is to the API and process the responses to get updated state of the game.
