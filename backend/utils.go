package main

import (
	"regexp"
	"strconv"
	"strings"
)

func formatText(input string) string {
	return strings.ToUpper(strings.TrimSpace(input))
}

func parseAge(input string) int {
	ageRegex := regexp.MustCompile(`\d+\s*(?i)(TAHUN|THN|TH)?`)
	match := ageRegex.FindStringSubmatch(input)

	var age int
	if match != nil {
		ageStr := match[0]
		ageRegex := regexp.MustCompile(`\d+`)
		ageMatch := ageRegex.FindString(ageStr)
		age = atoi(ageMatch)
	}

	return age
}

func atoi(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}
