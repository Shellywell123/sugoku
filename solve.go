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

// solve a cell by getting the unique set of numbers across row col square
func (s Suduko) SolveUnique(x int, y int) {

	row := s.GetRow(y)
	col := s.GetColumn(x)
	sqa := s.GetSquare(x, y)

	rowSet := Unique(row)
	colSet := Unique(col)
	sqaSet := Unique(sqa)

	set := Unique(slices.Concat(nil, rowSet, colSet, sqaSet))

	if len(set) == 8 {
		for n := 1; n <= 9; n++ {

			found := false
			for _, m := range set {
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

// // solve a cell by checking all other positions are blocked in the square
// func (s Suduko) SolveSquareByBlocked(x int, y int) {
// 	row := s.GetRow(y)
// 	col := s.GetColumn(x)
// 	sqa := s.GetSquare(x, y)

// 	rowSet := Unique(row)
// 	colSet := Unique(col)
// 	sqaSet := Unique(sqa)

// 	set := Unique(slices.Concat(nil, rowSet, colSet, sqaSet))

// 	choices := []int{}
// 	for n := 1; n <= 9; n++ {

// 		found := false
// 		for _, m := range set {
// 			if m == n {
// 				found = true
// 			}
// 		}

// 		if !found {
// 			choices = append(choices, n)
// 		}
// 	}

// 	// get other possilbe positions in square todo
// 	// positions := [][]int{} x,y coords

// 	for _, choice := range choices {
// 		// prove that the cell is the only place for the choice

// 		choiceMustGoInsekectedcell := true

// 		// logic todo
// 		if logic {
// 			choiceMustGoInsekectedcell = false
// 		}
// 		for _, position := range positions {
// 			if choice in GetRow(position x , position y) {

// 			}
// 		}

// 		if !choiceMustGoInsekectedcell {
// 			s.SetCell(x, y, choice)
// 			break
// 		}
// 	}
// }

func FindElementIndexsInSlice(slice []int, element int) []int {

	indexs := []int{}
	for i, e := range slice {
		if e == element {
			indexs = append(indexs, i)
		}
	}
	return indexs

}

// complete a row by checking all other positions are blocked
func (s Suduko) SolveSquareByRow(y int) {
	row := s.GetRow(y)
	// get other possilbe positions in square todo
	availableRowPostions := FindElementIndexsInSlice(row, 0)

	choices := []int{}
	for i := 1; i <= 9; i++ {
		if !slices.Contains(Unique(row), i) {
			choices = append(choices, i)
		}
	}
	// fmt.Println(y, choices,Unique(row) )

	for _, choice := range choices {
		unblockedPositions := []int{}
		positionsBlocked := 0
		for _, position := range availableRowPostions {
			positionBlocked := false

			coveringRowNumbers := Unique(s.GetColumn(position))
			coveringSquareNumbers := Unique(s.GetSquare(position, y))
			coveringSet := Unique(slices.Concat(nil, coveringRowNumbers, coveringSquareNumbers))

			if slices.Contains(coveringSet, choice) {
				positionBlocked = true
			}

			if positionBlocked {
				positionsBlocked++
			} else {
				unblockedPositions = append(unblockedPositions, position)
			}
		}

		if len(unblockedPositions) == 1 {
			s.SetCell(unblockedPositions[0], y, choice)
			break
		}
	}
}

// complete a column by checking all other positions are blocked
func (s Suduko) SolveSquareByColumn(x int) {
	col := s.GetColumn(x)
	// get other possilbe positions in square todo
	availableRowPostions := FindElementIndexsInSlice(col, 0)
	unblockedPositions := []int{}

	choices := []int{}
	for i := 1; i <= 9; i++ {
		if !slices.Contains(Unique(col), i) {
			choices = append(choices, i)
		}

	}
	// fmt.Println(choices,Unique(col) )

	for _, choice := range choices {
		positionsBlocked := 0
		for _, position := range availableRowPostions {
			positionBlocked := false

			coveringColumnNumbers := Unique(s.GetRow(position))
			coveringSquareNumbers := Unique(s.GetSquare(x, position))
			coveringSet := Unique(slices.Concat(nil, coveringColumnNumbers, coveringSquareNumbers))

			if slices.Contains(coveringSet, choice) {
				positionBlocked = true
			}

			if positionBlocked {
				positionsBlocked++
			} else {
				unblockedPositions = append(unblockedPositions, position)
			}
		}

		if len(unblockedPositions) == 1 {
			s.SetCell(x, unblockedPositions[0], choice)
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
func (s Suduko) Validate() {

	for x := 0; x < 9; x++ {
		if duplicateDetect(s.GetColumn(x)) {
			PrintSuduko(s)
			fmt.Println("duplicate detected in column", x)
			os.Exit(0)
		}
	}

	for y := 0; y < 9; y++ {
		if duplicateDetect(s.GetRow(y)) {
			PrintSuduko(s)
			fmt.Println("duplicate detected in row", y)
			os.Exit(0)

		}
	}

	for x := 0; x < 9; x += 3 {
		for y := 0; y < 9; y += 3 {
			if duplicateDetect(s.GetSquare(x, y)) {
				PrintSuduko(s)
				fmt.Println("duplicate detected in row", y)
				os.Exit(0)

			}
		}
	}

}

// main solve function
func (s Suduko) Solve() {

	cBefore := s.GetCompleted()

	// loop through all cells in the 9x9 grid
	for y := 0; y < 9; y++ {

		// if row complete skip
		if len(Unique(s.GetRow(y))) == 9 {
			continue
		}

		s.SolveSquareByRow(y)

	}

	// loop through all cells in the 9x9 grid
	for y := 0; y < 9; y++ {

		for x := 0; x < 9; x++ {

			// if cell populated skip
			if s.GetCell(x, y) != 0 {
				continue
			}

			s.SolveUnique(x, y)
		}
	}

	for x := 0; x < 9; x++ {

		// if col complete skip
		if len(Unique(s.GetColumn(x))) == 9 {
			continue
		}

		s.SolveSquareByColumn(x)
	}

	s.Validate()

	cAfter := s.GetCompleted()

	fmt.Printf("Cells completed (%d/81)\n", cAfter)
	// PrintSuduko(s)

	// if we completed more numbers recur
	if cBefore != cAfter {
		s.Solve()
	}
}
