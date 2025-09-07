////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bufio"
	"os"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// catFiles concatenates files specified in π.files into the destination file (π.dest).
// It checks if the destination file exists using FileExist (with an inline no-action)
// and removes it if it does. All errors are propagated using horus.PropagateErr.
func catFiles(options forgeOptions) error {
	// Clean prior copying:
	// Use an inline anonymous function for the NotFoundAction
	exists, err := domovoi.FileExist(options.OutPath, func(path string) (bool, error) {
		// No action required if the file is missing.
		return false, nil
	}, verbose)
	if err != nil {
		return horus.PropagateErr(
			"catFiles",
			"file_exist_error",
			"failed to check destination file existence",
			err,
			map[string]any{"dest": options.OutPath},
		)
	}

	// If the file exists, remove it.
	if exists {
		if err := os.Remove(options.OutPath); err != nil {
			return horus.PropagateErr(
				"catFiles",
				"file_remove_error",
				"failed to remove existing destination file",
				err,
				map[string]any{"dest": options.OutPath},
			)
		}
	}

	// Open the destination file for appending, creating it if necessary.
	fwrite, err := os.OpenFile(options.OutPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return horus.PropagateErr(
			"catFiles",
			"open_file_error",
			"failed to open destination file",
			err,
			map[string]any{"dest": options.OutPath},
		)
	}
	defer fwrite.Close()

	// Loop through each source file.
	for _, file := range options.Files {
		srcPath := options.InPath + "/" + file
		fread, err := os.Open(srcPath)
		if err != nil {
			return horus.PropagateErr(
				"catFiles",
				"open_source_error",
				"failed to open source file",
				err,
				map[string]any{"source": srcPath},
			)
		}
		// Close the source file when done.
		defer fread.Close()

		// Create a buffered writer for the destination file.
		writer := bufio.NewWriter(fwrite)
		scanner := bufio.NewScanner(fread)
		for scanner.Scan() {
			toPrint := scanner.Text() + "\n"
			if _, err := writer.WriteString(toPrint); err != nil {
				return horus.PropagateErr(
					"catFiles",
					"write_error",
					"failed to write to destination file",
					err,
					map[string]any{"line": scanner.Text()},
				)
			}
		}
		if err := scanner.Err(); err != nil {
			return horus.PropagateErr(
				"catFiles",
				"scan_error",
				"error reading from source file",
				err,
				nil,
			)
		}
		// Flush buffered writes to the destination file.
		if err := writer.Flush(); err != nil {
			return horus.PropagateErr(
				"catFiles",
				"flush_error",
				"failed to flush writer",
				err,
				map[string]any{"dest": options.OutPath},
			)
		}
	}

	// Perform replacements if provided.
	// if len(options.reps) > 0 {
		if err := replace(options.OutPath, replacePairs); err != nil {
			return horus.PropagateErr(
				"catFiles",
				"replace_error",
				"failed to perform replacements in the destination file",
				err,
				map[string]any{"dest": options.OutPath},
			)
		}
	// }

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
