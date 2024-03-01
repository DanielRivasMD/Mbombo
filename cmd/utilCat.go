////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bufio"
	"log"
	"os"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// copy file
func catFiles(params paramsCR) {
	// clean prior copying
	if fileExist(params.dest) {
		os.Remove(params.dest)
	}

	// open writer
	fwrite, ε := os.OpenFile(params.dest, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if ε != nil {
		log.Fatal(ε)
	}
	defer fwrite.Close()

	for _, file := range params.files {
		// open reader
		fread, ε := os.Open(params.orig + "/" + file)
		if ε != nil {
			log.Fatal(ε)
		}
		defer fread.Close()

		// declare writer
		ϖ := bufio.NewWriter(fwrite)

		// read file
		scanner := bufio.NewScanner(fread)

		// scan file
		for scanner.Scan() {
			// preallocate
			toPrint := scanner.Text() + "\n"
			// write
			_, ε = ϖ.WriteString(toPrint)
			if ε != nil {
				log.Fatal(ε)
			}
		}

		if ε := scanner.Err(); ε != nil {
			log.Fatal(ε)
		}

		// flush writer
		ϖ.Flush()
	}
	// replace
	if len(params.reps) > 0 {
		replace(params.dest, params.reps)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
