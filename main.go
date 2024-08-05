package main

import (
	"fmt"
	"time"
)

func main() {

	generations := 50
	rows, cols := 25, 25

	grid := newGrid(rows, cols)
	grid.applyPattern(GLIDER_PATTERN, rows/2-1, cols/2-1)

	for i := 0; i < generations; i++ {

		clearConsole()

		fmt.Println("GENERATION: ", i+1)

		grid.print()

		grid.click()

		time.Sleep(100 * time.Millisecond)
	}

}
