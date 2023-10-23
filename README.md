# Minesweeper Game in Go

This Minesweeper game is implemented in Go, utilizing the ebiten, a 2D game engine, and ebitenui, a user interface engine, for rendering graphics and user interface elements. The code is structured into three main packages: `model`, `view`, and `controller`, each responsible for different aspects of the game.

## Table of Contents

- [Code Structure](#code-structure)
  - [Model Package](#model-package)
  - [View Package](#view-package)
  - [Controller Package](#controller-package)
- [How to Build and Run](#how-to-build-and-run)

## Code Structure

### Model Package

Contains the data structures (2D array) and logic underlying the game. This includes the definitions of the game board, cells, and game state.

### View Package

Contains all of the code related to rendering the game to the screen. This includes the definitions of the game window, board, cells, and header user interface elements.

### Controller Package

The controller package acts as a bridge between the model and view packages. It contains functions that handle user interactions and update the game state accordingly.

## How to Build and Run

Minesweeper is a classic puzzle game. The objective is to clear a board containing hidden "mines" without detonating any of them, with help from clues about the number of neighboring mines in each field. The game is won when all safe cells are revealed, and lost when a mine is revealed.

1. [Install Go](https://go.dev/dl/), if not already installed.
2. Clone or download this repository.
3. Navigate to the directory containing your Go files in a terminal.
4. Run the command `go build .` to build the executable file for the game.

   - A file named `Minesweeper.exe` will be created in the same directory.

5. Running the executeable

   - Through TerminaL:

     - On Windows:

       ```sh
       .\Mineweeper.exe
       ```

     - On Linux/macOS:

       ```sh
       ./Minesweeper
       ```

   - Alternatively, you can double-click the executable file to run it in File Explorer/ Finder depening on your OS.
