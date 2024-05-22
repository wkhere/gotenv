package main

import (
	"fmt"
	"os"

	"github.com/wkhere/gotenv"
)

type action struct {
	files []string
}

func parseArgs(args []string) (a action, _ error) {
	// todo: parse flags
	a.files = args
	return a, nil
}

func run(a action) error {
	env, err := gotenv.Read(a.files...)
	if err != nil {
		return err
	}

	for _, e := range env {
		fmt.Printf("%s=%q\n", e.Key, e.Val)
	}
	return nil
}

func main() {
	a, err := parseArgs(os.Args[1:])
	if err != nil {
		die(2, err)
	}

	err = run(a)
	if err != nil {
		die(1, err)
	}
}

func die(code int, err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(code)
}
