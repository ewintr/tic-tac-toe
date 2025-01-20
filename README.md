# Tic-Tac-Toe

A simple implementation of Tic-Tac-Toe that runs in the terminal.

For information about this game, check [Wikipedia](https://en.wikipedia.org/wiki/Tic-tac-toe).

Run the game with:

```bash
go run .
```

It uses [escape codes](https://en.wikipedia.org/wiki/ANSI_escape_code) to clear the screen after each turn. If this is not supported by your terminal, or if you just don't like them, use the following command to run the game without them:

```bash
go run . plain
```

## Bot Player

### Simple

- Implement `Player` interface
- Select a random move


### Complex

Implement the [Minimax](https://en.wikipedia.org/wiki/Minimax) algorithm:

- Implement a way to score the board
- Recursively play all possible moves
  - Track score for each board state
- Score each play sequence on minimum and maximum score
  - Use maximum score for states where bot is going to move
  - Use minimum score fot states where player is going to move
- Pick best option

Example: https://www.neverstopbuilding.com/blog/minimax


