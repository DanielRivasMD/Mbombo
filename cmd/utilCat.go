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

// catFiles concatenates files specified in options.Files into the out file (options.OutPath)
// It checks if the out file exists using FileExist (with an inline no-action) & removes it if it does
// All errors are propagated using horus.PropagateErr
func catFiles(options forgeOptions) error {
	op := "cat"

	// Clean prior copying:
	// Use an inline anonymous function for the NotFoundAction
	exists, err := domovoi.FileExist(options.OutPath, func(path string) (bool, error) {
		// No action required if the file is missing
		return false, nil
	}, verbose)
	if err != nil {
		return horus.PropagateErr(
			op,
			"file_exist_error",
			"failed to check out file existence",
			err,
			map[string]any{
				"outpath": options.OutPath,
			},
		)
	}

	// If the file exists, remove it
	if exists {
		if err := os.Remove(options.OutPath); err != nil {
			return horus.PropagateErr(
				op,
				"file_remove_error",
				"failed to remove existing out file",
				err,
				map[string]any{
					"outpath": options.OutPath,
				},
			)
		}
	}

	// Open the out file for appending, creating it if necessary
	fwrite, err := os.OpenFile(options.OutPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return horus.PropagateErr(
			op,
			"open_file_error",
			"failed to open out file",
			err,
			map[string]any{
				"outpath": options.OutPath,
			},
		)
	}
	defer fwrite.Close()

	// Loop through each source file
	for _, file := range options.Files {
		srcPath := options.InPath + "/" + file
		fread, err := os.Open(srcPath)
		if err != nil {
			return horus.PropagateErr(
				op,
				"open_source_error",
				"failed to open source file",
				err,
				map[string]any{
					"source": srcPath,
				},
			)
		}
		// Close the source file when done
		defer fread.Close()

		// Create a buffered writer for the out file
		writer := bufio.NewWriter(fwrite)
		scanner := bufio.NewScanner(fread)
		for scanner.Scan() {
			toPrint := scanner.Text() + "\n"
			if _, err := writer.WriteString(toPrint); err != nil {
				return horus.PropagateErr(
					op,
					"write_error",
					"failed to write to out file",
					err,
					map[string]any{
						"line": scanner.Text(),
					},
				)
			}
		}
		if err := scanner.Err(); err != nil {
			return horus.PropagateErr(
				op,
				"scan_error",
				"error reading from source file",
				err,
				nil,
			)
		}
		// Flush buffered writes to the out file
		if err := writer.Flush(); err != nil {
			return horus.PropagateErr(
				op,
				"flush_error",
				"failed to flush writer",
				err,
				map[string]any{
					"outpath": options.OutPath,
				},
			)
		}
	}

	// Perform replacements if provided
	if err := replace(options.OutPath, replacePairs); err != nil {
		return horus.PropagateErr(
			op,
			"replace_error",
			"failed to perform replacements in the out file",
			err,
			map[string]any{
				"outpath": options.OutPath,
			},
		)
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
