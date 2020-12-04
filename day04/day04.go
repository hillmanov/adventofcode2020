package main

import (
	"adventofcode/utils"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	BirthYear      string `json:"byr"`
	IssueYear      string `json:"iyr"`
	ExpirationYear string `json:"eyr"`
	Height         string `json:"hgt"`
	HairColor      string `json:"hcl"`
	EyeColor       string `json:"ecl"`
	PassportID     string `json:"pid"`
	CountryID      string `json:"cid"`
}

func (p passport) LooseValidate() bool {
	if len(p.BirthYear) == 0 || len(p.IssueYear) == 0 || len(p.ExpirationYear) == 0 || len(p.Height) == 0 || len(p.HairColor) == 0 || len(p.EyeColor) == 0 || len(p.PassportID) == 0 {
		return false
	}
	return true
}

func (p passport) TightValidate() bool {
	// (Birth Year) - four digits; at least 1920 and at most 2002.
	if !inRange(p.BirthYear, 1920, 2002) {
		return false
	}

	// (Issue Year) - four digits; at least 2010 and at most 2020.
	if !inRange(p.IssueYear, 2010, 2020) {
		return false
	}

	// (Expiration Year) - four digits; at least 2020 and at most 2030.
	if !inRange(p.ExpirationYear, 2020, 2030) {
		return false
	}

	// (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	heightParts := regexp.MustCompile(`^(\d+)(cm|in)$`)
	if !heightParts.Match([]byte(p.Height)) {
		return false
	}
	parts := heightParts.FindStringSubmatch(p.Height)
	height := parts[1]
	unit := parts[2]
	switch unit {
	case "in":
		if !inRange(height, 59, 76) {
			return false
		}
	case "cm":
		if !inRange(height, 150, 193) {
			return false
		}
	}

	// (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	hairColorValidator := regexp.MustCompile(`^#[a-f0-9]{6}$`)
	if !hairColorValidator.Match([]byte(p.HairColor)) {
		return false
	}

	// (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	eyeColorValidator := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	if !eyeColorValidator.Match([]byte(p.EyeColor)) {
		return false
	}

	// (Passport ID) - a nine-digit number, including leading zeroes.

	passportIDValidator := regexp.MustCompile(`^\d{9}$`)
	if !passportIDValidator.Match([]byte(p.PassportID)) {
		return false
	}

	// (Country ID) - ignored, missing or not.
	return true
}

func main() {
	input, err := utils.ReadContents("./input.txt")
	if err != nil {
		panic(err)
	}

	passports := parsePassports(input)

	part1Solution := part1(passports)
	part2Solution := part2(passports)

	fmt.Printf("Day 04: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 04: Part 2: = %+v\n", part2Solution)
}

func part1(passports []passport) int {
	validCount := 0
	for _, p := range passports {
		if p.LooseValidate() {
			validCount++
		}
	}
	return validCount
}

func part2(passports []passport) int {
	validCount := 0
	for _, p := range passports {
		if p.TightValidate() {
			validCount++
		}
	}
	return validCount
}

func inRange(s string, min, max int) bool {
	value, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if value < min || value > max {
		return false
	}
	return true
}

func parsePassports(input string) []passport {
	block := regexp.MustCompile(`(?m)^$`)
	passportBatches := block.Split(input, -1)

	flatten := regexp.MustCompile(`\s`)
	keyValue := regexp.MustCompile(`(\w+):([\w#]+)`)
	passports := []passport{}
	for _, passportBatch := range passportBatches {
		// Convert multiline batch to single line json object
		passportBatch := flatten.ReplaceAllString(strings.TrimSpace(passportBatch), ",")
		passportBatch = keyValue.ReplaceAllString(passportBatch, `"${1}":"${2}"`)
		passportBatch = "{" + passportBatch + "}"

		pass := passport{}
		json.Unmarshal([]byte(passportBatch), &pass)
		passports = append(passports, pass)
	}
	return passports
}
