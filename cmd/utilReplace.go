////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bufio"
	"log"
	"os"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// replace values
func replace(target string, reps []rep) {
	// open reader
	fread, ε := os.Open(target)
	if ε != nil {
		log.Fatal(ε)
	}
	defer fread.Close()

	// open writer
	fwrite, ε := os.OpenFile(target, os.O_WRONLY|os.O_CREATE, 0666)
	if ε != nil {
		log.Fatal(ε)
	}
	defer fwrite.Close()

	// declare writer
	ϖ := bufio.NewWriter(fwrite)

	// read file
	scanner := bufio.NewScanner(fread)

	// scan file
	for scanner.Scan() {
		// preallocate
		toPrint := scanner.Text()

		// iterate replacements
		for _, rep := range reps {
			// replace
			toPrint = strings.Replace(toPrint, rep.old, rep.new, -1)
		}

		// format
		toPrint = toPrint + "\n"

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

////////////////////////////////////////////////////////////////////////////////////////////////////
