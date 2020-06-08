package main

import (
	"fmt"
)

const (
	// N is the size of the square matrix
	N = 9

	// BLANK represents an empty cell
	BLANK = 0
)

func usedInRow(grid [N][N]int, row, number int) bool {
	for col := 0; col < N; col++ {
		if grid[row][col] == number {
			return true
		}
	}

	return false
}

func usedInCol(grid [N][N]int, col, number int) bool {
	for row := 0; row < N; row++ {
		if grid[row][col] == number {
			return true
		}
	}

	return false
}

func usedInBlock(grid [N][N]int, row, col, number int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+row][j+col] == number {
				return true
			}
		}
	}

	return false
}

func isSafe(grid [N][N]int, row, col, number int) bool {
	if !usedInRow(grid, row, number) && !usedInCol(grid, col, number) && !usedInBlock(grid, row-row%3, col-col%3, number) {
		return true
	}

	return false
}

func getUnassignedLocation(grid [N][N]int) (int, int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == BLANK {
				return i, j
			}
		}

	}

	return 9, 9
}

func solveSudoku(grid *[N][N]int) bool {
	row, col := getUnassignedLocation(*grid)

	if row == N && col == N {
		return true
	}

	for number := 1; number <= N; number++ {

		// If placing the current number in the current
		// unassigned location is valid, go ahead
		if isSafe(*grid, row, col, number) {
			grid[row][col] = number

			// Do the same thing again recursively. If we go
			// through all of the recursions, and in the end
			// return true, then all of our number placements
			// on the Soduko grid are valid and we have fully
			// solved it
			if solveSudoku(grid) {
				return true
			}

			// As we were not able to validly go through all
			// of the recursions, we must have an invalid number
			// placement somewhere. Lets go back and try a
			// different number for this particular unassigned location
			grid[row][col] = BLANK
		}
	}

	// If we have gone through all possible numbers for the current unassigned
	// location, then we probably assigned a bad number early. Lets backtrack
	// and try a different number for the previous unassigned locations.
	return false

}

func main() {
	grid := [N][N]int{

		// Hard according to sudoku.com
		// https://sudoku.com/hard/

		// {3, 6, 0, 9, 0, 4, 1, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 7, 4},
		// {0, 0, 4, 0, 0, 3, 0, 6, 0},

		// {9, 0, 5, 0, 1, 0, 8, 0, 7},
		// {0, 0, 0, 7, 5, 0, 0, 0, 9},
		// {0, 0, 0, 0, 0, 0, 0, 3, 0},

		// {7, 0, 0, 0, 0, 8, 0, 9, 0},
		// {0, 0, 3, 5, 4, 0, 0, 0, 8},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0},

		// Expert according to sudoku.com
		// https://sudoku.com/expert/
		{0, 2, 0, 0, 7, 0, 4, 0, 0},
		{7, 4, 0, 0, 6, 0, 0, 0, 0},
		{0, 5, 0, 0, 0, 9, 0, 6, 0},

		{0, 0, 0, 0, 0, 0, 6, 0, 0},
		{0, 0, 6, 0, 0, 0, 0, 3, 1},
		{0, 7, 0, 0, 0, 3, 5, 0, 0},

		{0, 0, 0, 0, 9, 0, 2, 0, 4},
		{1, 0, 0, 0, 0, 5, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 7},
	}

	// Change the grid matrix by passing a reference to it
	if solveSudoku(&grid) == true {
		for i := 0; i < N; i++ {
			fmt.Println(grid[i])
		}
	} else {
		fmt.Println("No solution exists for the given Soduko")
	}

}
