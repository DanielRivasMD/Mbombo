/*
Copyright © 2024 danielrivasmd@gmail.com

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO: evaluate argument refactoring => one struct to capture all flags
// TODO: evaluate reading toml config
var (
	options      forgeOptions
	replacePairs replaceFlags
)

type forgeOptions struct {
	inPath  string
	outPath string
	outFile string
	inFiles []string
}

type replacement struct {
	old string // anchor or token
	new string // full replacement string
	// TODO: add tab completion?
	mode string // "token" or "line"
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// implement pflag.Value
type replaceFlags []replacement

func (r *replaceFlags) String() string {
	parts := make([]string, len(*r))
	for i, rep := range *r {
		parts[i] = rep.old + "=" + rep.new
	}
	return strings.Join(parts, ",")
}

func (r *replaceFlags) Set(val string) error {
	op := "flag.set"

	// Always split key/value first
	parts := strings.SplitN(val, "=", 2)
	if len(parts) != 2 {
		horus.CheckErr(
			errors.New(""),
			horus.WithOp(op),
			horus.WithMessage("invalid replace pair"),
			horus.WithExitCode(2),
			horus.WithFormatter(func(he *horus.Herror) string { return he.Message }),
		)
	}

	old := parts[0]
	newVal := parts[1]
	mode := "token"

	// Optional mode suffix on the RIGHT side: ...=VALUE:line or ...=VALUE:token
	switch {
	case strings.HasSuffix(newVal, ":line"):
		mode = "line"
		newVal = strings.TrimSuffix(newVal, ":line")
	case strings.HasSuffix(newVal, ":token"):
		mode = "token"
		newVal = strings.TrimSuffix(newVal, ":token")
	}

	*r = append(*r, replacement{old: old, new: newVal, mode: mode})
	return nil
}

func (r *replaceFlags) Type() string {
	return "old=new[:line|:token]"
}

func applyReplacements(content string, reps []replacement) string {
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		for _, rep := range reps {
			switch rep.mode {
			case "line":
				if strings.Contains(line, rep.old) {
					lines[i] = rep.new
				}
			default: // "token" or empty
				lines[i] = strings.ReplaceAll(lines[i], rep.old, rep.new)
			}
		}
	}

	return strings.Join(lines, "\n")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(forgeCmd)

	forgeCmd.Flags().StringVarP(&options.inPath, "in", "", "", "Where are the itmes to be forged?")
	forgeCmd.Flags().StringVarP(&options.outPath, "out", "", "", "Where will the forge be delivered?")
	forgeCmd.Flags().StringArrayVarP(&options.inFiles, "files", "", []string{}, "These items will create...")
	forgeCmd.Flags().VarP(&replacePairs, "replace", "", "replacement in form old=new, comma-separated")

	// TODO: error out as one-liner if required flags are absent
	// horus.CheckErr(forgeCmd.MarkFlagRequired("in"))
	horus.CheckErr(forgeCmd.MarkFlagRequired("out"))
	horus.CheckErr(forgeCmd.MarkFlagRequired("files"))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var forgeCmd = &cobra.Command{
	Use:     "forge",
	Short:   "Forge products",
	Long:    helpForge,
	Example: exampleForge,

	Run: runForge,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runForge(cmd *cobra.Command, args []string) {
	op := "mbombo.forge"

	normalizeForgeOptions(&options)

	horus.CheckErr(
		catFiles(options),
		horus.WithOp(op),
		horus.WithMessage("Error during concatenation execution"),
		horus.WithExitCode(2),
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func catFiles(options forgeOptions) error {
	op := "cat"

	// Detect overwrite mode
	overwrite := false
	if len(options.inFiles) == 1 && options.inFiles[0] == options.outFile {
		overwrite = true
	} else if len(options.inFiles) > 1 && slices.Contains(options.inFiles, options.outFile) {
		return horus.PropagateErr(
			op,
			"overwrite_conflict",
			"cannot overwrite output when multiple input files are used",
			errors.New("overwrite conflict"),
			map[string]any{
				"outpath": options.outPath,
				"files":   options.inFiles,
			},
		)
	}

	// Prepare source files
	var sourceFiles []string
	if overwrite {
		tmpFile := options.outFile + ".tmp"
		src := filepath.Join(options.outPath, options.outFile)
		dst := filepath.Join(options.outPath, tmpFile)

		if err := domovoi.CopyFile(src, dst); err != nil {
			return horus.PropagateErr(op, "copy_error", "failed to copy file for overwrite", err, map[string]any{
				"source": src,
				"temp":   dst,
			})
		}
		sourceFiles = []string{dst}
		defer os.Remove(dst)
	} else {
		for _, f := range options.inFiles {
			sourceFiles = append(sourceFiles, filepath.Join(options.inPath, f))
		}
	}

	// Clean prior output (only now!)
	outFull := filepath.Join(options.outPath, options.outFile)
	exists, err := domovoi.FileExist(outFull, func(path string) (bool, error) {
		return false, nil
	}, verbose)
	if err != nil {
		return horus.PropagateErr(op, "file_exist_error", "failed to check out file existence", err, map[string]any{"outpath": outFull})
	}
	if exists {
		if err := os.Remove(outFull); err != nil {
			return horus.PropagateErr(op, "file_remove_error", "failed to remove existing out file", err, map[string]any{"outpath": outFull})
		}
	}

	// Prepare output writer
	fwrite, err := os.OpenFile(outFull, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return horus.PropagateErr(op, "open_file_error", "failed to open out file", err, map[string]any{"outpath": outFull})
	}
	defer fwrite.Close()
	writer := bufio.NewWriter(fwrite)

	// Process each file
	for _, srcPath := range sourceFiles {
		fread, err := os.ReadFile(srcPath)
		if err != nil {
			return horus.PropagateErr(op, "read_source_error", "failed to read source file", err, map[string]any{"source": srcPath})
		}

		content := applyReplacements(string(fread), replacePairs)

		if _, err := writer.WriteString(content + "\n"); err != nil {
			return horus.PropagateErr(op, "write_error", "failed to write to out file", err, map[string]any{"source": srcPath})
		}
	}

	if err := writer.Flush(); err != nil {
		return horus.PropagateErr(op, "flush_error", "failed to flush writer", err, map[string]any{"outpath": outFull})
	}

	return nil
}

func normalizeForgeOptions(opts *forgeOptions) {
	// normalize input file path if only one
	if len(opts.inFiles) == 1 {
		full := opts.inFiles[0]
		dir := filepath.Dir(full)
		base := filepath.Base(full)

		// if the path contains a directory, update inPath and inFiles
		if dir != "." && strings.Contains(full, string(filepath.Separator)) {
			opts.inPath = dir
			opts.inFiles = []string{base}
		}
	}

	// normalize output path
	if opts.outPath != "" {
		dir := filepath.Dir(opts.outPath)
		base := filepath.Base(opts.outPath)

		// if outPath contains a file, split it
		if dir != "." && strings.Contains(opts.outPath, string(filepath.Separator)) {
			opts.outPath = dir
			opts.outFile = base
		} else {
			// if outPath is just a file name, treat current dir as path
			opts.outFile = opts.outPath
			opts.outPath = "."
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
