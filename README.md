# Minesweeper Game

This is a Minesweeper game written in Golang, using the Fyne library for the graphical user interface.

## Architecture
The game is architected using the layered architecture, using the main layers Model and View, with additional packages for the command-line interface (CLI) and build configuration.

### Model
The model package contains the game logic for Minesweeper. This includes the game board, mines, and methods for updating the board as the player makes moves.

### View
The view package contains the Fyne GUI code for rendering the game window and board. It also handles player input, such as mouse clicks, and updates the game model as necessary.

### CLI
The cmd package contains a command-line interface for running the game. This can be used as an alternative to the Fyne GUI, or for automated testing.

### Build
The build package contains configuration files and scripts for building and packaging the game. This includes building binaries for different operating systems and architectures, and generating installation packages.

### Getting Started
To run the game using the Fyne GUI, you need to have Golang and Fyne installed on your system. You can follow the installation instructions for Golang from the official website, and for Fyne from the Fyne website.

Once you have installed Golang and Fyne, you can download the source code for the Minesweeper game from this repository. To build and run the game, navigate to the directory where the source code is located and run the following command:

```
go run cmd/minesweeper/main.go
```
This will launch the game window, and you can start playing Minesweeper.

To run the game using the command-line interface, navigate to the directory where the source code is located and run the following command:

```
go run cmd/minesweeper-cli/main.go
```
This will start the game in the terminal, and you can play using keyboard commands.

## How to Play
The goal of the game is to clear the minefield without detonating any mines. You can do this by revealing squares on the field, and if there is a mine, the game is over. Otherwise, the square will show a number indicating how many mines are adjacent to that square.

To reveal a square, simply click on it with the left mouse button. If you think there is a mine in a square, you can flag it with the right mouse button. To remove a flag, simply right-click on the square again.

The game is won when all non-mine squares are revealed, and lost when a mine is detonated. You can start a new game at any time by clicking on the "New Game" button.

## Contributing
If you would like to contribute to the Minesweeper game, you are welcome to fork this repository and submit a pull request with your changes. Please make sure to follow the contribution guidelines when submitting your pull request.

## License
This project is licensed under the MIT License - see the [LICENSE]() file for details.