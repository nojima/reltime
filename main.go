package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: reltime DURATION\n")
		os.Exit(2)
	}

	d, err := time.ParseDuration(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v", err)
		os.Exit(1)
	}
	t := time.Now().UTC().Add(d)
	fmt.Println(t.Format(time.RFC3339))
}
