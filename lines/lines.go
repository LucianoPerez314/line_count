package lines

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CountLines(fileptr io.Reader) {
	//Read from file
	//Iniialize scanner for file and make it scan line by line.
	scanner := bufio.NewScanner(fileptr)
	numLines := 0
	// Scanner.scan returns a boolean.
	for scanner.Scan() {
		fmt.Println("Readin from file:", scanner.Text()) // Println will add back the final '\n'
		numLines++
	}

	// Check for errors.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	} else {
		fmt.Println("Number of lines found:", numLines)
	}
}
