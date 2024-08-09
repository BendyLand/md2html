package main

import (
	"fmt"
	"md2html/utils"
)

func main() {
	file := utils.ReadFile("example.md")
	fmt.Println(file)
}
