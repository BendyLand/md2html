package main

import (
	"fmt"
	"os"
)

func main() {
	file := readFile("example.md")
	fmt.Println(file)
}

func readFile(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Unable to read file:", err)
		return ""
	}
	return string(file)
}

func writeFile(path string, contents string) {
	err := os.WriteFile(path, []byte(contents), 0644)
	if err != nil {
		fmt.Println("Unable to write file:", err)
	}
}
