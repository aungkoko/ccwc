package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Result struct {
	ByteCount int64
	CharCount int64
	LineCount int64
	WordCount int64
}

func (r *Result) SetResult(content []byte) {
	r.ByteCount = int64(len(content))

	content_str := string(content)
	r.LineCount = int64(strings.Count(content_str, "\n"))

	r.CharCount = int64(len([]rune(content_str)))

	r.WordCount = int64((len(bytes.Fields(content))))
}

func main() {
	var bytes []byte
	var filename string
	var err error
	args := os.Args

	defaultPrinter := map[string]bool{"-c": false, "-l": false, "-w": false, "-m": false}

	if len(args) == 3 {
		filename = args[2]
		bytes, err = os.ReadFile(filename)
		if err != nil {
			log.Fatal(err.Error())
		}
		defaultPrinter[args[1]] = true
	}

	if len(args) == 2 {
		first_arg := args[1]
		if first_arg == "-c" || first_arg == "-l" || first_arg == "-w" || first_arg == "-m" {
			defaultPrinter[args[1]] = true
			bytes, err = io.ReadAll(os.Stdin)
			if err != nil {
				log.Fatal(err.Error())
			}
		} else {
			filename = first_arg
			bytes, err = os.ReadFile(filename)
			if err != nil {
				log.Fatal(err.Error())
			}
			defaultPrinter["-c"] = true
			defaultPrinter["-l"] = true
			defaultPrinter["-w"] = true
		}
	}

	r := Result{}
	r.SetResult(bytes)

	if defaultPrinter["-l"] {
		fmt.Printf("%s ", strconv.Itoa(int(r.LineCount)))
	}

	if defaultPrinter["-w"] {
		fmt.Printf("%s ", strconv.Itoa(int(r.WordCount)))
	}

	if defaultPrinter["-c"] {
		fmt.Printf("%s ", strconv.Itoa(int(r.ByteCount)))
	}

	if defaultPrinter["-m"] {
		fmt.Printf("%s ", strconv.Itoa(int(r.CharCount)))
	}

	if filename == "" {
		fmt.Printf("\n")
	}

	if filename != "" {
		fmt.Printf("%s \n", filename)
	}
}
