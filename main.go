package main

import (
	"fmt"
	"md2html/utils"
	"os"
	"slices"
	"strings"
)

func main() {
	var name string
	if len(os.Args) > 1 {
		name = os.Args[1]
	} else {
		name = "example.md"
	}
	file := utils.ReadFile(name)
	lines := strings.Split(file, "\n")
	numLists := 0
	newFile := []string{"<html>", "<body>"}
	for _, line := range lines {
		switch {
		case utils.StartsWith(line, '#'):
			header := replaceHeader(line)
			newFile = append(newFile, header)
		case utils.StartsWith(line, '-'):
			if numLists < 1 {
				newFile = append(newFile, "<ul>\n")
				numLists++
			}
			item := convertListItem(line)
			newFile = append(newFile, fmt.Sprintf("\t%s", item))
		default:
			if numLists > 0 {
				newFile = append(newFile, "</ul>\n")
				numLists--
			} else {
				if len(line) > 0 {
					newFile = append(newFile, "<p>"+line+"</p>")
				}
			}
		}
	}
	path, err := writeToFile(newFile, name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("'%s' written succesfully!\n", path)
	}
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

func indentInnerLines(file string) string {
	result := ""
	lines := strings.Split(file, "\n")
	for i, line := range lines {
		if i > 1 && i < len(lines)-2 {
			result += fmt.Sprintf("\t%s\n", line)
		} else {
			result += fmt.Sprintf("%s\n", line)
		}
	}
	return result
}

func writeToFile(file []string, name string) (string, error) {
	file = utils.TrimLines(file)
	temp := slices.Concat(file, []string{"</body>", "</html>"})
	result := strings.Join(temp, "\n")
	result = indentInnerLines(result)
	path := utils.ExtractFileName(name) + ".html"
	err := utils.WriteFile(path, result)
	return path, err
}
