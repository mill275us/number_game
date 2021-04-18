package main

import (
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
}

func main() {

	// Initialize game board and row and col counts in the game struct
	// The int passed to game represent the board size
	game := InitGame(6)

	// Populate the board with random 1's
	FillBoard(game)

	/*
		// Calculate the 1's in rows and colums
		UpdateBoard(dimension, board, colcount, rowcount)

		// Output the board
		DisplayBoard(dimension, summary_dimemsion, board, colcount, rowcount)
	*/

	fmt.Println(game)
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

func UpdateBoard(dimension int, board [][]int, colcount [][]int, rowcount [][]int) {
	// Count number of 1's in each col
	var check int
	var lastWasTrue bool
	var colCountRow int
	for col := range board {
		check = 0
		lastWasTrue = true
		colCountRow = 0
		for row := range board[col] {
			if board[col][row] == 1 || check > 0 {

				if lastWasTrue && board[col][row] == 1 {
					colcount[col][colCountRow] += board[col][row]
					check++
				} else if !lastWasTrue && board[col][row] == 1 {
					lastWasTrue = true
					colCountRow++
					colcount[col][colCountRow] += board[col][row]
				} else if lastWasTrue && board[col][row] == 0 {
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
	for row := 0; row < dimension; row++ {
		isFirstCol = true
		rowCountCol = 0

		for col := range board {
			// 	if it is first col
			// 		if 1 then rowcount++ and  first col = false
			// 		else first col = false
			if isFirstCol {
				if board[col][row] == 1 {
					rowcount[rowCountCol][row]++
				}
				isFirstCol = false
			} else {
				// 	if not first col
				// 		if 1 and last was 1 then rowcount++ and store rowcount in array
				// 		if 0 and last was 1 then arrayindex++ and rowcount = 0
				// 		if 01and last was 0 then rowcount++ and store rowcount in array
				if board[col][row] == 1 && board[col-1][row] == 1 {
					rowcount[rowCountCol][row]++
				}
				if board[col][row] == 0 && board[col-1][row] == 1 {
					rowCountCol++
				}
				if board[col][row] == 1 && board[col-1][row] == 0 {
					rowcount[rowCountCol][row]++
				}
			}
		}
	}
}

func DisplayBoard(dimension int, summary_dimemsion int, board [][]int, colcount [][]int, rowcount [][]int) {
	var output string

	// Output the col counts as header rows
	for row := 0; row < summary_dimemsion; row++ {

		// Indeed based on summary_dimemsion and the widthe
		// of each of the columns
		output += "\n "
		for i := 0; i < summary_dimemsion; i++ {
			output += "    "
		}

		// Output actual colcount values into the head row
		for col := 0; col < dimension; col++ {
			output += fmt.Sprintf("  %d  ", colcount[col][row])
		}
	}
	// Add a divider between header rows and start of board
	output += "\n\n"
	output += PrintSeparator(dimension, summary_dimemsion)

	// Output main playing field and also the row counts at the front
	for row := 0; row < dimension; row++ {

		// Output the row counts as the first elements of the row
		for i := 0; i < summary_dimemsion; i++ {
			output += fmt.Sprintf(" %d ", rowcount[i][row])
		}
		output += " - |"

		// Loop through all of the columns for this row
		for col := 0; col < dimension; col++ {
			output += fmt.Sprintf("  %d |", board[col][row])
		}

		// Add a divider between rows
		output += "\n"
		output += PrintSeparator(dimension, summary_dimemsion)

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
