package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	beginLines := []string{`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`}

	for _, b := range beginLines {
		if regexp.MustCompile(b).MatchString(text) {
			return true
		}
	}

	return false
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<[~*-=]*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			count += 1
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var newLines []string
	re := regexp.MustCompile(`User\s+(\w+)`)
	for _, line := range lines {
		sl := re.FindStringSubmatch(line)
		if sl != nil && sl[1] != "" {
			newLines = append(newLines, fmt.Sprintf("[USR] %s %s", sl[1], line))
		} else {
			newLines = append(newLines, line)
		}
	}
	return newLines
}
