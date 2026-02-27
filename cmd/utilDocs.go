package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/DanielRivasMD/domovoi"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

//go:embed docs.json
var docsFS embed.FS

////////////////////////////////////////////////////////////////////////////////////////////////////

type CommandOpt func(*cobra.Command)

////////////////////////////////////////////////////////////////////////////////////////////////////

type DocEntry struct {
	Use                   string     `json:"use"`
	Aliases               []string   `json:"aliases,omitempty"`
	Hidden                bool       `json:"hidden,omitempty"`
	Short                 string     `json:"short,omitempty"`
	Long                  string     `json:"long"`
	ExampleUsages         [][]string `json:"example_usages,omitempty"`
	ValidArgs             []string   `json:"valid_args,omitempty"`
	DisableFlagsInUseLine bool       `json:"disable_flags_in_use_line,omitempty"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	once    sync.Once
	docs    map[string]DocEntry
	example map[string]string
	help    map[string]string
	short   map[string]string
	use     map[string]string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func formatHelp(text string, appName string) string {
	if text == "" {
		return ""
	}
	if strings.Contains(text, "%[1]s") || strings.Contains(text, "%s") {
		return fmt.Sprintf(text, appName)
	}
	return text
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func styleLongHelp(text string) string {
	if text == "" {
		return ""
	}
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "$") {
			lines[i] = chalk.White.Color(line)
		}
		if strings.HasPrefix(trimmed, "#") {
			lines[i] = chalk.Dim.TextStyle(chalk.Cyan.Color(line))
		}
	}

	return strings.Join(lines, "\n")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func formatExample(app string, usages ...[]string) string {
	return domovoi.FormatExample(app, usages...)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func loadDocs() {
	docs = make(map[string]DocEntry)
	example = make(map[string]string)
	help = make(map[string]string)
	short = make(map[string]string)
	use = make(map[string]string)

	data, err := docsFS.ReadFile("docs.json")
	if err != nil {
		// In production, we want to panic if docs are missing
		panic(fmt.Sprintf("Failed to load embedded documentation: %v", err))
	}

	if err := json.Unmarshal(data, &docs); err != nil {
		panic(fmt.Sprintf("Failed to parse embedded documentation: %v", err))
	}

	for key, entry := range docs {
		use[key] = entry.Use
		short[key] = entry.Short

		formattedHelp := formatHelp(entry.Long, APP)
		help[key] = styleLongHelp(formattedHelp)

		if len(entry.ExampleUsages) > 0 {
			example[key] = formatExample(APP, entry.ExampleUsages...)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func GetDocs() map[string]DocEntry {
	once.Do(loadDocs)
	return docs
}

func GetExample(key string) string {
	once.Do(loadDocs)
	return example[key]
}

func GetHelp(key string) string {
	once.Do(loadDocs)
	return help[key]
}

func GetShort(key string) string {
	once.Do(loadDocs)
	return short[key]
}

func GetUse(key string) string {
	once.Do(loadDocs)
	return use[key]
}

func GetDocEntry(key string) (DocEntry, bool) {
	once.Do(loadDocs)
	entry, exists := docs[key]
	return entry, exists
}

////////////////////////////////////////////////////////////////////////////////////////////////////

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

	if len(entry.ValidArgs) > 0 {
		cmd.ValidArgs = entry.ValidArgs
	}

	if entry.DisableFlagsInUseLine {
		cmd.DisableFlagsInUseLine = true
	}

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
