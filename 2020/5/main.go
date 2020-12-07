package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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
	partitions := make([]string, 0)
	for scanner.Scan() {
		partitions = append(partitions, scanner.Text())
	}

	runPart1(partitions)
	runPart2(partitions)
}

func runPart1(partitions []string) {
	best := 0
	for _, partition := range partitions {
		row := findRow(partition[0:7])
		column := findColumn(partition[7:])
		id := (row * 8) + column
		if id > best {
			best = id
		}
	}

	fmt.Println(best)
}

func runPart2(partitions []string) {
	ids := make([]int, 0, 0)
	for _, partition := range partitions {
		row := findRow(partition[0:7])
		column := findColumn(partition[7:])
		id := (row * 8) + column
		ids = append(ids, id)
	}

	sort.Ints(ids)

	previous := 0
	for _, id := range ids {
		if id - previous == 2 {
			fmt.Println(id - 1)
			break
		}

		previous = id
	}
}

func findRow(partition string) int {
	return binarySearch(partition, 'F', 'B', 0, 127)
}

func findColumn(partition string) int {
	return binarySearch(partition, 'L', 'R', 0, 7)
}

func binarySearch(partition string, lower, upper rune, min, max int) int {
	start := min
	end := max
	for _, char := range partition {
		diff := ((end - start)/2) + 1

		if char == lower {
			end -= diff
		} else if char == upper {
			start += diff
		} else {
			panic(fmt.Errorf("%c is not a valid character", char))
		}
	}

	if rune(partition[len(partition)-1]) == lower {
		return start
	} else {
		return end
	}
}
