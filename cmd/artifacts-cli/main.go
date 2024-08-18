package main

import (
	"fmt"
	"os"

	"github.com/0kate/artifactsmmo-scripts/internal/cli"
)

func main() {
	cmd := cli.NewRootCmd()

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
