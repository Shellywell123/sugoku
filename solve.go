package main

import (
	"fmt"
	"slices"

	"golang.org/x/text/number"
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

// solve a cell by checking all other positions are blocked in the row
func (s Suduko) SolveSquareByRow(x int, y int) {
	row := s.GetRow(y)
	col := s.GetColumn(x)
	sqa := s.GetSquare(x, y)

	rowSet := Unique(row)
	colSet := Unique(col)
	sqaSet := Unique(sqa)

	set := Unique(slices.Concat(nil, rowSet, colSet, sqaSet))

	choices := []int{}
	for n := 1; n <= 9; n++ {
		
		found := false
		for _, m := range set {
			if m == n {
				found = true
			}
		}

		if !found {
			choices = append(choices, n)
		}
	}

	// get other possilbe positions in square todo

	availableRowPostions := slices.Index(row, 0) // need a list of index not singular!
	// positions := [][]int{} x,y coords

	for _, choice := range choices {
		// prove that the cell is the only place for the choice

		choiceMustGoInsekectedcell := true

		// logic todo
		if logic {
			choiceMustGoInsekectedcell = false
		}
		for _, position := range availableRowPostions {
			positionBlocked := false
			for num := range GetColumn(position , y) {
				if choice == number {
					positionBlocked = true
				}
				
			}
		}

		if !choiceMustGoInsekectedcell {
			s.SetCell(x, y, choice)
			break
		}
	}
}

// main solve function
func (s Suduko) Solve() {

	cBefore := s.GetCompleted()

	// loop through all cells in the 9x9 grid
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {

			if s.GetCell(x, y) != 0 {
				continue
			}

			s.SolveUnique(x, y)
			// s.SolveSquareByBlocked(x,y)
			// s.SolveRowByBlocked(x,y)
			// s.SolveColumnByBlocked(x,y)
		}
	}
	cAfter := s.GetCompleted()

	fmt.Printf("Cells completed (%d/81)\n", cAfter)

	// if we completed more numbers recur
	if cBefore != cAfter {
		s.Solve()
	}
}
