package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Sample function returns list of items chosen from the slice
func Sample(slice []string, n uint) ([]string, error) {
	if n > uint(len(slice)) {
		return nil, fmt.Errorf("N should be less than %v", len(slice)+1)
	}

	shuffledSlice := Shuffle(slice)
	return shuffledSlice[:n], nil
}

// Shuffle randomizes the items of a list in place.
func Shuffle(slice []string) []string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]string, len(slice))
	n := len(slice)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(slice))
		ret[i] = slice[randIndex]
		slice = append(slice[:randIndex], slice[randIndex+1:]...)
	}
	return ret
}

// TrimSpace returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
func TrimSpace(input string) string {
	return strings.TrimSpace(input)
}
