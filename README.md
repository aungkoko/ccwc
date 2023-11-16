# John Crickett's Coding Challenges


### About

My solution of the coding challenge's `wc` cli tool. [wc cli tool challenge](https://codingchallenges.fyi/challenges/challenge-wc/)


### Prerequisites

You need to install [go version1.21.x](https://go.dev/doc/install)

### Tests
```
go test -v
```

### How to install
```
go build
```

### Usage

Sample command

```
    ./ccwc -l sample.txt
```

You can use these flags on your desire.
- -c for `bytes count` in file
- -l for `line count` in file
- -w for `words count` in file
- -m for `character count` in file
  


If you would like to read from standard input, you can use like this
```
    less test.txt | ./ccwc -l
```
