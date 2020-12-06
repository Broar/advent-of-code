package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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
	expenses := make([]int, 0)
	for scanner.Scan() {
		expense, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			panic(fmt.Errorf("failed to open parse integer: %v", err))
		}

		expenses = append(expenses, int(expense))
	}

	runPart1(expenses)
	runPart2(expenses)
}

func runPart1(expenses []int) {
	found := false
	for i := 0; i < len(expenses) && !found; i++ {
		for j := 0; j < len(expenses) && !found; j++ {
			if expenses[i]+expenses[j] == 2020 {
				fmt.Printf("%d * %d = %d\n", expenses[i], expenses[j], expenses[i]*expenses[j])
				found = true
			}
		}
	}
}

func runPart2(expenses []int) {
	found := false
	for i := 0; i < len(expenses) && !found; i++ {
		for j := 0; j < len(expenses) && !found; j++ {
			for k := 0; k < len(expenses) && !found; k++ {
				if expenses[i]+expenses[j]+expenses[k] == 2020 {
					fmt.Printf("%d * %d * %d= %d\n", expenses[i], expenses[j], expenses[k], expenses[i]*expenses[j]*expenses[k])
					found = true
				}
			}
		}
	}
}
