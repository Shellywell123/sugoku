package main

import (
	"fmt"
	"os"
	"slices"
)

// Returns unique items in a slice
func Unique(slice []int) []int {
	// create a map with all the values as key
	uniqMap := make(map[int]struct{})
	for _, v := range slice {
		if v != 0 {
			uniqMap[v] = struct{}{}
		}
	}

	// turn the map keys into a slice
	uniqSlice := make([]int, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice

}

// get the list of possible numbers for an empty cell
func (s Sudoku) GetOptions(x int, y int) []int {

	row := s.GetRow(y)
	col := s.GetColumn(x)
	sqa := s.GetSquare(x, y)

	return Unique(slices.Concat(nil, row, col, sqa))
}

// solve a cell by getting the unique set of numbers across row col square
func (s Sudoku) SolveUnique(x int, y int) {

	options := s.GetOptions(x, y)

	if len(options) == 8 {
		for n := 1; n <= 9; n++ {

			found := false
			for _, m := range options {
				if m == n {
					found = true
				}
			}

			if !found {
				s.SetCell(x, y, n)
				break
			}
		}
	}
}

func FindElementIndexesInSlice(slice []int, element int) []int {
	Indexes := []int{}
	for i, e := range slice {
		if e == element {
			Indexes = append(Indexes, i)
		}
	}
	return Indexes
}

// complete a row by checking all other positions are blocked
func (s Sudoku) SolveRowByBlocked(y int) {
	row := s.GetRow(y)
	// get other possible positions in row
	availableRowPositions := FindElementIndexesInSlice(row, 0)

	choices := []int{}
	for i := 1; i <= 9; i++ {
		if !slices.Contains(Unique(row), i) {
			choices = append(choices, i)
		}
	}

	for _, choice := range choices {
		unblockedPositions := []int{}
		positionsBlocked := 0
		for _, xPosition := range availableRowPositions {
			positionBlocked := false

			options := s.GetOptions(xPosition, y)

			if slices.Contains(options, choice) {
				positionBlocked = true
			}

			if positionBlocked {
				positionsBlocked++
			} else {
				unblockedPositions = append(unblockedPositions, xPosition)
			}
		}

		if len(unblockedPositions) == 1 {
			s.SetCell(unblockedPositions[0], y, choice)
			break
		}
	}
}

// complete a column by checking all other positions are blocked
func (s Sudoku) SolveColumnByBlocked(x int) {
	col := s.GetColumn(x)
	// get other possible positions in column
	availableRowPositions := FindElementIndexesInSlice(col, 0)
	unblockedPositions := []int{}

	choices := []int{}
	for i := 1; i <= 9; i++ {
		if !slices.Contains(Unique(col), i) {
			choices = append(choices, i)
		}
	}

	for _, choice := range choices {
		positionsBlocked := 0
		for _, yPosition := range availableRowPositions {
			positionBlocked := false

			options := s.GetOptions(x, yPosition)

			if slices.Contains(options, choice) {
				positionBlocked = true
			}

			if positionBlocked {
				positionsBlocked++
			} else {
				unblockedPositions = append(unblockedPositions, yPosition)
			}
		}

		if len(unblockedPositions) == 1 {
			s.SetCell(x, unblockedPositions[0], choice)
			break
		}
	}
}

// complete a square by checking all other positions are blocked
func (s Sudoku) SolveSquareByBlocked(x int, y int) {
	row := s.GetSquare(x, y)

	availableSquarePositions := [][]int{}

	// get other possible positions in square
	for ey, row := range [][]int{
		s.GetSquare(x, y)[0:3],
		s.GetSquare(x, y)[3:6],
		s.GetSquare(x, y)[6:9],
	} {
		for ex, e := range row {
			// i do not like how this is written
			sx := 0
			sy := 0

			if 0 <= x && x <= 2 {
				sx = 0
			}
			if 3 <= x && x <= 5 {
				sx = 3
			}
			if 6 <= x && x <= 8 {
				sx = 6
			}

			if 0 <= y && y <= 2 {
				sy = 0
			}
			if 3 <= y && y <= 5 {
				sy = 3
			}
			if 6 <= y && y <= 8 {
				sy = 6
			}

			if e == 0 {
				availableSquarePositions = append(availableSquarePositions, []int{sx + ex, sy + ey})
			}
		}
	}

	choices := []int{}
	for i := 1; i <= 9; i++ {
		if !slices.Contains(Unique(row), i) {
			choices = append(choices, i)
		}
	}

	for _, choice := range choices {
		unblockedPositions := [][]int{}
		positionsBlocked := 0
		for _, position := range availableSquarePositions {
			positionBlocked := false

			options := s.GetOptions(position[0], position[1])

			if slices.Contains(options, choice) {
				positionBlocked = true
			}

			if positionBlocked {
				positionsBlocked++
			} else {
				unblockedPositions = append(unblockedPositions, []int{position[0], position[1]})
			}
		}

		if len(unblockedPositions) == 1 {
			s.SetCell(unblockedPositions[0][0], unblockedPositions[0][1], choice)
			// fmt.Println("inputtin", choice, "into", unblockedPositions[0][0], unblockedPositions[0][1])
			break
		}
	}
}

func duplicateDetect(slice []int) bool {
	elements := []int{}
	for _, e := range slice {
		if e == 0 {
			continue
		}

		if slices.Contains(elements, e) {
			fmt.Println(e)
			return true
		} else {
			elements = append(elements, e)
		}
	}
	return false
}

// check every box, row, col contains no duplicate numbers
func (s Sudoku) Validate() {

	for x := 0; x < 9; x++ {
		if duplicateDetect(s.GetColumn(x)) {
			PrintSudoku(s)
			fmt.Println("duplicate detected in column", x)
			os.Exit(0)
		}
	}

	for y := 0; y < 9; y++ {
		if duplicateDetect(s.GetRow(y)) {
			PrintSudoku(s)
			fmt.Println("duplicate detected in row", y)
			os.Exit(0)
		}
	}

	for x := 0; x < 9; x += 3 {
		for y := 0; y < 9; y += 3 {
			if duplicateDetect(s.GetSquare(x, y)) {
				PrintSudoku(s)
				fmt.Println("duplicate detected in row", y)
				os.Exit(0)
			}
		}
	}
}

// main solve function
func (s Sudoku) Solve() {

	cBefore := s.GetCompleted()

	// loop through all cells in the 9x9 grid
	for y := 0; y < 9; y++ {

		// if row complete skip
		if len(Unique(s.GetRow(y))) == 9 {
			continue
		}

		s.SolveRowByBlocked(y)
	}

	// loop through all cells in the 9x9 grid
	for y := 0; y < 9; y++ {

		for x := 0; x < 9; x++ {

			// if cell populated skip
			if s.GetCell(x, y) != 0 {
				continue
			}

			s.SolveUnique(x, y)
			// PrintSudoku(s)
			s.SolveSquareByBlocked(x, y)
		}
	}

	for x := 0; x < 9; x++ {

		// if col complete skip
		if len(Unique(s.GetColumn(x))) == 9 {
			continue
		}

		s.SolveColumnByBlocked(x)
	}

	s.Validate()

	cAfter := s.GetCompleted()

	fmt.Printf("Cells completed (%d/81)\n", cAfter)

	// if we completed more numbers recur
	if cBefore != cAfter {
		s.Solve()
	} else {
		fmt.Println(s.GetName(), float32(cAfter)/float32(81)*100, "% Completed")
	}
}
