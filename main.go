package main

import (
	"os"

	"github.com/Lewvy/interpreter/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
