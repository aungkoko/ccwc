package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getBytesCount(filename string) (int64, error) {
	bytes, err := os.ReadFile(filename)
	var byteCount int64
	if err != nil {
		fmt.Println("Error reading file:", err)
		return byteCount, err
	}

	byteCount = int64(len(bytes))
	return byteCount, nil
}

func getLinesCount(filename string) (int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lineCount := 0

	for sc.Scan() {
		lineCount++
	}

	if err := sc.Err(); err != nil {
		return 0, err
	}

	return int64(lineCount), nil
}

func getWordsCount(filename string) (int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	wordCount := 0

	for sc.Scan() {
		words := strings.Fields(sc.Text())
		wordCount += len(words)
	}

	if err := sc.Err(); err != nil {
		return 0, err
	}

	return int64(wordCount), nil
}

func getCharCounter(filename string) (int64, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, err
	}

	content_str := string(content)

	return int64(len([]rune(content_str))), nil
}

func main() {
	args := os.Args
	if len(args) > 1 {
		option := args[1]
		filename := args[2]

		switch option {
		case "-c":
			if byteCount, err := getBytesCount(filename); err == nil {
				fmt.Println(byteCount, filename)
			} else {
				fmt.Println(err)
			}
		case "-l":
			if lineCount, err := getLinesCount(filename); err == nil {
				fmt.Println(lineCount, filename)
			} else {
				fmt.Println(err)
			}
		case "-w":
			if wordCount, err := getWordsCount(filename); err == nil {
				fmt.Println(wordCount, filename)
			} else {
				fmt.Println(err)
			}
		case "-m":
			if wordCount, err := getCharCounter(filename); err == nil {
				fmt.Println(wordCount, filename)
			} else {
				fmt.Println(err)
			}
		}
	}
}
