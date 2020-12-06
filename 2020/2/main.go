package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	inputFilePath       string
	passwordPolicyRegex = regexp.MustCompile(`^(\d+)-(\d+)\s+([a-zA-z]):\s+([a-zA-z0-9]+)$`)
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
	policies := make([]string, 0)
	for scanner.Scan() {
		policies = append(policies, scanner.Text())
	}

	runPart1(policies)
	runPart2(policies)
}

func runPart1(policies []string) {
	count := 0
	for _, raw := range policies {
		policy, err := parsePasswordPolicy(raw)
		if err != nil {
			panic(fmt.Errorf("failed to parse password policy: %v", err))
		}

		if policy.isValidPart1Password() {
			count++
		}
	}

	fmt.Println(count)
}

func runPart2(policies []string) {
	count := 0
	for _, raw := range policies {
		policy, err := parsePasswordPolicy(raw)
		if err != nil {
			panic(fmt.Errorf("failed to parse password policy: %v", err))
		}

		if policy.isValidPart2Password() {
			count++
		}
	}

	fmt.Println(count)
}

type passwordPolicy struct {
	minimum  int
	maximum  int
	letter   string
	password string
}

func (p *passwordPolicy) isValidPart1Password() bool {
	frequency := strings.Count(p.password, p.letter)
	return frequency >= p.minimum && frequency <= p.maximum
}

func (p *passwordPolicy) isValidPart2Password() bool {
	isLetterAtFirstPosition := string(p.password[p.minimum - 1]) == p.letter
	isLetterAtSecondPosition := string(p.password[p.maximum - 1]) == p.letter
	return isLetterAtFirstPosition != isLetterAtSecondPosition
}

func parsePasswordPolicy(raw string) (*passwordPolicy, error) {
	results := passwordPolicyRegex.FindAllStringSubmatch(raw, -1)
	if len(results) < 1 {
		return nil, fmt.Errorf("no matches found for password policy '%s'", raw)
	}

	// The regexp package sets the first submatch to the original string hence the length needs to be 5
	matches := results[0]
	if len(matches) < 5 {
		panic(fmt.Errorf("failed to find matches for all subexpressions for password policy '%s'", raw))
	}

	minimum, err := strconv.ParseInt(matches[1], 10, 32)
	if err != nil {
		return nil, err
	}

	maximum, err := strconv.ParseInt(matches[2], 10, 32)
	if err != nil {
		return nil, err
	}

	policy := &passwordPolicy{
		minimum:  int(minimum),
		maximum:  int(maximum),
		letter:   matches[3],
		password: matches[4],
	}


	return policy, nil
}
