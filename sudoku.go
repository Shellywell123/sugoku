package main

// import "fmt"

type sudoku interface {
	SetCell(x int, y int, v int)
	GetCell(x int, y int)
	GetRow(y int)
	GetColumn(x int)
	GetSquare(sx int, sy int)
	GetCompleted()
}

type Sudoku struct {
	Grid      [][]int
	Completed int
}

func (s Sudoku) GetCell(x int, y int) int {
	return s.Grid[y][x]
}

func (s *Sudoku) SetCell(x int, y int, v int) {
	s.Grid[y][x] = v
}

func (s Sudoku) GetRow(y int) []int {
	return s.Grid[y]
}

func (s *Sudoku) SetRow(y int, r []int) {
	s.Grid[y] = r
}

func (s Sudoku) GetColumn(x int) []int {
	column := []int{}
	for i := 0; i < 9; i++ {
		column = append(column, s.Grid[i][x])
	}
	return column
}

func (s Sudoku) GetSquare(x int, y int) []int {

	// i dont like how this is written
	sx := 0
	sy := 0

	if 0 <= x && x <= 2 {
		sx = 0
	}
	if 3 <= x && x <= 5 {
		sx = 1
	}
	if 6 <= x && x <= 8 {
		sx = 2
	}

	if 0 <= y && y <= 2 {
		sy = 0
	}
	if 3 <= y && y <= 5 {
		sy = 1
	}
	if 6 <= y && y <= 8 {
		sy = 2
	}
	square := []int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			square = append(square, s.GetCell(sx*3+j, sy*3+i))
		}
	}

	return square
}

func (s Sudoku) GetCompleted() int {

	c := 0
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if s.GetCell(x, y) != 0 {
				c++
			}
		}
	}
	return c
}
