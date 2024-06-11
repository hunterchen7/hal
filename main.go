package main

import (
	"fmt"
	"hal/repl"
	"os"
	"os/user"
)

func main() {
	currUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hi %s! Welcome to the hal programming langauge!\n", currUser.Username)
	fmt.Printf("Type in commands here: \n")
	repl.Start(os.Stdin, os.Stdout)
}
