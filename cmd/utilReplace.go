////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bufio"
	"os"
	"strings"

	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// replace applies the provided replacements on the content of the target file.
// It opens the file for reading and writing, iterates through each line, applies the replacements,
// and writes the modified content back into the file. Any errors are wrapped and propagated.
func replace(target string, replacements []replacement) error {
	// Open the target file for reading.
	fread, err := os.Open(target)
	if err != nil {
		return horus.PropagateErr(
			"replace",
			"file_open_read_error",
			"failed to open file for reading",
			err,
			map[string]any{"target": target},
		)
	}
	defer fread.Close()

	// Open the same file for writing.
	fwrite, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return horus.PropagateErr(
			"replace",
			"file_open_write_error",
			"failed to open file for writing",
			err,
			map[string]any{"target": target},
		)
	}
	defer fwrite.Close()

	// Create a buffered writer.
	writer := bufio.NewWriter(fwrite)

	// Create a scanner to read the file by lines.
	scanner := bufio.NewScanner(fread)
	for scanner.Scan() {
		line := scanner.Text()

		// Apply each replacement.
		for _, rep := range replacements {
			line = strings.Replace(line, rep.Old, rep.New, -1)
		}

		// Append a newline.
		line = line + "\n"

		// Write the modified line.
		_, err = writer.WriteString(line)
		if err != nil {
			return horus.PropagateErr(
				"replace",
				"write_error",
				"failed to write modified line to target file",
				err,
				map[string]any{"line": line},
			)
		}
	}

	// Check for scanning errors.
	if err := scanner.Err(); err != nil {
		return horus.PropagateErr(
			"replace",
			"scanner_error",
			"an error occurred during scanning",
			err,
			nil,
		)
	}

	// Flush the writer.
	if err := writer.Flush(); err != nil {
		return horus.PropagateErr(
			"replace",
			"flush_error",
			"failed to flush writer",
			err,
			map[string]any{"target": target},
		)
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
