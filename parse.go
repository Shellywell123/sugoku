package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// function to create a sudoku object from a input txt file
func ImportSudokuFromFile(filename string) Sudoku {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fscanner := bufio.NewScanner(file)

	sudoku := Sudoku{Grid: make([][]int, 10)}

	y := 0
	for fscanner.Scan() {
		line := fscanner.Text()
		row := []int{}

		for _, s := range strings.Split(line, " ") {
			if s == "" {
				continue
			}
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err)
			}

			row = append(row, num)
		}

		sudoku.SetRow(y, row)
		y++
	}

	return sudoku
}

// simple funciton to pretty print a sudokus current state in the terminal
func PrintSudoku(s Sudoku) {

	for rNumber, r := range s.Grid {
		if rNumber%3 == 0 || rNumber == 0 {
			fmt.Printf("+-------+-------+-------+\n")
		}
		for cNumber, c := range r {
			if cNumber%3 == 0 {
				fmt.Printf("| ")
			}
			cell := strconv.Itoa(c)
			if cell == "0" {
				cell = " "
			}
			fmt.Printf("%s ", cell)
		}
		if rNumber != 9 {
			fmt.Printf("|\n")
		}
	}
}
