package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	requiredFieldsCount = 7
)

var (
	inputFilePath       string
	validHairColorRegex = regexp.MustCompile(`#[0-9a-f]{6}`)
	heightRegex         = regexp.MustCompile(`([0-9])+(in|cm)`)
	validEyeColors      = map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
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
	passports := make([]Passport, 0)
	pairs := make([]string, 0, 0)
	for scanner.Scan() {
		if scanner.Text() == "" {
			passports = append(passports, NewPassport(pairs))
			pairs = make([]string, 0, 0)
			continue
		}

		pairs = append(pairs, strings.Split(scanner.Text(), " ")...)
	}

	// We need to explicitly add the last passport because the condition inside the for loop is never hit for it
	passports = append(passports, NewPassport(pairs))

	runPart1(passports)
	runPart2(passports)
}

func runPart1(passports []Passport) {
	count := 0
	for _, passport := range passports {
		if passport.areAllFieldsPresent() {
			count++
		}
	}

	fmt.Println(count)
}

func runPart2(passports []Passport) {
	count := 0
	for _, passport := range passports {
		if passport.isValid() {
			count++
		}
	}

	fmt.Println(count)
}

type Passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
	FieldsPresent  int
}

func NewPassport(raw []string) Passport {
	passport := Passport{}
	for _, pair := range raw {

		split := strings.Split(pair, ":")
		key := split[0]
		value := split[1]

		switch key {
		case "byr":
			passport.BirthYear = value
			passport.FieldsPresent++
		case "iyr":
			passport.IssueYear = value
			passport.FieldsPresent++
		case "eyr":
			passport.ExpirationYear = value
			passport.FieldsPresent++
		case "hgt":
			passport.Height = value
			passport.FieldsPresent++
		case "hcl":
			passport.HairColor = value
			passport.FieldsPresent++
		case "ecl":
			passport.EyeColor = value
			passport.FieldsPresent++
		case "pid":
			passport.PassportID = value
			passport.FieldsPresent++
		case "cid":
			// Nothing to do
		default:
			panic(fmt.Errorf("%s is not a value field: %s", key, pair))
		}
	}

	return passport
}

func (p Passport) isValid() bool {
	_, isValidEyeColor := validEyeColors[p.EyeColor]
	return p.areAllFieldsPresent() && p.areValidYears() && validHairColorRegex.MatchString(p.HairColor) && isValidEyeColor && p.isValidHeight() && len(p.PassportID) == 9
}

func (p Passport) areAllFieldsPresent() bool {
	return p.FieldsPresent == requiredFieldsCount
}

func (p Passport) areValidYears() bool {
	isValidBirthYear := len(p.BirthYear) == 4 && p.BirthYear >= "1920" && p.BirthYear <= "2002"
	isValidIssueYear := len(p.IssueYear) == 4 && p.IssueYear >= "2010" && p.IssueYear <= "2020"
	isValidExpirationYear := len(p.ExpirationYear) == 4 && p.ExpirationYear >= "2020" && p.ExpirationYear <= "2030"
	return isValidBirthYear && isValidIssueYear && isValidExpirationYear
}

func (p Passport) isValidHeight() bool {
	if i := strings.LastIndex(p.Height, "cm"); i != -1 {
		height := p.Height[:i]
		return height >= "150" && height <= "193"
	} else if i := strings.LastIndex(p.Height, "in"); i != -1 {
		i := strings.LastIndex(p.Height, "in")
		height := p.Height[:i]
		return height >= "59" && height <= "76"
	} else {
		return false
	}
}

func isValidPassport(passport ...string) bool {
	fmt.Println(passport)
	if len(passport) < 7 {
		return false
	}

	fields := 0
	for _, pair := range passport {

		split := strings.Split(pair, ":")
		key := split[0]

		switch key {
		case "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid":
			fields++
		case "cid":
			// Nothing to do
		default:
			panic(fmt.Errorf("%s is not a value field: %s", key, pair))
		}
	}

	return fields == requiredFieldsCount
}
