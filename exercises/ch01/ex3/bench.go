package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	loopcat()  // tests creation of vars, for loop, concatination, print
	argslice() // tests slice sub, joining the string, and printing the result
}

func loopcat() {
	startTime := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	duration := time.Since(startTime)

	fmt.Println("The loopcat routine completed in ", duration)
}

// absolutely the winner
func argslice() {
	startTime := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	duration := time.Since(startTime)

	fmt.Println("The argslice routine completed in ", duration)
}
