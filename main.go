package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/achmad-dev/leah/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err, "can't get current user")
	}
	fmt.Printf("Hello %s this is leah language!\n", user.Username)
	fmt.Printf("Give me your command")
	repl.Start(os.Stdin, os.Stdout)
}
