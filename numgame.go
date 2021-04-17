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

	// Output the col counts
	for row := 0; row < summary_dimemsion; row++ {
		output += "\n        "
		for col := 0; col < dimension; col++ {
			output += fmt.Sprintf("  %d  ", colcount[col][0])
		}
	}
	output += "\n\n"

	// Output main playing field and also the row counts at the front
	output += "       "
	for col := 0; col < dimension; col++ {
		output += "+----"
	}
	output += "+ \n"
	for row := 0; row < dimension; row++ {
		output += fmt.Sprintf(" %d --- |", rowcount[0][row])
		for col := 0; col < dimension; col++ {
			output += fmt.Sprintf("  %d |", board[col][row])
		}
		output += "\n       "
		for col := 0; col < dimension; col++ {
			output += "+----"
		}
		output += "+ \n"
	}
	fmt.Println(output)

}
