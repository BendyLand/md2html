package main

import (
	"fmt"
	"strings"
	"md2html/utils"
)

func main() {
	file := utils.ReadFile("example.md")
	lines := strings.Split(file, "\n")
	numLists := 0
	newFile := ""
	for _, line := range lines {
		switch {
		case utils.StartsWith(line, '#'):
			header := replaceHeader(line)
			newFile += header
		case utils.StartsWith(line, '-'):
			if numLists < 1 {
				newFile += "<ul>\n"
				numLists++
			}
			item := convertListItem(line)
			newFile += fmt.Sprintf("\t%s", item)
		default:
			if numLists > 0 {
				newFile += "</ul>"
				numLists--
			}
		}
	}
	fmt.Println(newFile)
}

func replaceHeader(line string) string {
	num := strings.Count(line, "#")
	start := strings.Index(line, " ")
	substr := strings.TrimLeft(line[start:], " ")
	result := fmt.Sprintf("<h%d>%s</h%d>\n", num, substr, num)
	return result
}

func convertListItem(line string) string {
	start := strings.Index(line, "-") + 1
	substr := strings.TrimLeft(line[start:], " ")
	result := fmt.Sprintf("<li>%s</li>\n", substr)
	return result
}

func constructUnorderedList(lines []string) string {
	result := "<ul>\n"
	for _, line := range lines {
		temp := convertListItem(line)
		result += fmt.Sprintf("\t%s\n", temp)
	}
	result += "</ul>"
	return result
}
