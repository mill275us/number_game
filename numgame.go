package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// Create and seed the generator.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create the board as a square array of variable size 'dimension'
	// Create a colcount array to hold the counts of 1's this should be
	// 1/2 of the dimension rounded up in size
	var dimension = 6
	var summary_dimemsion = int(math.Round(float64(dimension) / 2))

	board := make([][]int, dimension)

	colcount := make([][]int, dimension)
	for j := range colcount {
		colcount[j] = make([]int, summary_dimemsion)
	}

	rowcount := make([][]int, summary_dimemsion)
	for k := range rowcount {
		rowcount[k] = make([]int, dimension)
	}

	// Populate the board with random 1's
	for col := range board {
		board[col] = make([]int, dimension)
		for row := range board[col] {
			board[col][row] = int(math.Round(r.Float64()))
		}
	}

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
	for row := 0; row < dimension; row++ {
		for col := range board {
			rowcount[0][row] += board[col][row]
		}
	}

	/************** Output the board **************/
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


