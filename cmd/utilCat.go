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
func catFiles(path, out string, files []string) {
	// clean prior copying
	if fileExist(out) { os.Remove(out) }

	// open writer
	fwrite, ε := os.OpenFile(out, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if ε != nil {
		log.Fatal(ε)
	}
	defer fwrite.Close()

	for _, file := range files {
		// open reader
		fread, ε := os.Open(path + "/" + file)
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
}

////////////////////////////////////////////////////////////////////////////////////////////////////
