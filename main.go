package main

import (
	"fmt"
	"os"
	user2 "os/user"
	"toy-interpreter/repl"
)

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Toy interpreter console\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
