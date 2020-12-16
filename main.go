package main

import (
	"os"

	"github.com/mmorg031/Blockchain-Go.git/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}

	cmd.Run()

}
