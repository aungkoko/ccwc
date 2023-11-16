package main

import "testing"

const filename = "test.txt"

func TestGetBytesCount(t *testing.T) {
	want := 342190
	if got, _ := getBytesCount(filename); got != int64(want) {
		t.Errorf("want %v but got %v", want, got)
	}
}

func TestGetLinesCount(t *testing.T) {
	want := 7145
	if got, _ := getLinesCount(filename); got != int64(want) {
		t.Errorf("want %v but got %v", want, got)
	}
}

func TestGetWordsCount(t *testing.T) {
	want := 58164
	if got, _ := getWordsCount(filename); got != int64(want) {
		t.Errorf("want %v but got %v", want, got)
	}
}

func TestGetCharCount(t *testing.T) {
	want := 339292
	if got, _ := getCharCount(filename); got != int64(want) {
		t.Errorf("want %v but got %v", want, got)
	}
}
