package utils

import (
	"fmt"
	"os"
)

func ReadFile(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Unable to read file:", err)
		return ""
	}
	return string(file)
}

func WriteFile(path string, contents string) {
	err := os.WriteFile(path, []byte(contents), 0644)
	if err != nil {
		fmt.Println("Unable to write file:", err)
	}
}