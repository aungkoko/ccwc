package main

import (
	"os"
	"reflect"
	"testing"
)

const filename = "test.txt"

func TestGetBytesCount(t *testing.T) {
	bytes, _ := os.ReadFile(filename)
	r := Result{}
	r.SetResult(bytes)

	want := Result{ByteCount: 342190, CharCount: 339292, LineCount: 7145, WordCount: 58164}

	if !reflect.DeepEqual(r, want) {
		t.Errorf("want %v got %v", want, r)
	}
}
