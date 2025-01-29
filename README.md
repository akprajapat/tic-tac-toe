# Tic-Tac-Toe in Go

A simple command-line **Tic-Tac-Toe** game implemented in **Go (Golang)**.

## Features
- Two-player mode (Player X and Player O)
- Command-line interface (CLI)
- Simple game logic with a 3x3 board
- Detects wins, draws, and invalid moves

## Prerequisites
- **Go 1.18+** installed on your system

## Installation
Clone the repository:
```sh
git clone https://github.com/akprajapat/tic-tac-toe-go.git
cd tic-tac-toe-go
```

## Running the Game
Run the game using:
```sh
go run main.go
```

## How to Play
1. The game starts with an empty 3x3 board.
2. Players take turns entering their move as a row and column number (e.g., `1 2` for row 1, column 2).
3. The game checks for a winner after each move.
4. If a player wins or the game ends in a draw, the result is displayed.
5. Players can restart or exit the game.

## Example Gameplay
```
Welcome to Tic-Tac-Toe!

  ___|___|___
  ___|___|___
     |   |

Player 1, enter your move (row and column): 1 1

  ___|___|___
  ___|_X_|___
     |   |

Player 2, enter your move (row and column): 1 2
...
```

## Future Enhancements
- Implement AI for single-player mode
- Add a GUI version
- Save game histor

## Author
[ARPIT PRAJAPAT](https://github.com/akprajapat)

