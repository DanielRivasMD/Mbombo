package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

type CommandOpt func(*cobra.Command)

////////////////////////////////////////////////////////////////////////////////////////////////////

// authorHeader returns the formatted author information
func authorHeader() string {
	return chalk.Bold.TextStyle(
		chalk.Green.Color("Daniel Rivas "),
	) +
		chalk.Dim.TextStyle(
			chalk.Italic.TextStyle("<danielrivasmd@gmail.com>"),
		)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func MakeCmd(key string, run func(*cobra.Command, []string), opts ...CommandOpt) *cobra.Command {
	docs := GetDocs()
	entry, exists := docs[key]
	if !exists {
		keys := make([]string, 0, len(docs))
		for k := range docs {
			keys = append(keys, k)
		}
		log.Fatalf("No documentation found for command: %s. Available keys: %v", key, keys)
	}

	// Add author header to long help
	longHelp := GetHelp(key)
	if longHelp != "" {
		longHelp = authorHeader() + "\n\n" + longHelp
	} else {
		longHelp = authorHeader()
	}

	cmd := &cobra.Command{
		Use:     entry.Use,
		Short:   entry.Short,
		Long:    longHelp,
		Example: GetExample(key),
		Aliases: entry.Aliases,
		Hidden:  entry.Hidden,
		Run:     run,
	}

	// Set additional fields
	if len(entry.ValidArgs) > 0 {
		cmd.ValidArgs = entry.ValidArgs
	}

	if entry.DisableFlagsInUseLine {
		cmd.DisableFlagsInUseLine = true
	}

	// Apply all options
	for _, opt := range opts {
		opt(cmd)
	}

	return cmd
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func WithArgs(validator cobra.PositionalArgs) CommandOpt {
	return func(cmd *cobra.Command) {
		cmd.Args = validator
	}
}

func WithValidArgs(args []string) CommandOpt {
	return func(cmd *cobra.Command) {
		cmd.ValidArgs = args
	}
}

func WithAliases(aliases []string) CommandOpt {
	return func(cmd *cobra.Command) {
		cmd.Aliases = aliases
	}
}

func WithDisableFlagParsing(disable bool) CommandOpt {
	return func(cmd *cobra.Command) {
		cmd.DisableFlagParsing = disable
	}
}

func WithPreRun(preRun func(*cobra.Command, []string)) CommandOpt {
	return func(cmd *cobra.Command) {
		cmd.PreRun = preRun
	}
}

func WithPostRun(postRun func(*cobra.Command, []string)) CommandOpt {
	return func(cmd *cobra.Command) {
		cmd.PostRun = postRun
	}
}

func WithSilenceErrors(silence bool) CommandOpt {
	return func(cmd *cobra.Command) {
		cmd.SilenceErrors = silence
	}
}

func WithSilenceUsage(silence bool) CommandOpt {
	return func(cmd *cobra.Command) {
		cmd.SilenceUsage = silence
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
