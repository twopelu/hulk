package main

import(
	"fmt"
	"os"
	//"github.com/twopelu/hulk/cli"
	//"github.com/twopelu/hulk/gui"
	"github.com/twopelu/hulk/lib"
)

func main() {
	//cli.Launch()
	//gui.Launch()

	args := os.Args[1:] // args[0] is the program path
	if (args == nil || len(args) != 3) {
		fmt.Println("Hulk usage is: hulk <user> <pass> <owner>")
		return
	}

	// TODO: Get owner from user email

	user := args[0]
	pass := args[1]
	owner := args[2]

	lib.ConnectAndListRepos(user, pass, owner)
}