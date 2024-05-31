package main

import (
	"fmt"
	"os"
	"os/user"
	"waiig/repl"
)

func main() {
	currUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hi %s! Welcome to the Waiig programming langauge!\n", currUser.Username)
	fmt.Printf("Type in commands here: \n")
	repl.Start(os.Stdin, os.Stdout)
}
