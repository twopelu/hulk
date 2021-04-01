package main

import (
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
	if args == nil || len(args) != 3 {
		fmt.Println("Hulk usage is: hulk <user> <pass> <owner>")
		return
	}

	// TODO: Get owner from user email

	user := args[0]
	pass := args[1]
	owner := args[2]

	bbman := lib.BitBucketManager{}
	bbman.Connect(user, pass)
	_, errlr := bbman.ListRepos(owner)
	if errlr != nil {
		panic(errlr)
	}

	// FIXME
	// goroutine 1 [running]:
	// github.com/ktrysmt/go-bitbucket.(*Client).requestUrl(0x0, 0xc000086480, 0x5f, 0x0, 0x0, 0x0, 0x5f, 0xc000146080)
	//         C:/Users/dani/go/pkg/mod/github.com/ktrysmt/go-bitbucket@v0.9.9/client.go:376 +0x8a
	// github.com/ktrysmt/go-bitbucket.(*Repository).ListFiles(0xc0002004e0, 0xc00007fe38, 0xc00007fe08, 0x3, 0x3, 0x29, 0x0)
	//         C:/Users/dani/go/pkg/mod/github.com/ktrysmt/go-bitbucket@v0.9.9/repository.go:220 +0x17f
	// github.com/twopelu/hulk/lib.(*BitBucketManager).ListFiles(0xc00007ff60, 0xc0002004e0, 0xc0000140d8, 0x7, 0x5, 0x8, 0x0, 0x0, 0x0)
	//         D:/DANI/go/hulk/lib/lib.go:60 +0x1ae
	// main.main()
	//         D:/DANI/go/hulk/main.go:34 +0x21d

	// repos, errlr := bbman.ListRepos(owner)
	// if errlr != nil {
	// 	panic(errlr)
	// }
	// _, errlf := bbman.ListFiles(&repos[4], owner)
	// if errlf != nil {
	// 	panic(errlf)
	// }
}
