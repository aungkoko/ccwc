package main

import (
	"fmt"
	"os"
)

func CharCounter(filename string) {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	content_str := string(file)

	fmt.Println(len([]rune(content_str)))
}
