package main

import (
	"fmt"
	"os"

	"github.com/kazufusa/orp"
)

func main() {
	if isHelpNeeded() {
		help()
		return
	}

	env := orp.NewEnv("PATH", ":")
	for i := len(os.Args) - 1; i >= 1; i-- {
		env.MoveToTop(os.Args[i])
	}
	fmt.Printf(env.Export())
}

func isHelpNeeded() bool {
	for _, arg := range os.Args {
		if arg == "--help" || arg == "-h" || arg == "/?" {
			return true
		}
	}
	return false
}

func help() {
	fmt.Printf(`orp can arrange the order of  PATH environment variable

Usage:
    orp [regexp of high-priority dir]...
`)
}
