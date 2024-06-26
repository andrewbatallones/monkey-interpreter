package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/andrewbatallones/monkey_interpreter/repl"
)

func main() {
	// This will just grab the current user
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("feel free to type in commands\n")

	// Starts the repl loop
	repl.Start(os.Stdin, os.Stdout)
}
