// Package word uses various methods to count words in a string
package word

import "strings"

// UseCount counts the word frequncy in a string
// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns to the total number of words in a string
func Count(s string) int {
	// write the code for this func

	//my way
	return strings.Count(s, " ") + 1

	// √ word % go test -bench .
	// goos: darwin
	// goarch: amd64
	// pkg: github.com/imattf/go211/ninj1300/02/word
	// BenchmarkCount-6         9046837               132 ns/op
	// PASS
	// ok      github.com/imattf/go211/ninj1300/02/word        1.479s

	//todd's way
	// xs := strings.Fields(s)
	// return len(xs)

	// √ word % go test -bench .
	// goos: darwin
	// goarch: amd64
	// pkg: github.com/imattf/go211/ninj1300/02/word
	// BenchmarkCount-6           17313             68870 ns/op
	// PASS
	// ok      github.com/imattf/go211/ninj1300/02/word        2.767s
}
