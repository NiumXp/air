package main

import (
	"bufio"
	"flag"
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

	f, err := os.Open(*pathFile)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
