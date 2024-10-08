package main

import (
	"fmt"
	"os"

	"github.com/dhth/mult/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}
