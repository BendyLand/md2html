package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Unable to read file:", err)
		return ""
	}
	return string(file)
}

func WriteFile(path string, contents string) error {
	err := os.WriteFile(path, []byte(contents), 0644)
	if err != nil {
		return fmt.Errorf("Unable to write file: %s", err)
	}
	return nil
}

func StartsWith(line string, c rune) bool {
	for _, cur := range line {
		if cur == ' ' { 
			continue
		}
		return cur == c
	}
	return false
}

func ExtractFileName(filename string) string {
	idx := strings.LastIndex(filename, ".")
	return filename[:idx]
}

func TrimLines(lines []string) []string {
	result := make([]string, 0)
	for _, line := range lines {
		if len(line) == 0 || strings.Trim(line, " \n") == "" {
			continue
		}
		result = append(result, strings.Trim(line, " \n"))
	}
	return result
}