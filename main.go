package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Variables for command line
var urlPtr = flag.String("url", "", "Url for downloading file")
var filePtr = flag.String("file", "", "File output")

func init() {
	// Short version for long flag
	flag.StringVar(urlPtr, "u", "", "Url for downloading file")
	flag.StringVar(filePtr, "f", "", "File output")
}

func main() {
	// Parse our flags
	flag.Parse()

	// Validate the cli inputs
	validate(*urlPtr, *filePtr)

	// Get the url file
	resp, err := http.Get(*urlPtr)
	if err != nil {
		log.Fatal(err)
	}

	// Make sure we close connection
	defer resp.Body.Close()

	// Open file for writing
	file, err := os.Create(*filePtr)
	if err != nil {
		log.Fatal(err)
	}

	// io.Copy the respons to files
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	fmt.Printf("Your file is now written: %s\n", *filePtr)

	// Execute bootstrap and print the output
	output := execute(*filePtr)
	fmt.Println(output)

}
