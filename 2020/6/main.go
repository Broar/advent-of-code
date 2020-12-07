package main

import (
	"bufio"
	"flag"
	"fmt"
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
	patterns := make([]string, 0)
	pattern := ""
	groups := make([][]string, 0, 0)
	people := make([]string, 0, 0)
	for scanner.Scan() {
		if scanner.Text() == "" {
			patterns = append(patterns, pattern)
			groups = append(groups, people)
			pattern = ""
			people = make([]string, 0, 0)
			continue
		}

		pattern += scanner.Text()
		people = append(people, scanner.Text())
	}

	patterns = append(patterns, pattern)
	groups = append(groups, people)

	runPart1(patterns)
	runPart2(groups)
}

func runPart1(patterns []string) {
	sum := 0
	for _, pattern := range patterns {
		answered := make(map[string]bool)
		for _, answer := range pattern {
			answered[string(answer)] = true
		}

		sum += len(answered)
	}

	fmt.Println(sum)
}

func runPart2(groups [][]string) {
	sum := 0
	for _, group := range groups {
		counts := make(map[string]int)
		for _, person := range group {
			for _, answer := range person {
				counts[string(answer)] = counts[string(answer)] + 1
			}
		}

		for _, count := range counts {
			if count == len(group) {
				sum++
			}
		}
	}

	fmt.Println(sum)
}