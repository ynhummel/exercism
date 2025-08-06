package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	reg := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return reg.MatchString(text)
}

func SplitLogLine(text string) []string {
	reg := regexp.MustCompile(`<(\*|=|-|~)*>`)
	return reg.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	reg := regexp.MustCompile(`(?i)"[\w\s]*password[\w\s]*"`)

	count := 0
	for _, line := range lines {
		if reg.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	reg := regexp.MustCompile(`end-of-line\d*`)
	return reg.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	taggedLines := make([]string, len(lines))
	reg := regexp.MustCompile(`User\s+(\w+)`)

	for ind, line := range lines {
		submatches := reg.FindStringSubmatch(line)
		if len(submatches) > 0 {
			line = "[USR] " + submatches[1] + " " + line
		}
		taggedLines[ind] = line
	}
	return taggedLines
}
