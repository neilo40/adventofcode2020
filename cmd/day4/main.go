package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/neilo40/adventofcode2020/internal/common"
)

func main() {
	lines := common.ReadFileString("day4.input")
	validCount := 0
	var currentPassport string
	for _, l := range lines {
		if l == "" {
			if isValidPassport(currentPassport) {
				validCount++
			} else {
				fmt.Printf("Invalid passport: %s\n\n", currentPassport)
			}
			currentPassport = ""
		} else {
			currentPassport += " " + l
		}
	}

	// The last line has no newline after it, so handle it here
	if isValidPassport(currentPassport) {
		validCount++
	}

	fmt.Printf("Valid passports found: %d\n", validCount)
}

func isValidPassport(passport string) bool {
	if !yearIsValid(passport, "byr", 1920, 2002) {
		return false
	}
	if !yearIsValid(passport, "iyr", 2010, 2020) {
		return false
	}
	if !yearIsValid(passport, "eyr", 2020, 2030) {
		return false
	}
	if !heightIsValid(passport) {
		return false
	}
	if !hairColourIsValid(passport) {
		return false
	}
	if !eyeColourIsValid(passport) {
		return false
	}
	if !passportIdIsValid(passport) {
		return false
	}

	return true
}

func yearIsValid(passport string, field string, min int64, max int64) bool {
	matcher := regexp.MustCompile(fmt.Sprintf(`%s:(\d+)`, field))
	matches := matcher.FindStringSubmatch(passport)
	if len(matches) < 2 {
		fmt.Printf("%s is incorrectly formed or missing\n", field)
		return false
	} else {
		year, _ := strconv.ParseInt(matches[1], 10, 64)
		if year < min || year > max {
			fmt.Printf("%s is out of range\n", field)
			return false
		}
	}
	return true
}

func heightIsValid(passport string) bool {
	matcher := regexp.MustCompile(`hgt:(\d+)(cm|in)`)
	matches := matcher.FindStringSubmatch(passport)
	if len(matches) < 3 {
		fmt.Println("Height is incorrectly formed or missing")
		return false
	} else {
		height, _ := strconv.ParseInt(matches[1], 10, 64)
		if matches[2] == "cm" && (height < 150 || height > 193) {
			fmt.Println("Height in cm is out of range")
			return false
		}
		if matches[2] == "in" && (height < 59 || height > 76) {
			fmt.Println("Height in in is out of range")
			return false
		}
	}
	return true
}

func hairColourIsValid(passport string) bool {
	matcher := regexp.MustCompile(`hcl:#[0-9a-f]{6}(\s|$)`)
	matched := matcher.MatchString(passport)
	if !matched {
		fmt.Println("Hair colour is malformed or missing")
	}
	return matched
}

func eyeColourIsValid(passport string) bool {
	matcher := regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)`)
	matched := matcher.MatchString(passport)
	if !matched {
		fmt.Println("Eye colour is malformed or missing")
	}
	return matched
}

func passportIdIsValid(passport string) bool {
	matcher := regexp.MustCompile(`pid:\d{9}(\s|$)`)
	matched := matcher.MatchString(passport)
	if !matched {
		fmt.Println("Passport ID is malformed or missing")
	}
	return matched
}
