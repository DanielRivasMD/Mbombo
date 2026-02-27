package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/DanielRivasMD/domovoi"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

//go:embed docs.json
var docsFS embed.FS

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

// formatHelp intelligently formats help text only if it contains format specifiers
func formatHelp(text string, appName string) string {
	if text == "" {
		return ""
	}

	// Check if the string contains any format specifiers
	if strings.Contains(text, "%[1]s") || strings.Contains(text, "%s") {
		// No need to replace %% anymore since JSON doesn't require escaping
		return fmt.Sprintf(text, appName)
	}

	// No format specifiers, return as-is
	return text
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// styleLongHelp applies styling to long help text
func styleLongHelp(text string) string {
	if text == "" {
		return ""
	}

	// Split into lines to style sections
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Style shell commands (lines starting with $)
		if strings.HasPrefix(trimmed, "$") {
			lines[i] = chalk.White.Color(line)
		}
		// Style comments (lines starting with #)
		if strings.HasPrefix(trimmed, "#") {
			lines[i] = chalk.Dim.TextStyle(chalk.Cyan.Color(line))
		}
	}

	return strings.Join(lines, "\n")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// formatExample builds a multi‐line example block
func formatExample(app string, usages ...[]string) string {
	return domovoi.FormatExample(app, usages...)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// loadDocs initializes all documentation maps from the embedded JSON
func loadDocs() {
	// Initialize maps
	docs = make(map[string]DocEntry)
	example = make(map[string]string)
	help = make(map[string]string)
	short = make(map[string]string)
	use = make(map[string]string)

	// Load and parse JSON
	data, err := docsFS.ReadFile("docs.json")
	if err != nil {
		// In production, we want to panic if docs are missing
		panic(fmt.Sprintf("Failed to load embedded documentation: %v", err))
	}

	if err := json.Unmarshal(data, &docs); err != nil {
		panic(fmt.Sprintf("Failed to parse embedded documentation: %v", err))
	}

	// Populate all the helper maps
	for key, entry := range docs {
		use[key] = entry.Use
		short[key] = entry.Short

		// Format and style help text
		formattedHelp := formatHelp(entry.Long, "mbombo")
		help[key] = styleLongHelp(formattedHelp)

		// Format examples if they exist
		if len(entry.ExampleUsages) > 0 {
			example[key] = formatExample("mbombo", entry.ExampleUsages...)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// GetDocs returns the complete documentation map
func GetDocs() map[string]DocEntry {
	once.Do(loadDocs)
	return docs
}

// GetExample returns the formatted example for a given command key
func GetExample(key string) string {
	once.Do(loadDocs)
	return example[key]
}

// GetHelp returns the formatted help text for a given command key
func GetHelp(key string) string {
	once.Do(loadDocs)
	return help[key]
}

// GetShort returns the short description for a given command key
func GetShort(key string) string {
	once.Do(loadDocs)
	return short[key]
}

// GetUse returns the usage string for a given command key
func GetUse(key string) string {
	once.Do(loadDocs)
	return use[key]
}

// GetDocEntry returns a specific documentation entry
func GetDocEntry(key string) (DocEntry, bool) {
	once.Do(loadDocs)
	entry, exists := docs[key]
	return entry, exists
}

////////////////////////////////////////////////////////////////////////////////////////////////////
