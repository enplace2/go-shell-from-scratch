package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Uncomment this block to pass the first stage
		fmt.Print("$ ")

		// Wait for user input
		text, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
		}

		fmt.Printf("%s: command not found\n", text[:len(text)-1])
	}

}
