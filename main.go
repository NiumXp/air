package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

const (
	EMPTY_STRING = ""
)

func main() {
	pathFile := flag.String("p", EMPTY_STRING, "Path to an air file.")
	flag.Parse()

	if *pathFile == EMPTY_STRING {
		flag.PrintDefaults()
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(*pathFile)
	if err != nil {
		log.Fatal(err)
	}

	content := string(bytes)
	log.Println(content)
}
