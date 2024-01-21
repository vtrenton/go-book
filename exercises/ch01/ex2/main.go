package main

import (
	"fmt"
	"os"
)

func main() {
	for ind, val := range os.Args[1:] {
		fmt.Printf("Index: %d Value: %s\n", ind, val)
	}
}
