package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/LuckyLuciano314/line_count/lines"
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

	lines.CountLines(fileptr)

	fileptr.Close()
}
