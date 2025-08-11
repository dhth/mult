package main

import (
	"os"

	"github.com/dhth/mult/internal/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
