package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type game struct {
	dimension         int
	summary_dimemsion int
	board             [][]int
	colcount          [][]int
	rowcount          [][]int
	modePtr           *int
}

func main() {

	// Harvest command line arguments
	// size -- int representing the square size of the board Y by Y
	//		-- default size is 5
	// mode -- int representing the display mode
	//		-- 1: (Default) Hides the game values
	//		-- 0: Dispalys values for debugging
	sizePtr := flag.Int("size", 5, "Size of game board as an Int")
	modePtr := flag.Int("mode", 1, "Toggle board reveal for debugging as an Int")

	// After all flags are declared you must call Parse
	flag.Parse()

	// Initialize game board and row and col counts in the game struct
	// The int passed to game represent the board size
	game := InitGame(*sizePtr)
	game.modePtr = modePtr

	// Populate the board with random 1's
	FillBoard(game)

	// Calculate the 1's in rows and colums
	UpdateBoard(game)

	// Output the board
	DisplayBoard(game)

	// Temp code for playing
	// Wait for user to press any key then reveal the answer
	// Set modePtr to 0 to change to answer reveal
	var dummyString string
	fmt.Println("Press Any Key for Answer")
	fmt.Scanln(&dummyString)
	*game.modePtr = 0 // Set mode to 0 to change to answer reveal
	DisplayBoard(game)

}

// InitGame takes an integer argument for the size of the playing field
// returns a game struct
// Create the board as a square array of variable size 'dimension'
// Create a colcount and rowcount array to hold the counts of 1's this should be
// 1/2 of the dimension rounded up in size
func InitGame(dimension int) *game {
	game := game{dimension: dimension}
	game.summary_dimemsion = int(math.Round(float64(game.dimension) / 2))

	game.board = make([][]int, game.dimension)

	game.colcount = make([][]int, game.dimension)
	for j := range game.colcount {
		game.colcount[j] = make([]int, game.summary_dimemsion)
	}

	game.rowcount = make([][]int, game.summary_dimemsion)
	for k := range game.rowcount {
		game.rowcount[k] = make([]int, game.dimension)
	}
	return &game
}

// Populate the board with random 1's
// Takes a pointer to a game struct whose data will be modified
// This function further builds out the arrays for the board as it populates them
func FillBoard(gameStruct *game) {
	// Create and seed the generator.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for col := range gameStruct.board {
		gameStruct.board[col] = make([]int, gameStruct.dimension)
		for row := range gameStruct.board[col] {
			gameStruct.board[col][row] = int(math.Round(r.Float64()))
		}
	}
}

func UpdateBoard(gameStruct *game) {
	// Count number of 1's in each col
	var check int
	var lastWasTrue bool
	var colCountRow int
	for col := range gameStruct.board {
		check = 0
		lastWasTrue = true
		colCountRow = 0
		for row := range gameStruct.board[col] {
			if gameStruct.board[col][row] == 1 || check > 0 {

				if lastWasTrue && gameStruct.board[col][row] == 1 {
					gameStruct.colcount[col][colCountRow] += gameStruct.board[col][row]
					check++
				} else if !lastWasTrue && gameStruct.board[col][row] == 1 {
					lastWasTrue = true
					colCountRow++
					gameStruct.colcount[col][colCountRow] += gameStruct.board[col][row]
				} else if lastWasTrue && gameStruct.board[col][row] == 0 {
					lastWasTrue = false
				}
			}
		}
	}

	// Count number of 1's in each row
	// first col -- boolean whether this is the first col looked at
	// rowCountCol -- index for the col of the header array for storing rowcount
	var isFirstCol bool
	var rowCountCol int

	// Loop logic
	for row := 0; row < gameStruct.dimension; row++ {
		isFirstCol = true
		rowCountCol = 0

		for col := range gameStruct.board {
			// 	if it is first col
			// 		if 1 then rowcount++ and  first col = false
			// 		else first col = false
			if isFirstCol {
				if gameStruct.board[col][row] == 1 {
					gameStruct.rowcount[rowCountCol][row]++
				}
				isFirstCol = false
			} else {
				// 	if not first col
				// 		if 1 and last was 1 then rowcount++ and store rowcount in array
				// 		if 0 and last was 1 then arrayindex++ and rowcount = 0
				// 		if 01and last was 0 then rowcount++ and store rowcount in array
				if gameStruct.board[col][row] == 1 && gameStruct.board[col-1][row] == 1 {
					gameStruct.rowcount[rowCountCol][row]++
				}
				if gameStruct.board[col][row] == 0 && gameStruct.board[col-1][row] == 1 {
					rowCountCol++
				}
				if gameStruct.board[col][row] == 1 && gameStruct.board[col-1][row] == 0 {
					gameStruct.rowcount[rowCountCol][row]++
				}
			}
		}
	}
}

func DisplayBoard(gameStruct *game) {
	var output string

	// Output the col counts as header rows
	for row := 0; row < gameStruct.summary_dimemsion; row++ {

		// Indeed based on summary_dimemsion and the widthe
		// of each of the columns
		output += "\n "
		for i := 0; i < gameStruct.summary_dimemsion; i++ {
			output += "    "
		}

		// Output actual colcount values into the head row
		for col := 0; col < gameStruct.dimension; col++ {
			output += fmt.Sprintf("  %d  ", gameStruct.colcount[col][row])
		}
	}
	// Add a divider between header rows and start of board
	output += "\n\n"
	output += PrintSeparator(gameStruct.dimension, gameStruct.summary_dimemsion)

	// Output main playing field and also the row counts at the front
	for row := 0; row < gameStruct.dimension; row++ {

		// Output the row counts as the first elements of the row
		for i := 0; i < gameStruct.summary_dimemsion; i++ {
			output += fmt.Sprintf(" %d ", gameStruct.rowcount[i][row])
		}
		output += " - |"

		// Loop through all of the columns for this row
		// modePtr -- int representing the display mode
		//		-- 1: (Default) Hides the game values
		//		-- 0: Dispalys values for debugging
		// 				If it is a 0 output a blank otherwise print a 1
		for col := 0; col < gameStruct.dimension; col++ {
			if *gameStruct.modePtr == 0 {
				if gameStruct.board[col][row] == 0 {
					output += "    |"
				} else {
					output += fmt.Sprintf("  %d |", gameStruct.board[col][row])
				}
			} else {
				output += "    |"
			}
		}

		// Add a divider between rows
		output += "\n"
		output += PrintSeparator(gameStruct.dimension, gameStruct.summary_dimemsion)

	}
	fmt.Println(output)
}

// PrintSeparator functions returns a string containing row separator
// This string is indented based on the numbers of row count rows
// as defined by summary_dimemsion
// dimension === the Y by Y size of the game board
// summary_dimemsion === 1/2 of dimension rounded up
func PrintSeparator(dimension int, summary_dimemsion int) string {
	var rtnString = "   "
	for i := 0; i < summary_dimemsion; i++ {
		rtnString += "   "
	}

	for col := 0; col < dimension; col++ {
		rtnString += "+----"
	}
	rtnString += "+ \n"
	return rtnString
}
