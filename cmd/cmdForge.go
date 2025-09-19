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
	files   []string
}

type replacement struct {
	old  string // anchor or token
	new  string // full replacement string
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
	forgeCmd.Flags().StringArrayVarP(&options.files, "files", "", []string{}, "These items will create...")
	forgeCmd.Flags().VarP(&replacePairs, "replace", "r", "replacement in form old=new, comma-separated")

	horus.CheckErr(forgeCmd.MarkFlagRequired("in"))
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

	horus.CheckErr(
		catFiles(options),
		horus.WithOp(op),
		horus.WithMessage("Error during concatenation execution"),
		horus.WithExitCode(2),
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// catFiles concatenates files specified in options.Files into the out file (options.OutPath)
// It checks if the out file exists using FileExist (with an inline no-action) & removes it if it does
// All errors are propagated using horus.PropagateErr
func catFiles(options forgeOptions) error {
	op := "cat"

	exists, err := domovoi.FileExist(options.outPath, func(path string) (bool, error) {
		return false, nil
	}, verbose)
	if err != nil {
		return horus.PropagateErr(
			op,
			"file_exist_error",
			"failed to check out file existence",
			err,
			map[string]any{"outpath": options.outPath},
		)
	}

	if exists {
		if err := os.Remove(options.outPath); err != nil {
			return horus.PropagateErr(
				op,
				"file_remove_error",
				"failed to remove existing out file",
				err,
				map[string]any{"outpath": options.outPath},
			)
		}
	}

	fwrite, err := os.OpenFile(options.outPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return horus.PropagateErr(
			op,
			"open_file_error",
			"failed to open out file",
			err,
			map[string]any{"outpath": options.outPath},
		)
	}
	defer fwrite.Close()

	writer := bufio.NewWriter(fwrite)

	for _, file := range options.files {
		srcPath := options.inPath + "/" + file
		fread, err := os.ReadFile(srcPath)
		if err != nil {
			return horus.PropagateErr(
				op,
				"read_source_error",
				"failed to read source file",
				err,
				map[string]any{"source": srcPath},
			)
		}

		// Apply replacements before writing
		content := applyReplacements(string(fread), replacePairs)

		if _, err := writer.WriteString(content + "\n"); err != nil {
			return horus.PropagateErr(
				op,
				"write_error",
				"failed to write to out file",
				err,
				map[string]any{"source": srcPath},
			)
		}
	}

	if err := writer.Flush(); err != nil {
		return horus.PropagateErr(
			op,
			"flush_error",
			"failed to flush writer",
			err,
			map[string]any{"outpath": options.outPath},
		)
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
