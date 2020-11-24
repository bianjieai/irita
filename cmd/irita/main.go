package main

import (
	"os"

	"github.com/bianjieai/irita/cmd/irita/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
