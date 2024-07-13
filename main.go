package main

import (
	"fmt"
	"os"
	"os/user"
	"toy-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the toy interpreter\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
