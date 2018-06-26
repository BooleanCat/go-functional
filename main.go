package main

import "os"

func main() {
	_, err := parseArgs()
	if err != nil {
		os.Exit(1)
	}
}
