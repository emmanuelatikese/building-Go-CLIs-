package main

import (
	"bufio"
	"io"
)
func CountingWords(r io.Reader, lines bool) int {
    scanner := bufio.NewScanner(r)
    if !lines {
            scanner.Split(bufio.ScanWords)
    }

    num_words := 0
    for scanner.Scan() {
        num_words ++
    }
    return num_words
}