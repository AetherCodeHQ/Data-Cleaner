package main

import (
	"fmt"
	"os"
)

// data_cleaner - Clean messy datasets
func data_cleaner(path string) {
	fmt.Println("========================================")
	fmt.Println("  Data-Cleaner")
	fmt.Println("  Clean messy datasets")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	data_cleaner(path)
}
