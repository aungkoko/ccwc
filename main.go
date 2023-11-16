package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func getCharCount(filename string) (int64, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, err
	}

	content_str := string(content)

	return int64(len([]rune(content_str))), nil
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

func main() {
	args := os.Args

	var filename string
	var result []string

	if len(args) > 2 {
		flag := args[1]
		filename = args[2]

		switch flag {
		case "-c":
			if byteCount, err := getBytesCount(filename); err == nil {
				result = append(result, strconv.FormatInt(byteCount, 10))
			} else {
				fmt.Println(err)
			}
		case "-l":
			if lineCount, err := getLinesCount(filename); err == nil {
				result = append(result, strconv.FormatInt(lineCount, 10))
			} else {
				fmt.Println(err)
			}
		case "-w":
			if wordCount, err := getWordsCount(filename); err == nil {
				result = append(result, strconv.FormatInt(wordCount, 10))
			} else {
				fmt.Println(err)
			}
		case "-m":
			if charCount, err := getCharCount(filename); err == nil {
				result = append(result, strconv.FormatInt(charCount, 10))
			} else {
				fmt.Println(err)
			}
		}
	}

	if len(args) == 2 {
		filename = args[1]
		byteCount, _ := getBytesCount(filename)
		result = append(result, strconv.FormatInt(byteCount, 10))

		lineCount, _ := getLinesCount(filename)
		result = append(result, strconv.FormatInt(lineCount, 10))

		wordCount, _ := getWordsCount(filename)
		result = append(result, strconv.FormatInt(wordCount, 10))
	} else {

	}

	result = append(result, filename)

	fmt.Println(strings.Join(result, " "))
}
