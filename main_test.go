package main

import "testing"

const filename = "test.txt"

type Case struct {
	want int64
	fn   func(string) (int64, error)
}

func TestGetBytesCount(t *testing.T) {
	cases := []Case{
		{
			want: 342190,
			fn:   getBytesCount,
		},
		{
			want: 7145,
			fn:   getLinesCount,
		},
		{
			want: 58164,
			fn:   getWordsCount,
		},
		{
			want: 33929,
			fn:   getCharCount,
		},
	}

	for _, each_case := range cases {
		if got, _ := each_case.fn(filename); got != each_case.want {
			t.Errorf("want %v but got %v", each_case.want, got)
		}
	}
}
