# Tic Tac Toe Problem

## Problem Statement
Design and implement the classic game of Tic Tac Toe. The program should be able to interact with two players and handle their moves on a 3x3 grid.

### Requirements
- **Initialize Game:** Start with an empty 3x3 grid.
- **Player Moves:** Accept input from two players alternately, where player 1 is 'X' and player 2 is 'O'. Ensure the move is valid (i.e., the cell is unoccupied).
- **Check Win:** After every move, check whether there is a winner. A win is defined as three of a player's marks (either 'X' or 'O') in a row, either horizontally, vertically, or diagonally.
- **Check Draw:** If the board is full and there is no winner, the game ends in a draw.
- **Display Board:** After every move, display the current state of the board.
- **Restart Game:** Allow players to restart the game once it ends.

### Example
```
Player 1 make a move: (0-2),(0-2)0,0
 X |   |
   |   |
   |   |
Player 2 make a move: (0-2),(0-2)0,1
 X | O |
   |   |
   |   |
Player 1 make a move: (0-2),(0-2)1,1
 X | O |
   | X |
   |   |
Player 2 make a move: (0-2),(0-2)0,2
 X | O | O
   | X |
   |   |
Player 1 make a move: (0-2),(0-2)2,2
 X | O | O
   | X |
   |   | X
Winner is Player 1
```

### Problem Extension
1. **AI Opponent:** Implement a simple AI opponent so that a single player can play against the computer.
2. **Scalable Board:** Modify the game to support a scalable board size, e.g., 4x4, 5x5, etc., and adjust the winning condition accordingly.

## Design Pattern
- **Strategy Pattern:** Use this for the player move strategy. This allows the flexibility to switch between different algorithms for player moves, for example, from a human player input to an AI algorithm.
- **Observer Pattern:** This can be used to notify various components of the game (like display or end-game checks) when a move is made.

