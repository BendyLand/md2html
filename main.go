package main

import (
	"fmt"
	"strings"
	"md2html/utils"
)

func main() {
	file := utils.ReadFile("example.md")
	lines := strings.Split(file, "\n")
	test := replaceHeader(lines[0])
	fmt.Println(lines[0])
	fmt.Println(test)
}

func replaceHeader(line string) string {
	if !strings.Contains(line, "#") {
		return line
	}
	num := strings.Count(line, "#")
	start := strings.Index(line, " ")
	substr := strings.TrimLeft(line[start:], " ")
	result := fmt.Sprintf("<h%d>%s</h%d>", num, substr, num)
	return result
}

