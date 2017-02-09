package main

import (
	"fmt"
	"log"
	"os/exec"
)

func execute(file string) string {

	// here we perform the pwd command.
	// we can store the output of this in our out variable
	// and catch any errors in err
	out, err := exec.Command("sh", file).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Command Successfully Executed")
	// as the out variable defined above is of type []byte we need to convert
	// this to a string or else we will see garbage printed out in our console
	// this is how we convert it to a string
	output := string(out[:])

	return output
}
