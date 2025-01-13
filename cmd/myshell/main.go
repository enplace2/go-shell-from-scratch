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

		trimmed := text[:len(text)-1]

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
		}
		if trimmed == "" {
			continue
		}

		if trimmed == "exit 0" {
			os.Exit(0)
		}

		fmt.Printf("%s: command not found\n", trimmed)
	}

}
