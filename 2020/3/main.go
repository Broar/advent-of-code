package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
)

var (
	inputFilePath string
)

func init() {
	flag.StringVar(&inputFilePath, "input-filepath", "", "The filepath to the input for the challenge")
	flag.Parse()
}

func main() {
	input, err := os.Open(inputFilePath)
	if err != nil {
		panic(fmt.Errorf("failed to open input: %v", err))
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)
	rows := make([][]bool, 0)
	for scanner.Scan() {
		row := make([]bool, len(scanner.Text()))
		for i, letter := range scanner.Text() {
			if letter == '#' {
				row[i] = true
			}
		}

		rows = append(rows, row)
	}

	runPart1(rows)
	runPart2(rows)
}

func runPart1(rows [][]bool) {
	fmt.Println(countTrees(rows, 3, 1))
}

type pair struct {
	right int
	down int
}

func runPart2(rows [][]bool) {
	pairs := []pair{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	total := 1
	for _, pair := range pairs {
		trees := countTrees(rows, pair.right, pair.down)
		total *= trees
	}

	fmt.Println(total)
}

func countTrees(rows [][]bool, right, down int) int {
	trees := 0
	x := 0
	y := 0
	for {
		x += right
		y += down
		if len(rows) <= y {
			break
		}

		if len(rows[y]) <= x {
			// Account for overflow of x into the next section
			x = int(math.Abs(float64(len(rows[y]) - x)))
		}

		if rows[y][x] {
			trees++
		}
	}

	return trees
}
