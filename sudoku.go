package main

import (
	"math"
)

type sudoku interface {
	GetName()
	SetName(name string)
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
	Name      string
}

func (s Sudoku) GetName() string {
	return s.Name
}

func (s *Sudoku) SetName(name string) {
	s.Name = name
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

	// get square position 
	sx := int(math.Floor(float64(x)/3))
	sy := int(math.Floor(float64(y)/3))

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
