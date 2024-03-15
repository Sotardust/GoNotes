package main

import (
	"GoNotes/src/demo1"
	"bufio"
	"os"
)

func main() {

	demo1.Main1()

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
