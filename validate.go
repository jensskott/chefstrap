package main

import (
	"flag"
	"log"
	"os"

	"github.com/asaskevich/govalidator"
)

func validate(url string, file string) {
	if url == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if file == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	testURL := govalidator.IsURL(url)
	if testURL == false {
		log.Fatal(url, " is not a valid url.")
		os.Exit(1)
	}

}
