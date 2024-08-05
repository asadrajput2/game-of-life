package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Grid struct {
	height, width int
	matrix        [][]int
}

func newGrid(height, width int) Grid {
	return Grid{
		height: height,
		width:  width,
		matrix: makeEmptyMatrix(height, width),
	}
}

func (grid *Grid) print() {
	for i := range grid.matrix {

		for j := range grid.matrix[i] {

			if grid.matrix[i][j] == 1 {
				fmt.Printf("█ ")
			} else {
				fmt.Printf(". ")
			}

		}
		fmt.Println()
	}

	fmt.Println("\n-----------------------------------------------")
}

func (grid *Grid) click() {

	rows, columns := 25, 25
	newMatrix := makeEmptyMatrix(rows, columns)
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {

			liveNeighborsCount := grid.getLiveNeighborsCount(r, c)

			if grid.matrix[r][c] == 0 {

				if liveNeighborsCount == 3 {
					newMatrix[r][c] = 1
				}

			} else {

				if liveNeighborsCount < 2 || liveNeighborsCount > 3 {
					newMatrix[r][c] = 0
				} else {
					newMatrix[r][c] = 1
				}

			}

		}
	}

	grid.matrix = newMatrix

}

func (grid *Grid) getLiveNeighborsCount(r, c int) int {

	dirs := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0
	for _, d := range dirs {
		newR, newC := r+d[0], c+d[1]

		if newR >= 0 && newR < grid.width && newC >= 0 && newC < grid.height {
			count += grid.matrix[newR][newC]
		}

	}

	return count

}

func (grid *Grid) applyPattern(pattern [][]int, gridRow, gridCol int) {
	for i := range pattern {
		for j := range pattern[i] {
			grid.matrix[gridRow+i][gridCol+j] = pattern[i][j]

		}
	}
}

func makeEmptyMatrix(rows, cols int) [][]int {

	grid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]int, cols)
	}

	return grid
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
