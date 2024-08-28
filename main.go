package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Read filename from cmd line
	filenameFlag := flag.String("f", "", "Indicates filename will be specified")

	flag.Parse()

	fmt.Println("Filename provided is", *filenameFlag)

	//Open file via filename
	fileptr, err := os.Open(*filenameFlag)
	if err != nil {
		log.Fatal(err)
	}

	fileptr.Close()
}
